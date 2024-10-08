// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: florin/v2/query.proto

package florinv2

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
	Query_Authority_FullMethodName         = "/florin.v2.Query/Authority"
	Query_AllowedDenoms_FullMethodName     = "/florin.v2.Query/AllowedDenoms"
	Query_Owners_FullMethodName            = "/florin.v2.Query/Owners"
	Query_Owner_FullMethodName             = "/florin.v2.Query/Owner"
	Query_Systems_FullMethodName           = "/florin.v2.Query/Systems"
	Query_SystemsByDenom_FullMethodName    = "/florin.v2.Query/SystemsByDenom"
	Query_Admins_FullMethodName            = "/florin.v2.Query/Admins"
	Query_AdminsByDenom_FullMethodName     = "/florin.v2.Query/AdminsByDenom"
	Query_MaxMintAllowances_FullMethodName = "/florin.v2.Query/MaxMintAllowances"
	Query_MaxMintAllowance_FullMethodName  = "/florin.v2.Query/MaxMintAllowance"
	Query_MintAllowances_FullMethodName    = "/florin.v2.Query/MintAllowances"
	Query_MintAllowance_FullMethodName     = "/florin.v2.Query/MintAllowance"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	Authority(ctx context.Context, in *QueryAuthority, opts ...grpc.CallOption) (*QueryAuthorityResponse, error)
	AllowedDenoms(ctx context.Context, in *QueryAllowedDenoms, opts ...grpc.CallOption) (*QueryAllowedDenomsResponse, error)
	Owners(ctx context.Context, in *QueryOwners, opts ...grpc.CallOption) (*QueryOwnersResponse, error)
	Owner(ctx context.Context, in *QueryOwner, opts ...grpc.CallOption) (*QueryOwnerResponse, error)
	Systems(ctx context.Context, in *QuerySystems, opts ...grpc.CallOption) (*QuerySystemsResponse, error)
	SystemsByDenom(ctx context.Context, in *QuerySystemsByDenom, opts ...grpc.CallOption) (*QuerySystemsByDenomResponse, error)
	Admins(ctx context.Context, in *QueryAdmins, opts ...grpc.CallOption) (*QueryAdminsResponse, error)
	AdminsByDenom(ctx context.Context, in *QueryAdminsByDenom, opts ...grpc.CallOption) (*QueryAdminsByDenomResponse, error)
	MaxMintAllowances(ctx context.Context, in *QueryMaxMintAllowances, opts ...grpc.CallOption) (*QueryMaxMintAllowancesResponse, error)
	MaxMintAllowance(ctx context.Context, in *QueryMaxMintAllowance, opts ...grpc.CallOption) (*QueryMaxMintAllowanceResponse, error)
	MintAllowances(ctx context.Context, in *QueryMintAllowances, opts ...grpc.CallOption) (*QueryMintAllowancesResponse, error)
	MintAllowance(ctx context.Context, in *QueryMintAllowance, opts ...grpc.CallOption) (*QueryMintAllowanceResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Authority(ctx context.Context, in *QueryAuthority, opts ...grpc.CallOption) (*QueryAuthorityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryAuthorityResponse)
	err := c.cc.Invoke(ctx, Query_Authority_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllowedDenoms(ctx context.Context, in *QueryAllowedDenoms, opts ...grpc.CallOption) (*QueryAllowedDenomsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryAllowedDenomsResponse)
	err := c.cc.Invoke(ctx, Query_AllowedDenoms_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Owners(ctx context.Context, in *QueryOwners, opts ...grpc.CallOption) (*QueryOwnersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryOwnersResponse)
	err := c.cc.Invoke(ctx, Query_Owners_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Owner(ctx context.Context, in *QueryOwner, opts ...grpc.CallOption) (*QueryOwnerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryOwnerResponse)
	err := c.cc.Invoke(ctx, Query_Owner_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Systems(ctx context.Context, in *QuerySystems, opts ...grpc.CallOption) (*QuerySystemsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QuerySystemsResponse)
	err := c.cc.Invoke(ctx, Query_Systems_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SystemsByDenom(ctx context.Context, in *QuerySystemsByDenom, opts ...grpc.CallOption) (*QuerySystemsByDenomResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QuerySystemsByDenomResponse)
	err := c.cc.Invoke(ctx, Query_SystemsByDenom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Admins(ctx context.Context, in *QueryAdmins, opts ...grpc.CallOption) (*QueryAdminsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryAdminsResponse)
	err := c.cc.Invoke(ctx, Query_Admins_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AdminsByDenom(ctx context.Context, in *QueryAdminsByDenom, opts ...grpc.CallOption) (*QueryAdminsByDenomResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryAdminsByDenomResponse)
	err := c.cc.Invoke(ctx, Query_AdminsByDenom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) MaxMintAllowances(ctx context.Context, in *QueryMaxMintAllowances, opts ...grpc.CallOption) (*QueryMaxMintAllowancesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryMaxMintAllowancesResponse)
	err := c.cc.Invoke(ctx, Query_MaxMintAllowances_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) MaxMintAllowance(ctx context.Context, in *QueryMaxMintAllowance, opts ...grpc.CallOption) (*QueryMaxMintAllowanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryMaxMintAllowanceResponse)
	err := c.cc.Invoke(ctx, Query_MaxMintAllowance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) MintAllowances(ctx context.Context, in *QueryMintAllowances, opts ...grpc.CallOption) (*QueryMintAllowancesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryMintAllowancesResponse)
	err := c.cc.Invoke(ctx, Query_MintAllowances_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) MintAllowance(ctx context.Context, in *QueryMintAllowance, opts ...grpc.CallOption) (*QueryMintAllowanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryMintAllowanceResponse)
	err := c.cc.Invoke(ctx, Query_MintAllowance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility.
type QueryServer interface {
	Authority(context.Context, *QueryAuthority) (*QueryAuthorityResponse, error)
	AllowedDenoms(context.Context, *QueryAllowedDenoms) (*QueryAllowedDenomsResponse, error)
	Owners(context.Context, *QueryOwners) (*QueryOwnersResponse, error)
	Owner(context.Context, *QueryOwner) (*QueryOwnerResponse, error)
	Systems(context.Context, *QuerySystems) (*QuerySystemsResponse, error)
	SystemsByDenom(context.Context, *QuerySystemsByDenom) (*QuerySystemsByDenomResponse, error)
	Admins(context.Context, *QueryAdmins) (*QueryAdminsResponse, error)
	AdminsByDenom(context.Context, *QueryAdminsByDenom) (*QueryAdminsByDenomResponse, error)
	MaxMintAllowances(context.Context, *QueryMaxMintAllowances) (*QueryMaxMintAllowancesResponse, error)
	MaxMintAllowance(context.Context, *QueryMaxMintAllowance) (*QueryMaxMintAllowanceResponse, error)
	MintAllowances(context.Context, *QueryMintAllowances) (*QueryMintAllowancesResponse, error)
	MintAllowance(context.Context, *QueryMintAllowance) (*QueryMintAllowanceResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQueryServer struct{}

func (UnimplementedQueryServer) Authority(context.Context, *QueryAuthority) (*QueryAuthorityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authority not implemented")
}
func (UnimplementedQueryServer) AllowedDenoms(context.Context, *QueryAllowedDenoms) (*QueryAllowedDenomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowedDenoms not implemented")
}
func (UnimplementedQueryServer) Owners(context.Context, *QueryOwners) (*QueryOwnersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Owners not implemented")
}
func (UnimplementedQueryServer) Owner(context.Context, *QueryOwner) (*QueryOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Owner not implemented")
}
func (UnimplementedQueryServer) Systems(context.Context, *QuerySystems) (*QuerySystemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Systems not implemented")
}
func (UnimplementedQueryServer) SystemsByDenom(context.Context, *QuerySystemsByDenom) (*QuerySystemsByDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemsByDenom not implemented")
}
func (UnimplementedQueryServer) Admins(context.Context, *QueryAdmins) (*QueryAdminsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Admins not implemented")
}
func (UnimplementedQueryServer) AdminsByDenom(context.Context, *QueryAdminsByDenom) (*QueryAdminsByDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminsByDenom not implemented")
}
func (UnimplementedQueryServer) MaxMintAllowances(context.Context, *QueryMaxMintAllowances) (*QueryMaxMintAllowancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MaxMintAllowances not implemented")
}
func (UnimplementedQueryServer) MaxMintAllowance(context.Context, *QueryMaxMintAllowance) (*QueryMaxMintAllowanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MaxMintAllowance not implemented")
}
func (UnimplementedQueryServer) MintAllowances(context.Context, *QueryMintAllowances) (*QueryMintAllowancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintAllowances not implemented")
}
func (UnimplementedQueryServer) MintAllowance(context.Context, *QueryMintAllowance) (*QueryMintAllowanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintAllowance not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}
func (UnimplementedQueryServer) testEmbeddedByValue()               {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	// If the following call pancis, it indicates UnimplementedQueryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Authority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAuthority)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Authority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Authority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Authority(ctx, req.(*QueryAuthority))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllowedDenoms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllowedDenoms)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllowedDenoms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AllowedDenoms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllowedDenoms(ctx, req.(*QueryAllowedDenoms))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Owners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOwners)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Owners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Owners_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Owners(ctx, req.(*QueryOwners))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Owner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOwner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Owner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Owner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Owner(ctx, req.(*QueryOwner))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Systems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySystems)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Systems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Systems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Systems(ctx, req.(*QuerySystems))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SystemsByDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySystemsByDenom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SystemsByDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_SystemsByDenom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SystemsByDenom(ctx, req.(*QuerySystemsByDenom))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Admins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAdmins)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Admins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Admins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Admins(ctx, req.(*QueryAdmins))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AdminsByDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAdminsByDenom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AdminsByDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AdminsByDenom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AdminsByDenom(ctx, req.(*QueryAdminsByDenom))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_MaxMintAllowances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMaxMintAllowances)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MaxMintAllowances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_MaxMintAllowances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MaxMintAllowances(ctx, req.(*QueryMaxMintAllowances))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_MaxMintAllowance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMaxMintAllowance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MaxMintAllowance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_MaxMintAllowance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MaxMintAllowance(ctx, req.(*QueryMaxMintAllowance))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_MintAllowances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMintAllowances)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MintAllowances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_MintAllowances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MintAllowances(ctx, req.(*QueryMintAllowances))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_MintAllowance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMintAllowance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MintAllowance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_MintAllowance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MintAllowance(ctx, req.(*QueryMintAllowance))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "florin.v2.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authority",
			Handler:    _Query_Authority_Handler,
		},
		{
			MethodName: "AllowedDenoms",
			Handler:    _Query_AllowedDenoms_Handler,
		},
		{
			MethodName: "Owners",
			Handler:    _Query_Owners_Handler,
		},
		{
			MethodName: "Owner",
			Handler:    _Query_Owner_Handler,
		},
		{
			MethodName: "Systems",
			Handler:    _Query_Systems_Handler,
		},
		{
			MethodName: "SystemsByDenom",
			Handler:    _Query_SystemsByDenom_Handler,
		},
		{
			MethodName: "Admins",
			Handler:    _Query_Admins_Handler,
		},
		{
			MethodName: "AdminsByDenom",
			Handler:    _Query_AdminsByDenom_Handler,
		},
		{
			MethodName: "MaxMintAllowances",
			Handler:    _Query_MaxMintAllowances_Handler,
		},
		{
			MethodName: "MaxMintAllowance",
			Handler:    _Query_MaxMintAllowance_Handler,
		},
		{
			MethodName: "MintAllowances",
			Handler:    _Query_MintAllowances_Handler,
		},
		{
			MethodName: "MintAllowance",
			Handler:    _Query_MintAllowance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "florin/v2/query.proto",
}
