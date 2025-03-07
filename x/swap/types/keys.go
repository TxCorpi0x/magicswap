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
	PartialSendKey      = []byte{0x00}
	PartialSendCountKey = []byte{0x01}
)

func GetPartialSendPrefix(creator string) []byte {
	return append(PartialSendKey, []byte(creator)...)
}

func GetPartialSendKey(creator string, id uint64) []byte {
	return append(GetPartialSendPrefix(creator), Uint64ToBytes(id)...)
}

func Uint64ToBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
