package types

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	DefaultRatio               = math.LegacyNewDecWithPrec(50, 2)
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
	if len(p.SwapRules) == 0 {
		return errorsmod.Wrap(ErrInvalidParams, "SwapRules cannot be empty")
	}

	seenDenoms := make(map[string]struct{})
	for _, rule := range p.SwapRules {
		srcCoin := sdk.NewCoin(rule.SrcDenom, math.ZeroInt())
		err := srcCoin.Validate()
		if err != nil {
			return errorsmod.Wrap(ErrInvalidParams, "invalid src coin")
		}

		dstCoin := sdk.NewCoin(rule.DstDenom, math.ZeroInt())
		err = dstCoin.Validate()
		if err != nil {
			return errorsmod.Wrap(ErrInvalidParams, "invalid dst coin")
		}

		if rule.Ratio.LT(math.LegacyZeroDec()) || rule.Ratio.GT(math.LegacyOneDec()) {
			return errorsmod.Wrap(ErrInvalidParams, "ratio must be between 0 and 1")
		}
		if rule.MinSupplyRatioLimit.LT(math.LegacyZeroDec()) || rule.MinSupplyRatioLimit.GT(math.LegacyOneDec()) {
			return errorsmod.Wrap(ErrInvalidParams, "min supply ratio limit must be between 0 and 1")
		}

		if _, exists := seenDenoms[rule.SrcDenom]; exists {
			return errorsmod.Wrap(ErrInvalidParams, "duplicate src and dst token pair found")
		}
		seenDenoms[rule.SrcDenom] = struct{}{}
	}

	return nil
}

func (p *Params) GetSwapRule(denom string) (SwapRule, bool) {
	for _, rule := range p.SwapRules {
		if rule.SrcDenom == denom {
			return rule, true
		}
	}

	return SwapRule{}, false
}
