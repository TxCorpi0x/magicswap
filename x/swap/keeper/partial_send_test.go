package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/TxCorpi0x/magicswap/testutil/keeper"
	"github.com/TxCorpi0x/magicswap/testutil/nullify"
	"github.com/TxCorpi0x/magicswap/x/swap/keeper"
	"github.com/TxCorpi0x/magicswap/x/swap/types"
	"github.com/stretchr/testify/require"
)

func createNPartialSend(keeper keeper.Keeper, ctx context.Context, n int) []types.PartialSend {
	items := make([]types.PartialSend, n)
	for i := range items {
		items[i].Creator = "creator"
		items[i].Id = keeper.AppendPartialSend(ctx, items[i])
	}
	return items
}

func TestPartialSendGet(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	items := createNPartialSend(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetPartialSend(ctx, item.Creator, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestPartialSendGetAll(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	items := createNPartialSend(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPartialSend(ctx)),
	)
}

func TestPartialSendCount(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	items := createNPartialSend(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetPartialSendCount(ctx))
}
