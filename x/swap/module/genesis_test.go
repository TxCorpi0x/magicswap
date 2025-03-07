package swap_test

import (
	"testing"

	keepertest "github.com/TxCorpi0x/magicswap/testutil/keeper"
	"github.com/TxCorpi0x/magicswap/testutil/nullify"
	swap "github.com/TxCorpi0x/magicswap/x/swap/module"
	"github.com/TxCorpi0x/magicswap/x/swap/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PartialSendList: []types.PartialSend{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PartialSendCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SwapKeeper(t)
	swap.InitGenesis(ctx, k, genesisState)
	got := swap.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PartialSendList, got.PartialSendList)
	require.Equal(t, genesisState.PartialSendCount, got.PartialSendCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
