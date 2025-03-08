package types

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	// DefaultRation is the genesis default value for the swap ratio
	// 50% of the source token will be burned from source account and swapped to the destination
	DefaultRatio = math.LegacyNewDecWithPrec(50, 2)
	// DefaultMinSupplyRatioLimit is the genesis default value for the min supply ratio limit
	// 66.67% of the source token supply must remain after the swap
	DefaultMinSupplyRatioLimit = math.LegacyNewDecWithPrec(6667, 4)
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(swapRules []SwapRule) Params {
	return Params{swapRules}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams([]SwapRule{
		{
			SrcDenom:            "stake",
			DstDenom:            "stake1",
			Ratio:               DefaultRatio,
			MinSupplyRatioLimit: DefaultMinSupplyRatioLimit,
		},
	})
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	// ensure that the swap rules are not empty
	if len(p.SwapRules) == 0 {
		return errorsmod.Wrap(ErrInvalidParams, "SwapRules cannot be empty")
	}

	seenDenoms := make(map[string]struct{})
	for _, rule := range p.SwapRules {
		// validate the source coin
		srcCoin := sdk.NewCoin(rule.SrcDenom, math.ZeroInt())
		err := srcCoin.Validate()
		if err != nil {
			return errorsmod.Wrap(ErrInvalidParams, "invalid src coin")
		}

		// validate the destination coin
		dstCoin := sdk.NewCoin(rule.DstDenom, math.ZeroInt())
		err = dstCoin.Validate()
		if err != nil {
			return errorsmod.Wrap(ErrInvalidParams, "invalid dst coin")
		}

		// the module supports both all-send and all-swap
		// this means if the ration is set as 1, the source token will be swapped entirely
		// and nothing will be sent to the destination account directly.
		// if the ratio is set as 0, the source token will be sent entirely to the destination account
		// and nothing will be swapped.
		if rule.Ratio.LT(math.LegacyZeroDec()) || rule.Ratio.GT(math.LegacyOneDec()) {
			return errorsmod.Wrap(ErrInvalidParams, "ratio must be between 0 and 1")
		}

		// the min supply ratio limit is the minimum ratio of the source token supply
		// that must remain after the swap as well.
		// this is to ensure that the source token supply does not get depleted entirely.
		if rule.MinSupplyRatioLimit.LT(math.LegacyZeroDec()) || rule.MinSupplyRatioLimit.GT(math.LegacyOneDec()) {
			return errorsmod.Wrap(ErrInvalidParams, "min supply ratio limit must be between 0 and 1")
		}

		// ensure that the src and dst token pair is unique
		if _, exists := seenDenoms[rule.SrcDenom]; exists {
			return errorsmod.Wrap(ErrInvalidParams, "duplicate src and dst token pair found")
		}
		seenDenoms[rule.SrcDenom] = struct{}{}
	}

	return nil
}

// GetSwapRule returns the swap rule for the given denom
// if the swap rule does not exist, it returns false
func (p *Params) GetSwapRule(denom string) (SwapRule, bool) {
	for _, rule := range p.SwapRules {
		if rule.SrcDenom == denom {
			return rule, true
		}
	}

	return SwapRule{}, false
}
