package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PartialSendList: []PartialSend{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in partialSend
	partialSendIdMap := make(map[uint64]bool)
	partialSendCount := gs.GetPartialSendCount()
	for _, elem := range gs.PartialSendList {
		if _, ok := partialSendIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for partialSend")
		}
		if elem.Id >= partialSendCount {
			return fmt.Errorf("partialSend id should be lower or equal than the last id")
		}
		partialSendIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
