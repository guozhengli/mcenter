// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: apps/gateway/pb/rpc.proto

package gateway

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

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	QueryGateway(ctx context.Context, in *QueryGatewayRequest, opts ...grpc.CallOption) (*GatewaySet, error)
	DescribeGateway(ctx context.Context, in *DescribeGatewayRequest, opts ...grpc.CallOption) (*Gateway, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) QueryGateway(ctx context.Context, in *QueryGatewayRequest, opts ...grpc.CallOption) (*GatewaySet, error) {
	out := new(GatewaySet)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.gateway.RPC/QueryGateway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DescribeGateway(ctx context.Context, in *DescribeGatewayRequest, opts ...grpc.CallOption) (*Gateway, error) {
	out := new(Gateway)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.gateway.RPC/DescribeGateway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	QueryGateway(context.Context, *QueryGatewayRequest) (*GatewaySet, error)
	DescribeGateway(context.Context, *DescribeGatewayRequest) (*Gateway, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) QueryGateway(context.Context, *QueryGatewayRequest) (*GatewaySet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryGateway not implemented")
}
func (UnimplementedRPCServer) DescribeGateway(context.Context, *DescribeGatewayRequest) (*Gateway, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeGateway not implemented")
}
func (UnimplementedRPCServer) mustEmbedUnimplementedRPCServer() {}

// UnsafeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServer will
// result in compilation errors.
type UnsafeRPCServer interface {
	mustEmbedUnimplementedRPCServer()
}

func RegisterRPCServer(s grpc.ServiceRegistrar, srv RPCServer) {
	s.RegisterService(&RPC_ServiceDesc, srv)
}

func _RPC_QueryGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QueryGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.gateway.RPC/QueryGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QueryGateway(ctx, req.(*QueryGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DescribeGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DescribeGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.gateway.RPC/DescribeGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DescribeGateway(ctx, req.(*DescribeGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.mcenter.gateway.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryGateway",
			Handler:    _RPC_QueryGateway_Handler,
		},
		{
			MethodName: "DescribeGateway",
			Handler:    _RPC_DescribeGateway_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/gateway/pb/rpc.proto",
}
