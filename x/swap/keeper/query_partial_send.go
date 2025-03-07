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

func (k Keeper) PartialSendAll(ctx context.Context, req *types.QueryAllPartialSendRequest) (*types.QueryAllPartialSendResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var partialSends []types.PartialSend

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	partialSendStore := prefix.NewStore(store, types.KeyPrefix(types.PartialSendKey))

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

func (k Keeper) PartialSend(ctx context.Context, req *types.QueryGetPartialSendRequest) (*types.QueryGetPartialSendResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	partialSend, found := k.GetPartialSend(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetPartialSendResponse{PartialSend: partialSend}, nil
}
