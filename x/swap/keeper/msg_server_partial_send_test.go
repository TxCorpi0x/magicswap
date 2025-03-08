package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/stretchr/testify/require"

	"github.com/TxCorpi0x/magicswap/testutil/sample"
	"github.com/TxCorpi0x/magicswap/x/swap/types"
)

const (
	srcToken = "stake"
	dstToken = "stake1"
)

var (
	defaultSwapRules = []types.SwapRule{
		{
			SrcDenom:            srcToken,
			DstDenom:            dstToken,
			Ratio:               math.LegacyNewDecWithPrec(50, 2),
			MinSupplyRatioLimit: math.LegacyNewDecWithPrec(6667, 4),
		},
	}
)

// TestPartialSendMsgServerEven tests the send and swap process in the iterations.
// The supply and balance of the tokens gets calculated after each iteration and validated.
// This checks the even value for each amount.
func TestPartialSendMsgServerEven(t *testing.T) {
	k, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	k.SetParams(ctx, types.NewParams(defaultSwapRules))

	creator := sample.AccAddress()
	// the supply is set as 100 because in each iterations we have
	// 10 tokens burnt and swapped and 10 tokens sent.
	initialStakeSupply := sdk.NewCoin(srcToken, math.NewInt(100))
	initialStake2Supply := sdk.NewCoin(dstToken, math.ZeroInt())
	creatorBalance := sdk.NewCoins(initialStakeSupply)
	require.NoError(t, k.BankKeeper.MintCoins(ctx, minttypes.ModuleName, creatorBalance))
	require.NoError(t, k.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, sdk.MustAccAddressFromBech32(creator), creatorBalance))

	recipient := sample.AccAddress()
	for i := 0; i < 5; i++ {
		resp, err := srv.PartialSend(wctx, &types.MsgPartialSend{
			Creator:   creator,
			Recipient: recipient,
			Amount:    sdk.NewCoin(srcToken, math.NewInt(20)),
		})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))

		stakeSupply := k.BankKeeper.GetSupply(ctx, srcToken)
		require.Equal(t, initialStakeSupply.Sub(sdk.NewCoin(srcToken, math.NewInt(int64(i+1)*10))), stakeSupply)

		stake2Supply := k.BankKeeper.GetSupply(ctx, dstToken)
		require.Equal(t, initialStake2Supply.Add(sdk.NewCoin(dstToken, math.NewInt(int64(i+1)*10))), stake2Supply)

		recipientBalance := k.BankKeeper.SpendableCoins(ctx, sdk.MustAccAddressFromBech32(recipient))
		require.Equal(t, math.NewInt(int64(i+1)*10), recipientBalance.AmountOf(srcToken))
		require.Equal(t, math.NewInt(int64(i+1)*10), recipientBalance.AmountOf(dstToken))
	}
}

// TestPartialSendMsgServerOdd tests the send and swap process in the iterations.
// The supply and balance of the tokens gets calculated after each iteration and validated.
// This checks the odd value for each amount.
func TestPartialSendMsgServerOdd(t *testing.T) {
	k, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	k.SetParams(ctx, types.NewParams(defaultSwapRules))

	creator := sample.AccAddress()
	// the supply is set as 105 because in each iterations we have 11 tokens sent
	// and 10 tokens burnt and swapped.
	initialStakeSupply := sdk.NewCoin(srcToken, math.NewInt(105))
	initialStake2Supply := sdk.NewCoin(dstToken, math.ZeroInt())
	creatorBalance := sdk.NewCoins(initialStakeSupply)
	require.NoError(t, k.BankKeeper.MintCoins(ctx, minttypes.ModuleName, creatorBalance))
	require.NoError(t, k.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, sdk.MustAccAddressFromBech32(creator), creatorBalance))

	recipient := sample.AccAddress()
	for i := 0; i < 5; i++ {
		amount := math.NewInt(21)
		resp, err := srv.PartialSend(wctx, &types.MsgPartialSend{
			Creator:   creator,
			Recipient: recipient,
			Amount:    sdk.NewCoin(srcToken, amount),
		})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))

		amountDec := math.LegacyNewDecFromInt(amount)
		amountToBurn := defaultSwapRules[0].Ratio.Mul(amountDec).RoundInt()
		amountToSend := amount.Sub(amountToBurn)

		stakeSupply := k.BankKeeper.GetSupply(ctx, srcToken)
		require.Equal(t, initialStakeSupply.Sub(sdk.NewCoin(srcToken, math.NewInt(int64(i+1)*amountToBurn.Int64()))), stakeSupply)

		stake2Supply := k.BankKeeper.GetSupply(ctx, dstToken)
		require.Equal(t, initialStake2Supply.Add(sdk.NewCoin(dstToken, math.NewInt(int64(i+1)*amountToBurn.Int64()))), stake2Supply)

		recipientBalance := k.BankKeeper.SpendableCoins(ctx, sdk.MustAccAddressFromBech32(recipient))
		require.Equal(t, math.NewInt(int64(i+1)*amountToSend.Int64()), recipientBalance.AmountOf(srcToken))
		require.Equal(t, math.NewInt(int64(i+1)*amountToBurn.Int64()), recipientBalance.AmountOf(dstToken))
	}
}

