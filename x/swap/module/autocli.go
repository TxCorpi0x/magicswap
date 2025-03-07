package swap

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/TxCorpi0x/magicswap/api/magicswap/swap"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "PartialSendAll",
					Use:       "list-partial-send",
					Short:     "List all partialSend",
				},
				{
					RpcMethod:      "PartialSendByCreator",
					Use:            "list-partial-send-by-creator [creator] ",
					Short:          "Shows all partialSend by a specific creator",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "creator"}},
				},
				{
					RpcMethod:      "PartialSend",
					Use:            "show-partial-send [creator] [id]",
					Short:          "Shows a partialSend by creator and id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "creator"}, {ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "PartialSend",
					Use:            "partial-send [recipient] [amount]",
					Short:          "Do a partialSend",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "recipient"}, {ProtoField: "amount"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
