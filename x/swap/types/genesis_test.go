package types_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/TxCorpi0x/magicswap/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				PartialSendList: []types.PartialSend{
					{
						Id:            0,
						Creator:       "creator",
						Recipient:     "recipient",
						BurntAmount:   sdk.NewCoin("stake", math.ZeroInt()),
						SentAmount:    sdk.NewCoin("stake", math.ZeroInt()),
						SwappedAmount: sdk.NewCoin("stake", math.ZeroInt()),
					},
					{
						Id:            1,
						Creator:       "creator1",
						Recipient:     "recipient1",
						BurntAmount:   sdk.NewCoin("stake", math.ZeroInt()),
						SentAmount:    sdk.NewCoin("stake", math.ZeroInt()),
						SwappedAmount: sdk.NewCoin("stake", math.ZeroInt()),
					},
				},
				PartialSendCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated partialSend",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				PartialSendList: []types.PartialSend{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid partialSend count",
			genState: &types.GenesisState{
				PartialSendList: []types.PartialSend{
					{
						Id: 1,
					},
				},
				PartialSendCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