func TestPartialSendMsgServerSelfSwap(t *testing.T) {
	k, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	k.SetParams(ctx, types.NewParams(defaultSwapRules))

	creator := sample.AccAddress()
	initialStakeSupply := sdk.NewCoin(srcToken, math.NewInt(200))
	initialStake2Supply := sdk.NewCoin(dstToken, math.ZeroInt())
	creatorBalance := sdk.NewCoins(initialStakeSupply)
	require.NoError(t, k.BankKeeper.MintCoins(ctx, minttypes.ModuleName, creatorBalance))
	require.NoError(t, k.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, sdk.MustAccAddressFromBech32(creator), creatorBalance))

	for i := 0; i < 5; i++ {
		resp, err := srv.PartialSend(wctx, &types.MsgPartialSend{
			Creator:   creator,
			Recipient: creator,
			Amount:    sdk.NewCoin(srcToken, math.NewInt(20)),
		})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))

		stake2Supply := k.BankKeeper.GetSupply(ctx, dstToken)
		require.Equal(t, initialStake2Supply.Add(sdk.NewCoin(dstToken, math.NewInt(int64(i+1)*20))), stake2Supply)

		balance := k.BankKeeper.SpendableCoins(ctx, sdk.MustAccAddressFromBech32(creator))
		require.Equal(t, initialStakeSupply.Amount.Sub(math.NewInt(int64(i+1)*20)), balance.AmountOf(srcToken))
		require.Equal(t, math.NewInt(int64(i+1)*20), balance.AmountOf(dstToken))
	}
}

func TestPartialSendMsgServer_SanityErrors(t *testing.T) {
	k, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	k.SetParams(ctx, types.NewParams(defaultSwapRules))

	creator := sample.AccAddress()
	initialStakeSupply := sdk.NewCoin(srcToken, math.NewInt(100))
	creatorBalance := sdk.NewCoins(initialStakeSupply)
	require.NoError(t, k.BankKeeper.MintCoins(ctx, minttypes.ModuleName, creatorBalance))
	require.NoError(t, k.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, sdk.MustAccAddressFromBech32(creator), creatorBalance))

	recipient := sample.AccAddress()

	t.Run("invalid creator address", func(t *testing.T) {
		_, err := srv.PartialSend(wctx, &types.MsgPartialSend{
			Creator:   "invalid_address",
			Recipient: recipient,
			Amount:    sdk.NewCoin(srcToken, math.NewInt(20)),
		})
		require.ErrorContains(t, err, "invalid creator address")
	})

	t.Run("not enough balance", func(t *testing.T) {
		_, err := srv.PartialSend(wctx, &types.MsgPartialSend{
			Creator:   sample.AccAddress(),
			Recipient: recipient,
			Amount:    sdk.NewCoin(srcToken, math.NewInt(20)),
		})
		require.ErrorContains(t, err, "insufficient balance")
	})

	t.Run("invalid recipient address", func(t *testing.T) {
		_, err := srv.PartialSend(wctx, &types.MsgPartialSend{
			Creator:   creator,
			Recipient: "invalid_address",
			Amount:    sdk.NewCoin(srcToken, math.NewInt(20)),
		})
		require.ErrorContains(t, err, "invalid recipient address")
	})

	k.SetParams(ctx, types.NewParams([]types.SwapRule{
		{
			SrcDenom:            "new_token",
			DstDenom:            dstToken,
			Ratio:               math.LegacyNewDecWithPrec(50, 2),
			MinSupplyRatioLimit: math.LegacyNewDecWithPrec(6667, 4),
		},
	}))

	t.Run("swap rule not found", func(t *testing.T) {
		_, err := srv.PartialSend(wctx, &types.MsgPartialSend{
			Creator:   creator,
			Recipient: recipient,
			Amount:    sdk.NewCoin(srcToken, math.NewInt(20)),
		})
		require.ErrorContains(t, err, "swap rule not found for source token")
	})
}

