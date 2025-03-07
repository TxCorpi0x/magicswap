package swap

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/TxCorpi0x/magicswap/testutil/sample"
	swapsimulation "github.com/TxCorpi0x/magicswap/x/swap/simulation"
	"github.com/TxCorpi0x/magicswap/x/swap/types"
)

// avoid unused import issue
var (
	_ = swapsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgPartialSend = "op_weight_msg_partial_send"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPartialSend int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	swapGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PartialSendList: []types.PartialSend{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		PartialSendCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&swapGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgPartialSend int
	simState.AppParams.GetOrGenerate(opWeightMsgPartialSend, &weightMsgPartialSend, nil,
		func(_ *rand.Rand) {
			weightMsgPartialSend = defaultWeightMsgPartialSend
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPartialSend,
		swapsimulation.SimulateMsgPartialSend(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgPartialSend,
			defaultWeightMsgPartialSend,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				swapsimulation.SimulateMsgPartialSend(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
