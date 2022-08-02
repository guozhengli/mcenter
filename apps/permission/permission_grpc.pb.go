// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: apps/permission/pb/permission.proto

package permission

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
	QueryPermission(ctx context.Context, in *QueryPermissionRequest, opts ...grpc.CallOption) (*PermissionSet, error)
	DescribePermission(ctx context.Context, in *DescribePermissionRequest, opts ...grpc.CallOption) (*Permission, error)
	AddPermissionToRole(ctx context.Context, in *AddPermissionToRoleRequest, opts ...grpc.CallOption) (*PermissionSet, error)
	RemovePermissionFromRole(ctx context.Context, in *RemovePermissionFromRoleRequest, opts ...grpc.CallOption) (*PermissionSet, error)
	UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*Permission, error)
	CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*Permission, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) QueryPermission(ctx context.Context, in *QueryPermissionRequest, opts ...grpc.CallOption) (*PermissionSet, error) {
	out := new(PermissionSet)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.permission.RPC/QueryPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DescribePermission(ctx context.Context, in *DescribePermissionRequest, opts ...grpc.CallOption) (*Permission, error) {
	out := new(Permission)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.permission.RPC/DescribePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) AddPermissionToRole(ctx context.Context, in *AddPermissionToRoleRequest, opts ...grpc.CallOption) (*PermissionSet, error) {
	out := new(PermissionSet)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.permission.RPC/AddPermissionToRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) RemovePermissionFromRole(ctx context.Context, in *RemovePermissionFromRoleRequest, opts ...grpc.CallOption) (*PermissionSet, error) {
	out := new(PermissionSet)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.permission.RPC/RemovePermissionFromRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*Permission, error) {
	out := new(Permission)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.permission.RPC/UpdatePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*Permission, error) {
	out := new(Permission)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.permission.RPC/CheckPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	QueryPermission(context.Context, *QueryPermissionRequest) (*PermissionSet, error)
	DescribePermission(context.Context, *DescribePermissionRequest) (*Permission, error)
	AddPermissionToRole(context.Context, *AddPermissionToRoleRequest) (*PermissionSet, error)
	RemovePermissionFromRole(context.Context, *RemovePermissionFromRoleRequest) (*PermissionSet, error)
	UpdatePermission(context.Context, *UpdatePermissionRequest) (*Permission, error)
	CheckPermission(context.Context, *CheckPermissionRequest) (*Permission, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) QueryPermission(context.Context, *QueryPermissionRequest) (*PermissionSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPermission not implemented")
}
func (UnimplementedRPCServer) DescribePermission(context.Context, *DescribePermissionRequest) (*Permission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePermission not implemented")
}
func (UnimplementedRPCServer) AddPermissionToRole(context.Context, *AddPermissionToRoleRequest) (*PermissionSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPermissionToRole not implemented")
}
func (UnimplementedRPCServer) RemovePermissionFromRole(context.Context, *RemovePermissionFromRoleRequest) (*PermissionSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePermissionFromRole not implemented")
}
func (UnimplementedRPCServer) UpdatePermission(context.Context, *UpdatePermissionRequest) (*Permission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePermission not implemented")
}
func (UnimplementedRPCServer) CheckPermission(context.Context, *CheckPermissionRequest) (*Permission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermission not implemented")
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

func _RPC_QueryPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QueryPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.permission.RPC/QueryPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QueryPermission(ctx, req.(*QueryPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DescribePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DescribePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.permission.RPC/DescribePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DescribePermission(ctx, req.(*DescribePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_AddPermissionToRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPermissionToRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).AddPermissionToRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.permission.RPC/AddPermissionToRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).AddPermissionToRole(ctx, req.(*AddPermissionToRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_RemovePermissionFromRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePermissionFromRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).RemovePermissionFromRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.permission.RPC/RemovePermissionFromRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).RemovePermissionFromRole(ctx, req.(*RemovePermissionFromRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_UpdatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).UpdatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.permission.RPC/UpdatePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).UpdatePermission(ctx, req.(*UpdatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_CheckPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).CheckPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.permission.RPC/CheckPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).CheckPermission(ctx, req.(*CheckPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.mcenter.permission.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryPermission",
			Handler:    _RPC_QueryPermission_Handler,
		},
		{
			MethodName: "DescribePermission",
			Handler:    _RPC_DescribePermission_Handler,
		},
		{
			MethodName: "AddPermissionToRole",
			Handler:    _RPC_AddPermissionToRole_Handler,
		},
		{
			MethodName: "RemovePermissionFromRole",
			Handler:    _RPC_RemovePermissionFromRole_Handler,
		},
		{
			MethodName: "UpdatePermission",
			Handler:    _RPC_UpdatePermission_Handler,
		},
		{
			MethodName: "CheckPermission",
			Handler:    _RPC_CheckPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/permission/pb/permission.proto",
}