func TestPartialSendMsgServer_RatioFailure(t *testing.T) {
	k, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	k.SetParams(ctx, types.NewParams(defaultSwapRules))

	creator := sample.AccAddress()
	initialStakeSupply := sdk.NewCoin(srcToken, math.NewInt(120))
	creatorBalance := sdk.NewCoins(initialStakeSupply)
	require.NoError(t, k.BankKeeper.MintCoins(ctx, minttypes.ModuleName, creatorBalance))
	require.NoError(t, k.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, sdk.MustAccAddressFromBech32(creator), creatorBalance))

	initialStake1Supply := sdk.NewCoin(dstToken, math.NewInt(30))
	creator2Balance := sdk.NewCoins(initialStake1Supply)
	require.NoError(t, k.BankKeeper.MintCoins(ctx, minttypes.ModuleName, creator2Balance))
	require.NoError(t, k.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, sdk.MustAccAddressFromBech32(creator), creator2Balance))

	recipient := sample.AccAddress()
	for i := 0; i < 6; i++ {
		resp, err := srv.PartialSend(wctx, &types.MsgPartialSend{
			Creator:   creator,
			Recipient: recipient,
			Amount:    sdk.NewCoin(srcToken, math.NewInt(20)),
		})
		if i < 5 {
			require.NoError(t, err)
			require.Equal(t, i, int(resp.Id))
		} else {
			require.ErrorContains(t, err, "ratio 0.666666666666666667 is below the minimum supply ratio limit 0.666700000000000000")
		}
	}
}

func TestPartialSendMsgServer_AllBurn(t *testing.T) {
	k, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	k.SetParams(ctx, types.NewParams([]types.SwapRule{
		{
			SrcDenom:            srcToken,
			DstDenom:            dstToken,
			Ratio:               math.LegacyNewDecWithPrec(100, 2),
			MinSupplyRatioLimit: math.LegacyZeroDec(),
		},
	}))

	creator := sample.AccAddress()
	// initialStakeSupply := sdk.NewCoin(srcToken, math.NewInt(200))
	creatorBalance := sdk.NewCoins(sdk.NewCoin(srcToken, math.NewInt(100)))
	require.NoError(t, k.BankKeeper.MintCoins(ctx, minttypes.ModuleName, creatorBalance))
	require.NoError(t, k.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, sdk.MustAccAddressFromBech32(creator), creatorBalance))

	creator2Balance := sdk.NewCoins(sdk.NewCoin(srcToken, math.NewInt(100)))
	require.NoError(t, k.BankKeeper.MintCoins(ctx, minttypes.ModuleName, creator2Balance))
	require.NoError(t, k.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, sdk.MustAccAddressFromBech32(creator), creator2Balance))

	recipient := sample.AccAddress()

	t.Run("successful one-shot burn", func(t *testing.T) {
		_, err := srv.PartialSend(wctx, &types.MsgPartialSend{
			Creator:   creator,
			Recipient: recipient,
			Amount:    sdk.NewCoin(srcToken, math.NewInt(100)),
		})
		require.NoError(t, err)
	})
}

func TestPartialSendMsgServer_AllSwap(t *testing.T) {
	k, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	k.SetParams(ctx, types.NewParams([]types.SwapRule{
		{
			SrcDenom:            srcToken,
			DstDenom:            dstToken,
			Ratio:               math.LegacyZeroDec(),
			MinSupplyRatioLimit: math.LegacyZeroDec(),
		},
	}))

	creator := sample.AccAddress()
	initialStakeSupply := sdk.NewCoin(srcToken, math.NewInt(200))
	creatorBalance := sdk.NewCoins(initialStakeSupply)
	require.NoError(t, k.BankKeeper.MintCoins(ctx, minttypes.ModuleName, creatorBalance))
	require.NoError(t, k.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, sdk.MustAccAddressFromBech32(creator), creatorBalance))

	recipient := sample.AccAddress()

	t.Run("successful one-shot swap", func(t *testing.T) {
		_, err := srv.PartialSend(wctx, &types.MsgPartialSend{
			Creator:   creator,
			Recipient: recipient,
			Amount:    sdk.NewCoin(srcToken, initialStakeSupply.Amount),
		})
		require.NoError(t, err)

		stakeSupply := k.BankKeeper.GetSupply(ctx, srcToken)
		require.Equal(t, initialStakeSupply, stakeSupply)

		stake2Supply := k.BankKeeper.GetSupply(ctx, dstToken)
		require.Equal(t, sdk.NewCoin(dstToken, math.ZeroInt()), stake2Supply)

		recipientBalance := k.BankKeeper.SpendableCoins(ctx, sdk.MustAccAddressFromBech32(recipient))
		require.Equal(t, creatorBalance.AmountOf(srcToken), recipientBalance.AmountOf(srcToken))
		require.Equal(t, math.ZeroInt(), recipientBalance.AmountOf(dstToken))
	})
}
