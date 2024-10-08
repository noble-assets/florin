// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: florin/blacklist/v1/tx.proto

package blacklistv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Msg_AcceptOwnership_FullMethodName    = "/florin.blacklist.v1.Msg/AcceptOwnership"
	Msg_AddAdminAccount_FullMethodName    = "/florin.blacklist.v1.Msg/AddAdminAccount"
	Msg_Ban_FullMethodName                = "/florin.blacklist.v1.Msg/Ban"
	Msg_RemoveAdminAccount_FullMethodName = "/florin.blacklist.v1.Msg/RemoveAdminAccount"
	Msg_TransferOwnership_FullMethodName  = "/florin.blacklist.v1.Msg/TransferOwnership"
	Msg_Unban_FullMethodName              = "/florin.blacklist.v1.Msg/Unban"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	AcceptOwnership(ctx context.Context, in *MsgAcceptOwnership, opts ...grpc.CallOption) (*MsgAcceptOwnershipResponse, error)
	AddAdminAccount(ctx context.Context, in *MsgAddAdminAccount, opts ...grpc.CallOption) (*MsgAddAdminAccountResponse, error)
	Ban(ctx context.Context, in *MsgBan, opts ...grpc.CallOption) (*MsgBanResponse, error)
	RemoveAdminAccount(ctx context.Context, in *MsgRemoveAdminAccount, opts ...grpc.CallOption) (*MsgRemoveAdminAccountResponse, error)
	TransferOwnership(ctx context.Context, in *MsgTransferOwnership, opts ...grpc.CallOption) (*MsgTransferOwnershipResponse, error)
	Unban(ctx context.Context, in *MsgUnban, opts ...grpc.CallOption) (*MsgUnbanResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AcceptOwnership(ctx context.Context, in *MsgAcceptOwnership, opts ...grpc.CallOption) (*MsgAcceptOwnershipResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgAcceptOwnershipResponse)
	err := c.cc.Invoke(ctx, Msg_AcceptOwnership_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddAdminAccount(ctx context.Context, in *MsgAddAdminAccount, opts ...grpc.CallOption) (*MsgAddAdminAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgAddAdminAccountResponse)
	err := c.cc.Invoke(ctx, Msg_AddAdminAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Ban(ctx context.Context, in *MsgBan, opts ...grpc.CallOption) (*MsgBanResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgBanResponse)
	err := c.cc.Invoke(ctx, Msg_Ban_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveAdminAccount(ctx context.Context, in *MsgRemoveAdminAccount, opts ...grpc.CallOption) (*MsgRemoveAdminAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgRemoveAdminAccountResponse)
	err := c.cc.Invoke(ctx, Msg_RemoveAdminAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TransferOwnership(ctx context.Context, in *MsgTransferOwnership, opts ...grpc.CallOption) (*MsgTransferOwnershipResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgTransferOwnershipResponse)
	err := c.cc.Invoke(ctx, Msg_TransferOwnership_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Unban(ctx context.Context, in *MsgUnban, opts ...grpc.CallOption) (*MsgUnbanResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgUnbanResponse)
	err := c.cc.Invoke(ctx, Msg_Unban_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility.
type MsgServer interface {
	AcceptOwnership(context.Context, *MsgAcceptOwnership) (*MsgAcceptOwnershipResponse, error)
	AddAdminAccount(context.Context, *MsgAddAdminAccount) (*MsgAddAdminAccountResponse, error)
	Ban(context.Context, *MsgBan) (*MsgBanResponse, error)
	RemoveAdminAccount(context.Context, *MsgRemoveAdminAccount) (*MsgRemoveAdminAccountResponse, error)
	TransferOwnership(context.Context, *MsgTransferOwnership) (*MsgTransferOwnershipResponse, error)
	Unban(context.Context, *MsgUnban) (*MsgUnbanResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMsgServer struct{}

func (UnimplementedMsgServer) AcceptOwnership(context.Context, *MsgAcceptOwnership) (*MsgAcceptOwnershipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptOwnership not implemented")
}
func (UnimplementedMsgServer) AddAdminAccount(context.Context, *MsgAddAdminAccount) (*MsgAddAdminAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAdminAccount not implemented")
}
func (UnimplementedMsgServer) Ban(context.Context, *MsgBan) (*MsgBanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ban not implemented")
}
func (UnimplementedMsgServer) RemoveAdminAccount(context.Context, *MsgRemoveAdminAccount) (*MsgRemoveAdminAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAdminAccount not implemented")
}
func (UnimplementedMsgServer) TransferOwnership(context.Context, *MsgTransferOwnership) (*MsgTransferOwnershipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferOwnership not implemented")
}
func (UnimplementedMsgServer) Unban(context.Context, *MsgUnban) (*MsgUnbanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unban not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}
func (UnimplementedMsgServer) testEmbeddedByValue()             {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	// If the following call pancis, it indicates UnimplementedMsgServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_AcceptOwnership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAcceptOwnership)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AcceptOwnership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AcceptOwnership_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AcceptOwnership(ctx, req.(*MsgAcceptOwnership))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddAdminAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddAdminAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddAdminAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AddAdminAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddAdminAccount(ctx, req.(*MsgAddAdminAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Ban_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBan)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Ban(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Ban_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Ban(ctx, req.(*MsgBan))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveAdminAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveAdminAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveAdminAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RemoveAdminAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveAdminAccount(ctx, req.(*MsgRemoveAdminAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TransferOwnership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransferOwnership)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransferOwnership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_TransferOwnership_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransferOwnership(ctx, req.(*MsgTransferOwnership))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Unban_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnban)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Unban(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Unban_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Unban(ctx, req.(*MsgUnban))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "florin.blacklist.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AcceptOwnership",
			Handler:    _Msg_AcceptOwnership_Handler,
		},
		{
			MethodName: "AddAdminAccount",
			Handler:    _Msg_AddAdminAccount_Handler,
		},
		{
			MethodName: "Ban",
			Handler:    _Msg_Ban_Handler,
		},
		{
			MethodName: "RemoveAdminAccount",
			Handler:    _Msg_RemoveAdminAccount_Handler,
		},
		{
			MethodName: "TransferOwnership",
			Handler:    _Msg_TransferOwnership_Handler,
		},
		{
			MethodName: "Unban",
			Handler:    _Msg_Unban_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "florin/blacklist/v1/tx.proto",
}
