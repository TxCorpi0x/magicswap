package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/TxCorpi0x/magicswap/x/swap/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PartialSendAll gets all the partialSend
func (k Keeper) PartialSendAll(ctx context.Context, req *types.QueryAllPartialSendRequest) (*types.QueryAllPartialSendResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var partialSends []types.PartialSend

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	partialSendStore := prefix.NewStore(store, types.PartialSendKey)

	pageRes, err := query.Paginate(partialSendStore, req.Pagination, func(key []byte, value []byte) error {
		var partialSend types.PartialSend
		if err := k.cdc.Unmarshal(value, &partialSend); err != nil {
			return err
		}

		partialSends = append(partialSends, partialSend)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPartialSendResponse{PartialSend: partialSends, Pagination: pageRes}, nil
}

// PartialSendByCreator gets partialSend by creator
func (k Keeper) PartialSendByCreator(ctx context.Context, req *types.QueryGetPartialSendByCreatorRequest) (*types.QueryGetPartialSendByCreatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var partialSends []types.PartialSend

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	partialSendStore := prefix.NewStore(store, types.PartialSendKey)
	filterredPartialSendStore := prefix.NewStore(partialSendStore, types.GetPartialSendPrefix(req.Creator))

	pageRes, err := query.Paginate(filterredPartialSendStore, req.Pagination, func(key []byte, value []byte) error {
		var partialSend types.PartialSend
		if err := k.cdc.Unmarshal(value, &partialSend); err != nil {
			return err
		}

		partialSends = append(partialSends, partialSend)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetPartialSendByCreatorResponse{PartialSend: partialSends, Pagination: pageRes}, nil
}

// PartialSend gets the partialSend by creator and id
func (k Keeper) PartialSend(ctx context.Context, req *types.QueryGetPartialSendRequest) (*types.QueryGetPartialSendResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	partialSend, found := k.GetPartialSend(ctx, req.Creator, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetPartialSendResponse{PartialSend: partialSend}, nil
}
