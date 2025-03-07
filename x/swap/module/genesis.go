package swap

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/TxCorpi0x/magicswap/x/swap/keeper"
	"github.com/TxCorpi0x/magicswap/x/swap/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the partialSend
	for _, elem := range genState.PartialSendList {
		k.SetPartialSend(ctx, elem)
	}

	// Set partialSend count
	k.SetPartialSendCount(ctx, genState.PartialSendCount)
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PartialSendList = k.GetAllPartialSend(ctx)
	genesis.PartialSendCount = k.GetPartialSendCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
