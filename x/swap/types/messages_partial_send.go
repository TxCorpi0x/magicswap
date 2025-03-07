package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgPartialSend{}

func NewMsgPartialSend(creator string, recipient string, amount sdk.Coin) *MsgPartialSend {
	return &MsgPartialSend{
		Creator:   creator,
		Recipient: recipient,
		Amount:    amount,
	}
}

func (msg *MsgPartialSend) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
