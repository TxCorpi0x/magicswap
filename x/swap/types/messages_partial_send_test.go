package types

import (
	"testing"

	"github.com/TxCorpi0x/magicswap/testutil/sample"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgPartialSend_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgPartialSend
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgPartialSend{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgPartialSend{
				Creator:   sample.AccAddress(),
				Recipient: sample.AccAddress(),
				Amount:    sdk.NewInt64Coin("stake", 100),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
