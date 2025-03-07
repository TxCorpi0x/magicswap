package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/TxCorpi0x/magicswap/x/swap/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetPartialSendCount get the total number of partialSend
func (k Keeper) GetPartialSendCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.PartialSendCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPartialSendCount set the total number of partialSend
func (k Keeper) SetPartialSendCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.PartialSendCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPartialSend appends a partialSend in the store with a new id and update the count
func (k Keeper) AppendPartialSend(
	ctx context.Context,
	partialSend types.PartialSend,
) uint64 {
	// Create the partialSend
	count := k.GetPartialSendCount(ctx)

	// Set the ID of the appended value
	partialSend.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PartialSendKey))
	appendedValue := k.cdc.MustMarshal(&partialSend)
	store.Set(GetPartialSendIDBytes(partialSend.Id), appendedValue)

	// Update partialSend count
	k.SetPartialSendCount(ctx, count+1)

	return count
}

// SetPartialSend set a specific partialSend in the store
func (k Keeper) SetPartialSend(ctx context.Context, partialSend types.PartialSend) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PartialSendKey))
	b := k.cdc.MustMarshal(&partialSend)
	store.Set(GetPartialSendIDBytes(partialSend.Id), b)
}

// GetPartialSend returns a partialSend from its id
func (k Keeper) GetPartialSend(ctx context.Context, id uint64) (val types.PartialSend, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PartialSendKey))
	b := store.Get(GetPartialSendIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllPartialSend returns all partialSend
func (k Keeper) GetAllPartialSend(ctx context.Context) (list []types.PartialSend) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PartialSendKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PartialSend
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPartialSendIDBytes returns the byte representation of the ID
func GetPartialSendIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.PartialSendKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
