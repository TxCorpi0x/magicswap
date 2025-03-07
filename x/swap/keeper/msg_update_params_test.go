package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/TxCorpi0x/magicswap/x/swap/types"
)

func TestMsgUpdateParams(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	params := types.DefaultParams()
	require.NoError(t, k.SetParams(ctx, params))
	wctx := sdk.UnwrapSDKContext(ctx)

	// default params
	testCases := []struct {
		name      string
		input     *types.MsgUpdateParams
		expErr    bool
		expErrMsg string
	}{
		{
			name: "invalid authority",
			input: &types.MsgUpdateParams{
				Authority: "invalid",
				Params:    params,
			},
			expErr:    true,
			expErrMsg: "invalid authority",
		},
		{
			name: "invalid ratio",
			input: &types.MsgUpdateParams{
				Authority: k.GetAuthority(),
				Params: types.Params{
					SwapRules: []types.SwapRule{
						{
							SrcDenom:            "stake",
							DstDenom:            "stake1",
							Ratio:               math.LegacyNewDecWithPrec(-1, 2), // invalid ratio
							MinSupplyRatioLimit: types.DefaultMinSupplyRatioLimit,
						},
					},
				},
			},
			expErr:    true,
			expErrMsg: "ratio must be between 0 and 1",
		},
		{
			name: "invalid min supply ratio limit",
			input: &types.MsgUpdateParams{
				Authority: k.GetAuthority(),
				Params: types.Params{
					SwapRules: []types.SwapRule{
						{
							SrcDenom:            "stake",
							DstDenom:            "stake1",
							Ratio:               types.DefaultRatio,
							MinSupplyRatioLimit: math.LegacyNewDecWithPrec(-1, 2), // invalid min supply ratio limit
						},
					},
				},
			},
			expErr:    true,
			expErrMsg: "min supply ratio limit must be between 0 and 1",
		},
		{
			name: "duplicate src and dst token pair",
			input: &types.MsgUpdateParams{
				Authority: k.GetAuthority(),
				Params: types.Params{
					SwapRules: []types.SwapRule{
						{
							SrcDenom:            "stake",
							DstDenom:            "stake1",
							Ratio:               types.DefaultRatio,
							MinSupplyRatioLimit: types.DefaultMinSupplyRatioLimit,
						},
						{
							SrcDenom:            "stake",
							DstDenom:            "stake2",
							Ratio:               types.DefaultRatio,
							MinSupplyRatioLimit: types.DefaultMinSupplyRatioLimit,
						},
					},
				},
			},
			expErr:    true,
			expErrMsg: "duplicate src and dst token pair found",
		},
		{
			name: "all good",
			input: &types.MsgUpdateParams{
				Authority: k.GetAuthority(),
				Params: types.Params{
					SwapRules: []types.SwapRule{
						{
							SrcDenom:            "stake",
							DstDenom:            "stake1",
							Ratio:               types.DefaultRatio,
							MinSupplyRatioLimit: types.DefaultMinSupplyRatioLimit,
						},
					},
				},
			},
			expErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ms.UpdateParams(wctx, tc.input)

			if tc.expErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expErrMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
