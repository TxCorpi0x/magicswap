package types

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// SwapRule defines the rule for swapping tokens
func (s *SwapRule) GetCoinSplit(coin sdk.Coin, selfSwap bool) (sdk.Coin, sdk.Coin, sdk.Coin) {
	if selfSwap {
		// if it is a self swap, the entire amount will be burned and swapped
		return sdk.NewCoin(s.SrcDenom, math.ZeroInt()), coin, sdk.NewCoin(s.DstDenom, coin.Amount)
	}
	// the amount will be converted to decimal for the further calculation,
	// with ratio equal to 50%. Odd values will be rounded, so it is possible for send and swap
	// to have different amounts to cover entire amount of source token input.
	amountDec := math.LegacyNewDecFromInt(coin.Amount)

	// the same amount will be used for both burning and swapping
	amountToBurnAndSwap := s.Ratio.Mul(amountDec).RoundInt()
	// the only difference is the token denom.
	coinsToBurn := sdk.NewCoin(s.SrcDenom, amountToBurnAndSwap)
	coinsToSwap := sdk.NewCoin(s.DstDenom, amountToBurnAndSwap)

	// the remaining amount will be sent to the recipient with the same source denom.
	amountToSend := coin.Amount.Sub(amountToBurnAndSwap)
	coinsToSend := sdk.NewCoin(s.SrcDenom, amountToSend)
	return coinsToSend, coinsToBurn, coinsToSwap
}

func (s *SwapRule) Validate(newSrcTokenSupply, newDstTokenSupply sdk.Coin) error {
	// this is a new token, no need to check the ratio and prevent division by zero.
	if newDstTokenSupply.Amount.IsZero() {
		return nil
	}

	// check if the supply ratio is above the minimum supply ratio limit
	// this ensures that the swap will not burn all of the source token supply.
	ratio := math.LegacyNewDecFromInt(newSrcTokenSupply.Amount).Quo(math.LegacyNewDecFromInt(newDstTokenSupply.Amount))
	if ratio.LT(s.MinSupplyRatioLimit) {
		return errorsmod.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"ratio %s is below the minimum supply ratio limit %s",
			ratio.String(),
			s.MinSupplyRatioLimit.String(),
		)
	}

	return nil
}
