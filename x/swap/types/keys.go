package types

import "encoding/binary"

const (
	// ModuleName defines the module name
	ModuleName = "swap"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_swap"
)

var (
	ParamsKey = []byte("p_swap")
)

var (
	// Prefix for partialSend
	PartialSendKey = []byte{0x00}
	// Prefix for partialSend count, this is used to create id sequence
	PartialSendCountKey = []byte{0x01}
)

// GetPartialSendPrefix returns the prefix for partialSend
func GetPartialSendPrefix(creator string) []byte {
	return append(PartialSendKey, []byte(creator)...)
}

// GetPartialSendKey returns the key of a partialSend with the given creator address and id
func GetPartialSendKey(creator string, id uint64) []byte {
	return append(GetPartialSendPrefix(creator), uint64ToBytes(id)...)
}

func uint64ToBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
