package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/TxCorpi0x/magicswap/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// PartialSend handles a message to partially send coins from the creator to the recipient.
func (k msgServer) PartialSend(goCtx context.Context, msg *types.MsgPartialSend) (*types.MsgPartialSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// do sanity checks  on the message input values

	// check if the creator address is valid
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", msg.Creator)
	}

	// check if the recipient address is valid
	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address: %s", msg.Recipient)
	}

	// check if the input amount is positive
	if !msg.Amount.IsPositive() {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "amount cannot be zero or negative")
	}

	// check if the creator balance is sufficient to cover all of the input amount
	if k.BankKeeper.SpendableCoins(ctx, creator).AmountOf(msg.Amount.Denom).LT(msg.Amount.Amount) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "insufficient balance: %s", msg.Amount.String())
	}

	// get the swap rule for the source token
	params := k.GetParams(ctx)
	swapRule, found := params.GetSwapRule(msg.Amount.Denom)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "swap rule not found for source token: %s", msg.Amount.Denom)
	}

	// check if the recipient is the same as the creator
	isSelfSwap := creator.Equals(recipient)

	coinsToSend, coinsToBurn, coinsToSwap := swapRule.GetCoinSplit(msg.Amount, isSelfSwap)

	// send coins to the recipient
	if coinsToSend.IsPositive() {
		if err := k.BankKeeper.SendCoins(ctx, creator, recipient, sdk.NewCoins(coinsToSend)); err != nil {
			return nil, errorsmod.Wrapf(err, "failed to send coins from %s to %s", msg.Creator, msg.Recipient)
		}
	}

	// burn source coins from the creator account balance.
	if coinsToBurn.IsPositive() {
		if err := k.burn(ctx, creator, sdk.NewCoins(coinsToBurn)); err != nil {
			return nil, errorsmod.Wrapf(err, "failed to burn coins from %s", creator.String())
		}
	}

	// mint destination coins to the recipient account balance.
	if coinsToSwap.IsPositive() {
		if err := k.mint(ctx, recipient, sdk.NewCoins(coinsToSwap)); err != nil {
			return nil, errorsmod.Wrapf(err, "failed to mint coins to %s", recipient.String())
		}
	}

	// validate the swap rule after the flow completion
	// to ensure the swap rule is still valid after processing.
	newSrcTokenSupply := k.BankKeeper.GetSupply(ctx, swapRule.SrcDenom)
	newDstTokenSupply := k.BankKeeper.GetSupply(ctx, swapRule.DstDenom)

	if err := swapRule.Validate(newSrcTokenSupply, newDstTokenSupply); err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "swap rule validation failed: %v", err)
	}

	// create a new partial send record
	partialSend := types.PartialSend{
		Creator:       msg.Creator,
		Recipient:     msg.Recipient,
		BurntAmount:   coinsToBurn,
		SentAmount:    coinsToSend,
		SwappedAmount: coinsToSwap,
	}

	// append the partial send record to the store
	id := k.AppendPartialSend(ctx, partialSend)

	return &types.MsgPartialSendResponse{Id: id}, nil
}

// burn transfers coins from the specified account to the module account and then burns them.
// to be able to burn tokens, the amount should be sent to the module account first.
// this is because the module account is the configured in the app permissions that can burn tokens.
func (k Keeper) burn(ctx sdk.Context, account sdk.AccAddress, coins sdk.Coins) error {
	// Send coins from the account to the module account
	if err := k.BankKeeper.SendCoinsFromAccountToModule(ctx, account, types.ModuleName, coins); err != nil {
		return errorsmod.Wrapf(err, "could not send coins from account %s to module %s", account.String(), types.ModuleName)
	}

	// Burn the coins from the module account
	if err := k.BankKeeper.BurnCoins(ctx, types.ModuleName, coins); err != nil {
		return errorsmod.Wrapf(err, "could not burn %s for the module %s", coins.String(), types.ModuleName)
	}

	return nil
}

// mint mints new coins to the module account and then sends them to the specified account.
// the module account has the minting permissions to mint new coins, this is added to the app permissions.
// then after minting, the coins are sent to the specified account.
func (k Keeper) mint(ctx sdk.Context, account sdk.AccAddress, coins sdk.Coins) error {
	// Mint new coins to the module account
	if err := k.BankKeeper.MintCoins(ctx, types.ModuleName, coins); err != nil {
		return errorsmod.Wrapf(err, "could not mint %s for the module %s", coins.String(), types.ModuleName)
	}

	// Send the minted coins from the module account to the specified account
	if err := k.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, account, coins); err != nil {
		return errorsmod.Wrapf(err, "could not send minted coins from module %s to account %s", types.ModuleName, account.String())
	}

	return nil
}
