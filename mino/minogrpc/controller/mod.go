package controller

import (
	"time"

	"go.dedis.ch/dela"
	"go.dedis.ch/dela/cli"
	"go.dedis.ch/dela/cli/node"
	"go.dedis.ch/dela/mino/minogrpc"
	"go.dedis.ch/dela/mino/minogrpc/routing"
	"golang.org/x/xerrors"
)

const (
	treeRoutingHeight = 3
)

// Minimal is an initializer with the minimum set of commands.
//
// - implements node.Initializer
type minimal struct{}

// NewMinimal returns a new initializer to start an instance of Minogrpc.
func NewMinimal() node.Initializer {
	return minimal{}
}

// Build implements node.Initializer. It populates the builder with the commands
// to control Minogrpc.
func (m minimal) SetCommands(builder node.Builder) {
	builder.SetStartFlags(
		cli.IntFlag{
			Name:  "port",
			Usage: "set the port to listen on",
			Value: 2000,
		},
	)

	cmd := builder.SetCommand("minogrpc")
	cmd.SetDescription("Network overlay administration")

	sub := cmd.SetSubCommand("certificates")
	sub.SetDescription("list the certificates of the server")
	sub.SetAction(builder.MakeAction(certAction{}))

	sub = cmd.SetSubCommand("token")
	sub.SetDescription("generate a token to share to others to join the network")
	sub.SetFlags(
		cli.DurationFlag{
			Name:  "expiration",
			Usage: "amount of time before expiration",
			Value: 24 * time.Hour,
		},
	)
	sub.SetAction(builder.MakeAction(tokenAction{}))

	sub = cmd.SetSubCommand("join")
	sub.SetDescription("join a network of participants")
	sub.SetFlags(
		cli.StringFlag{
			Name:     "token",
			Usage:    "secret token generated by the node to join",
			Required: true,
		},
		cli.StringFlag{
			Name:     "address",
			Usage:    "address of the node to join",
			Required: true,
		},
		cli.StringFlag{
			Name:     "cert-hash",
			Usage:    "certificate hash of the distant server",
			Required: true,
		},
	)
	sub.SetAction(builder.MakeAction(joinAction{}))
}

// Run implements node.Initializer. It starts the minogrpc instance and inject
// it in the dependency resolver.
func (m minimal) Inject(ctx cli.Flags, inj node.Injector) error {
	rf := routing.NewTreeRoutingFactory(treeRoutingHeight, minogrpc.AddressFactory{})

	port := ctx.Int("port")
	if port < 0 || port > 65535 {
		return xerrors.Errorf("invalid port value %d", port)
	}

	o, err := minogrpc.NewMinogrpc("127.0.0.1", uint16(port), rf)
	if err != nil {
		return xerrors.Errorf("couldn't make overlay: %v", err)
	}

	inj.Inject(o)

	dela.Logger.Info().Msgf("%v is running", o)

	return nil
}