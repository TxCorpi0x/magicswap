// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: magicswap/swap/query.proto

package swap

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Query_Params_FullMethodName               = "/magicswap.swap.Query/Params"
	Query_PartialSend_FullMethodName          = "/magicswap.swap.Query/PartialSend"
	Query_PartialSendByCreator_FullMethodName = "/magicswap.swap.Query/PartialSendByCreator"
	Query_PartialSendAll_FullMethodName       = "/magicswap.swap.Query/PartialSendAll"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of PartialSend items.
	PartialSend(ctx context.Context, in *QueryGetPartialSendRequest, opts ...grpc.CallOption) (*QueryGetPartialSendResponse, error)
	// Queries PartialSend items of an address.
	PartialSendByCreator(ctx context.Context, in *QueryGetPartialSendByCreatorRequest, opts ...grpc.CallOption) (*QueryGetPartialSendByCreatorResponse, error)
	PartialSendAll(ctx context.Context, in *QueryAllPartialSendRequest, opts ...grpc.CallOption) (*QueryAllPartialSendResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PartialSend(ctx context.Context, in *QueryGetPartialSendRequest, opts ...grpc.CallOption) (*QueryGetPartialSendResponse, error) {
	out := new(QueryGetPartialSendResponse)
	err := c.cc.Invoke(ctx, Query_PartialSend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PartialSendByCreator(ctx context.Context, in *QueryGetPartialSendByCreatorRequest, opts ...grpc.CallOption) (*QueryGetPartialSendByCreatorResponse, error) {
	out := new(QueryGetPartialSendByCreatorResponse)
	err := c.cc.Invoke(ctx, Query_PartialSendByCreator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PartialSendAll(ctx context.Context, in *QueryAllPartialSendRequest, opts ...grpc.CallOption) (*QueryAllPartialSendResponse, error) {
	out := new(QueryAllPartialSendResponse)
	err := c.cc.Invoke(ctx, Query_PartialSendAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of PartialSend items.
	PartialSend(context.Context, *QueryGetPartialSendRequest) (*QueryGetPartialSendResponse, error)
	// Queries PartialSend items of an address.
	PartialSendByCreator(context.Context, *QueryGetPartialSendByCreatorRequest) (*QueryGetPartialSendByCreatorResponse, error)
	PartialSendAll(context.Context, *QueryAllPartialSendRequest) (*QueryAllPartialSendResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) PartialSend(context.Context, *QueryGetPartialSendRequest) (*QueryGetPartialSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PartialSend not implemented")
}
func (UnimplementedQueryServer) PartialSendByCreator(context.Context, *QueryGetPartialSendByCreatorRequest) (*QueryGetPartialSendByCreatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PartialSendByCreator not implemented")
}
func (UnimplementedQueryServer) PartialSendAll(context.Context, *QueryAllPartialSendRequest) (*QueryAllPartialSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PartialSendAll not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PartialSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetPartialSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PartialSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PartialSend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PartialSend(ctx, req.(*QueryGetPartialSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PartialSendByCreator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetPartialSendByCreatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PartialSendByCreator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PartialSendByCreator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PartialSendByCreator(ctx, req.(*QueryGetPartialSendByCreatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PartialSendAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllPartialSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PartialSendAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PartialSendAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PartialSendAll(ctx, req.(*QueryAllPartialSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magicswap.swap.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "PartialSend",
			Handler:    _Query_PartialSend_Handler,
		},
		{
			MethodName: "PartialSendByCreator",
			Handler:    _Query_PartialSendByCreator_Handler,
		},
		{
			MethodName: "PartialSendAll",
			Handler:    _Query_PartialSendAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "magicswap/swap/query.proto",
}
