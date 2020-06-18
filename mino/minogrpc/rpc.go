package minogrpc

import (
	context "context"

	"go.dedis.ch/dela/mino"
	"go.dedis.ch/dela/serde"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/metadata"
)

// RPC represents an RPC that has been registered by a client, which allows
// clients to call an RPC that will execute the provided handler.
//
// - implements mino.RPC
type RPC struct {
	overlay overlay
	uri     string
	factory serde.Factory
}

// Call implements mino.RPC. It calls the RPC on each provided address.
func (rpc *RPC) Call(ctx context.Context, req serde.Message,
	players mino.Players) (<-chan serde.Message, <-chan error) {

	out := make(chan serde.Message, players.Len())
	errs := make(chan error, players.Len())

	data, err := rpc.overlay.serializer.Serialize(req)
	if err != nil {
		errs <- xerrors.Errorf("failed to marshal msg to any: %v", err)
		return out, errs
	}

	sendMsg := &Message{Payload: data}

	go func() {
		iter := players.AddressIterator()
		for iter.HasNext() {
			addr := iter.GetNext()

			clientConn, err := rpc.overlay.connFactory.FromAddress(addr)
			if err != nil {
				errs <- xerrors.Errorf("failed to get client conn for '%v': %v",
					addr, err)
				continue
			}

			cl := NewOverlayClient(clientConn)

			header := metadata.New(map[string]string{headerURIKey: rpc.uri})
			newCtx := metadata.NewOutgoingContext(ctx, header)

			callResp, err := cl.Call(newCtx, sendMsg)
			if err != nil {
				errs <- xerrors.Errorf("failed to call client '%s': %v", addr, err)
				continue
			}

			var resp serde.Message
			err = rpc.overlay.serializer.Deserialize(callResp.GetPayload(), rpc.factory, &resp)
			if err != nil {
				errs <- xerrors.Errorf("couldn't unmarshal payload: %v", err)
				continue
			}

			out <- resp
		}

		close(out)
	}()

	return out, errs
}

// Stream implements mino.RPC. TODO: errors
func (rpc RPC) Stream(ctx context.Context,
	players mino.Players) (mino.Sender, mino.Receiver, error) {

	root := newRootAddress()

	rting, err := rpc.overlay.routingFactory.FromIterator(rpc.overlay.me,
		players.AddressIterator())
	if err != nil {
		return nil, nil, xerrors.Errorf("couldn't generate routing: %v", err)
	}

	header := metadata.New(map[string]string{headerURIKey: rpc.uri})

	receiver := receiver{
		serializer:     rpc.overlay.serializer,
		factory:        rpc.factory,
		addressFactory: rpc.overlay.routingFactory.GetAddressFactory(),
		errs:           make(chan error, 1),
		queue:          newNonBlockingQueue(),
	}

	gateway := rting.GetRoot()

	sender := sender{
		me:             root,
		serializer:     rpc.overlay.serializer,
		addressFactory: AddressFactory{},
		gateway:        gateway,
		clients:        map[mino.Address]chan OutContext{},
		receiver:       &receiver,
		traffic:        rpc.overlay.traffic,
	}

	relayCtx := metadata.NewOutgoingContext(ctx, header)

	// The orchestrator opens a connection to the entry point of the routing map
	// and it will relay the messages by this gateway by default. The entry
	// point of the routing will have the orchestrator stream opens which will
	// allow the messages to be routed back to the orchestrator.
	err = rpc.overlay.setupRelay(relayCtx, gateway, &sender, &receiver, rting)
	if err != nil {
		return nil, nil, xerrors.Errorf("couldn't setup relay: %v", err)
	}

	return sender, receiver, nil
}