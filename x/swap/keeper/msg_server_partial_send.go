package keeper

import (
	"context"

	"github.com/TxCorpi0x/magicswap/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PartialSend(goCtx context.Context, msg *types.MsgPartialSend) (*types.MsgPartialSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var partialSend = types.PartialSend{
		Creator:   msg.Creator,
		Recipient: msg.Recipient,
		// TODO: add the swap logic and fill the fields
	}

	id := k.AppendPartialSend(
		ctx,
		partialSend,
	)

	return &types.MsgPartialSendResponse{
		Id: id,
	}, nil
}
