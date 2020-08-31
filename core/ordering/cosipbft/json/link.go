package json

import (
	"go.dedis.ch/dela/consensus/viewchange"
	"go.dedis.ch/dela/core/ordering/cosipbft/types"
	"go.dedis.ch/dela/serde"
	"golang.org/x/xerrors"
)

type linkFormat struct{}

func (fmt linkFormat) Encode(ctx serde.Context, msg serde.Message) ([]byte, error) {
	link, ok := msg.(types.BlockLink)
	if !ok {
		return nil, xerrors.Errorf("unsupported message '%T'", msg)
	}

	block, err := link.GetTo().Serialize(ctx)
	if err != nil {
		return nil, err
	}

	prepare, err := link.GetPrepareSignature().Serialize(ctx)
	if err != nil {
		return nil, err
	}

	commit, err := link.GetCommitSignature().Serialize(ctx)
	if err != nil {
		return nil, err
	}

	changeset, err := link.GetChangeSet().Serialize(ctx)
	if err != nil {
		return nil, err
	}

	m := BlockLinkJSON{
		From:             link.GetFrom(),
		Block:            block,
		PrepareSignature: prepare,
		CommitSignature:  commit,
		ChangeSet:        changeset,
	}

	data, err := ctx.Marshal(m)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (fmt linkFormat) Decode(ctx serde.Context, data []byte) (serde.Message, error) {
	m := BlockLinkJSON{}
	err := ctx.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	factory := ctx.GetFactory(types.BlockKey{})

	msg, err := factory.Deserialize(ctx, m.Block)
	if err != nil {
		return nil, err
	}

	block, ok := msg.(types.Block)
	if !ok {
		return nil, xerrors.Errorf("invalid block '%T'", msg)
	}

	prepare, err := decodeSignature(ctx, m.PrepareSignature)
	if err != nil {
		return nil, err
	}

	commit, err := decodeSignature(ctx, m.CommitSignature)
	if err != nil {
		return nil, err
	}

	changeset, err := decodeChangeSet(ctx, m.ChangeSet)
	if err != nil {
		return nil, err
	}

	opts := []types.BlockLinkOption{
		types.WithSignatures(prepare, commit),
		types.WithChangeSet(changeset),
	}

	link, err := types.NewBlockLink(m.From, block, opts...)
	if err != nil {
		return nil, err
	}

	return link, nil
}

func decodeChangeSet(ctx serde.Context, data []byte) (viewchange.ChangeSet, error) {
	factory := ctx.GetFactory(types.ChangeSetKey{})

	fac, ok := factory.(viewchange.ChangeSetFactory)
	if !ok {
		return nil, xerrors.Errorf("invalid change set factory '%T'", factory)
	}

	changeset, err := fac.ChangeSetOf(ctx, data)
	if err != nil {
		return nil, xerrors.Errorf("change set decoding failed: %v", err)
	}

	return changeset, nil
}
