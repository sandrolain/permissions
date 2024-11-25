// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: pkg/grpc/permissions.proto

package permissionsGrpc

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
	PermissionsService_GetUserRoles_FullMethodName     = "/permissions.PermissionsService/GetUserRoles"
	PermissionsService_SetUserRole_FullMethodName      = "/permissions.PermissionsService/SetUserRole"
	PermissionsService_UnsetUserRole_FullMethodName    = "/permissions.PermissionsService/UnsetUserRole"
	PermissionsService_GetGlobalScopes_FullMethodName  = "/permissions.PermissionsService/GetGlobalScopes"
	PermissionsService_GetRoleScopes_FullMethodName    = "/permissions.PermissionsService/GetRoleScopes"
	PermissionsService_GetUserScopes_FullMethodName    = "/permissions.PermissionsService/GetUserScopes"
	PermissionsService_SetGlobalScope_FullMethodName   = "/permissions.PermissionsService/SetGlobalScope"
	PermissionsService_SetRoleScope_FullMethodName     = "/permissions.PermissionsService/SetRoleScope"
	PermissionsService_SetUserScope_FullMethodName     = "/permissions.PermissionsService/SetUserScope"
	PermissionsService_UnsetGlobalScope_FullMethodName = "/permissions.PermissionsService/UnsetGlobalScope"
	PermissionsService_UnsetRoleScope_FullMethodName   = "/permissions.PermissionsService/UnsetRoleScope"
	PermissionsService_UnsetUserScope_FullMethodName   = "/permissions.PermissionsService/UnsetUserScope"
	PermissionsService_GlobalAllowed_FullMethodName    = "/permissions.PermissionsService/GlobalAllowed"
	PermissionsService_RoleAllowed_FullMethodName      = "/permissions.PermissionsService/RoleAllowed"
	PermissionsService_UserAllowed_FullMethodName      = "/permissions.PermissionsService/UserAllowed"
)

// PermissionsServiceClient is the client API for PermissionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionsServiceClient interface {
	GetUserRoles(ctx context.Context, in *GetUserRolesRequest, opts ...grpc.CallOption) (*GetUserRolesResponse, error)
	SetUserRole(ctx context.Context, in *SetUserRoleRequest, opts ...grpc.CallOption) (*SetUserRoleResponse, error)
	UnsetUserRole(ctx context.Context, in *UnsetUserRoleRequest, opts ...grpc.CallOption) (*UnsetUserRoleResponse, error)
	GetGlobalScopes(ctx context.Context, in *GetGlobalScopesRequest, opts ...grpc.CallOption) (*GetGlobalScopesResponse, error)
	GetRoleScopes(ctx context.Context, in *GetRoleScopesRequest, opts ...grpc.CallOption) (*GetRoleScopesResponse, error)
	GetUserScopes(ctx context.Context, in *GetUserScopesRequest, opts ...grpc.CallOption) (*GetUserScopesResponse, error)
	SetGlobalScope(ctx context.Context, in *SetGlobalScopeRequest, opts ...grpc.CallOption) (*SetGlobalScopeResponse, error)
	SetRoleScope(ctx context.Context, in *SetRoleScopeRequest, opts ...grpc.CallOption) (*SetRoleScopeResponse, error)
	SetUserScope(ctx context.Context, in *SetUserScopeRequest, opts ...grpc.CallOption) (*SetUserScopeResponse, error)
	UnsetGlobalScope(ctx context.Context, in *UnsetGlobalScopeRequest, opts ...grpc.CallOption) (*UnsetGlobalScopeResponse, error)
	UnsetRoleScope(ctx context.Context, in *UnsetRoleScopeRequest, opts ...grpc.CallOption) (*UnsetRoleScopeResponse, error)
	UnsetUserScope(ctx context.Context, in *UnsetUserScopeRequest, opts ...grpc.CallOption) (*UnsetUserScopeResponse, error)
	GlobalAllowed(ctx context.Context, in *GlobalAllowedRequest, opts ...grpc.CallOption) (*GlobalAllowedResponse, error)
	RoleAllowed(ctx context.Context, in *RoleAllowedRequest, opts ...grpc.CallOption) (*RoleAllowedResponse, error)
	UserAllowed(ctx context.Context, in *UserAllowedRequest, opts ...grpc.CallOption) (*UserAllowedResponse, error)
}

type permissionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionsServiceClient(cc grpc.ClientConnInterface) PermissionsServiceClient {
	return &permissionsServiceClient{cc}
}

func (c *permissionsServiceClient) GetUserRoles(ctx context.Context, in *GetUserRolesRequest, opts ...grpc.CallOption) (*GetUserRolesResponse, error) {
	out := new(GetUserRolesResponse)
	err := c.cc.Invoke(ctx, PermissionsService_GetUserRoles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) SetUserRole(ctx context.Context, in *SetUserRoleRequest, opts ...grpc.CallOption) (*SetUserRoleResponse, error) {
	out := new(SetUserRoleResponse)
	err := c.cc.Invoke(ctx, PermissionsService_SetUserRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) UnsetUserRole(ctx context.Context, in *UnsetUserRoleRequest, opts ...grpc.CallOption) (*UnsetUserRoleResponse, error) {
	out := new(UnsetUserRoleResponse)
	err := c.cc.Invoke(ctx, PermissionsService_UnsetUserRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) GetGlobalScopes(ctx context.Context, in *GetGlobalScopesRequest, opts ...grpc.CallOption) (*GetGlobalScopesResponse, error) {
	out := new(GetGlobalScopesResponse)
	err := c.cc.Invoke(ctx, PermissionsService_GetGlobalScopes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) GetRoleScopes(ctx context.Context, in *GetRoleScopesRequest, opts ...grpc.CallOption) (*GetRoleScopesResponse, error) {
	out := new(GetRoleScopesResponse)
	err := c.cc.Invoke(ctx, PermissionsService_GetRoleScopes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) GetUserScopes(ctx context.Context, in *GetUserScopesRequest, opts ...grpc.CallOption) (*GetUserScopesResponse, error) {
	out := new(GetUserScopesResponse)
	err := c.cc.Invoke(ctx, PermissionsService_GetUserScopes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) SetGlobalScope(ctx context.Context, in *SetGlobalScopeRequest, opts ...grpc.CallOption) (*SetGlobalScopeResponse, error) {
	out := new(SetGlobalScopeResponse)
	err := c.cc.Invoke(ctx, PermissionsService_SetGlobalScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) SetRoleScope(ctx context.Context, in *SetRoleScopeRequest, opts ...grpc.CallOption) (*SetRoleScopeResponse, error) {
	out := new(SetRoleScopeResponse)
	err := c.cc.Invoke(ctx, PermissionsService_SetRoleScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) SetUserScope(ctx context.Context, in *SetUserScopeRequest, opts ...grpc.CallOption) (*SetUserScopeResponse, error) {
	out := new(SetUserScopeResponse)
	err := c.cc.Invoke(ctx, PermissionsService_SetUserScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) UnsetGlobalScope(ctx context.Context, in *UnsetGlobalScopeRequest, opts ...grpc.CallOption) (*UnsetGlobalScopeResponse, error) {
	out := new(UnsetGlobalScopeResponse)
	err := c.cc.Invoke(ctx, PermissionsService_UnsetGlobalScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) UnsetRoleScope(ctx context.Context, in *UnsetRoleScopeRequest, opts ...grpc.CallOption) (*UnsetRoleScopeResponse, error) {
	out := new(UnsetRoleScopeResponse)
	err := c.cc.Invoke(ctx, PermissionsService_UnsetRoleScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) UnsetUserScope(ctx context.Context, in *UnsetUserScopeRequest, opts ...grpc.CallOption) (*UnsetUserScopeResponse, error) {
	out := new(UnsetUserScopeResponse)
	err := c.cc.Invoke(ctx, PermissionsService_UnsetUserScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) GlobalAllowed(ctx context.Context, in *GlobalAllowedRequest, opts ...grpc.CallOption) (*GlobalAllowedResponse, error) {
	out := new(GlobalAllowedResponse)
	err := c.cc.Invoke(ctx, PermissionsService_GlobalAllowed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) RoleAllowed(ctx context.Context, in *RoleAllowedRequest, opts ...grpc.CallOption) (*RoleAllowedResponse, error) {
	out := new(RoleAllowedResponse)
	err := c.cc.Invoke(ctx, PermissionsService_RoleAllowed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsServiceClient) UserAllowed(ctx context.Context, in *UserAllowedRequest, opts ...grpc.CallOption) (*UserAllowedResponse, error) {
	out := new(UserAllowedResponse)
	err := c.cc.Invoke(ctx, PermissionsService_UserAllowed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionsServiceServer is the server API for PermissionsService service.
// All implementations must embed UnimplementedPermissionsServiceServer
// for forward compatibility
type PermissionsServiceServer interface {
	GetUserRoles(context.Context, *GetUserRolesRequest) (*GetUserRolesResponse, error)
	SetUserRole(context.Context, *SetUserRoleRequest) (*SetUserRoleResponse, error)
	UnsetUserRole(context.Context, *UnsetUserRoleRequest) (*UnsetUserRoleResponse, error)
	GetGlobalScopes(context.Context, *GetGlobalScopesRequest) (*GetGlobalScopesResponse, error)
	GetRoleScopes(context.Context, *GetRoleScopesRequest) (*GetRoleScopesResponse, error)
	GetUserScopes(context.Context, *GetUserScopesRequest) (*GetUserScopesResponse, error)
	SetGlobalScope(context.Context, *SetGlobalScopeRequest) (*SetGlobalScopeResponse, error)
	SetRoleScope(context.Context, *SetRoleScopeRequest) (*SetRoleScopeResponse, error)
	SetUserScope(context.Context, *SetUserScopeRequest) (*SetUserScopeResponse, error)
	UnsetGlobalScope(context.Context, *UnsetGlobalScopeRequest) (*UnsetGlobalScopeResponse, error)
	UnsetRoleScope(context.Context, *UnsetRoleScopeRequest) (*UnsetRoleScopeResponse, error)
	UnsetUserScope(context.Context, *UnsetUserScopeRequest) (*UnsetUserScopeResponse, error)
	GlobalAllowed(context.Context, *GlobalAllowedRequest) (*GlobalAllowedResponse, error)
	RoleAllowed(context.Context, *RoleAllowedRequest) (*RoleAllowedResponse, error)
	UserAllowed(context.Context, *UserAllowedRequest) (*UserAllowedResponse, error)
	mustEmbedUnimplementedPermissionsServiceServer()
}

// UnimplementedPermissionsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPermissionsServiceServer struct {
}

func (UnimplementedPermissionsServiceServer) GetUserRoles(context.Context, *GetUserRolesRequest) (*GetUserRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRoles not implemented")
}
func (UnimplementedPermissionsServiceServer) SetUserRole(context.Context, *SetUserRoleRequest) (*SetUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserRole not implemented")
}
func (UnimplementedPermissionsServiceServer) UnsetUserRole(context.Context, *UnsetUserRoleRequest) (*UnsetUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsetUserRole not implemented")
}
func (UnimplementedPermissionsServiceServer) GetGlobalScopes(context.Context, *GetGlobalScopesRequest) (*GetGlobalScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGlobalScopes not implemented")
}
func (UnimplementedPermissionsServiceServer) GetRoleScopes(context.Context, *GetRoleScopesRequest) (*GetRoleScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleScopes not implemented")
}
func (UnimplementedPermissionsServiceServer) GetUserScopes(context.Context, *GetUserScopesRequest) (*GetUserScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserScopes not implemented")
}
func (UnimplementedPermissionsServiceServer) SetGlobalScope(context.Context, *SetGlobalScopeRequest) (*SetGlobalScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGlobalScope not implemented")
}
func (UnimplementedPermissionsServiceServer) SetRoleScope(context.Context, *SetRoleScopeRequest) (*SetRoleScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRoleScope not implemented")
}
func (UnimplementedPermissionsServiceServer) SetUserScope(context.Context, *SetUserScopeRequest) (*SetUserScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserScope not implemented")
}
func (UnimplementedPermissionsServiceServer) UnsetGlobalScope(context.Context, *UnsetGlobalScopeRequest) (*UnsetGlobalScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsetGlobalScope not implemented")
}
func (UnimplementedPermissionsServiceServer) UnsetRoleScope(context.Context, *UnsetRoleScopeRequest) (*UnsetRoleScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsetRoleScope not implemented")
}
func (UnimplementedPermissionsServiceServer) UnsetUserScope(context.Context, *UnsetUserScopeRequest) (*UnsetUserScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsetUserScope not implemented")
}
func (UnimplementedPermissionsServiceServer) GlobalAllowed(context.Context, *GlobalAllowedRequest) (*GlobalAllowedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GlobalAllowed not implemented")
}
func (UnimplementedPermissionsServiceServer) RoleAllowed(context.Context, *RoleAllowedRequest) (*RoleAllowedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleAllowed not implemented")
}
func (UnimplementedPermissionsServiceServer) UserAllowed(context.Context, *UserAllowedRequest) (*UserAllowedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAllowed not implemented")
}
func (UnimplementedPermissionsServiceServer) mustEmbedUnimplementedPermissionsServiceServer() {}

// UnsafePermissionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionsServiceServer will
// result in compilation errors.
type UnsafePermissionsServiceServer interface {
	mustEmbedUnimplementedPermissionsServiceServer()
}

func RegisterPermissionsServiceServer(s grpc.ServiceRegistrar, srv PermissionsServiceServer) {
	s.RegisterService(&PermissionsService_ServiceDesc, srv)
}

func _PermissionsService_GetUserRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).GetUserRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_GetUserRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).GetUserRoles(ctx, req.(*GetUserRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_SetUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).SetUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_SetUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).SetUserRole(ctx, req.(*SetUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_UnsetUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsetUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).UnsetUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_UnsetUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).UnsetUserRole(ctx, req.(*UnsetUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_GetGlobalScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGlobalScopesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).GetGlobalScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_GetGlobalScopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).GetGlobalScopes(ctx, req.(*GetGlobalScopesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_GetRoleScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleScopesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).GetRoleScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_GetRoleScopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).GetRoleScopes(ctx, req.(*GetRoleScopesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_GetUserScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserScopesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).GetUserScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_GetUserScopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).GetUserScopes(ctx, req.(*GetUserScopesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_SetGlobalScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGlobalScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).SetGlobalScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_SetGlobalScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).SetGlobalScope(ctx, req.(*SetGlobalScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_SetRoleScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRoleScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).SetRoleScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_SetRoleScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).SetRoleScope(ctx, req.(*SetRoleScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_SetUserScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).SetUserScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_SetUserScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).SetUserScope(ctx, req.(*SetUserScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_UnsetGlobalScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsetGlobalScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).UnsetGlobalScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_UnsetGlobalScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).UnsetGlobalScope(ctx, req.(*UnsetGlobalScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_UnsetRoleScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsetRoleScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).UnsetRoleScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_UnsetRoleScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).UnsetRoleScope(ctx, req.(*UnsetRoleScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_UnsetUserScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsetUserScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).UnsetUserScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_UnsetUserScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).UnsetUserScope(ctx, req.(*UnsetUserScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_GlobalAllowed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GlobalAllowedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).GlobalAllowed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_GlobalAllowed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).GlobalAllowed(ctx, req.(*GlobalAllowedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_RoleAllowed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleAllowedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).RoleAllowed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_RoleAllowed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).RoleAllowed(ctx, req.(*RoleAllowedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionsService_UserAllowed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAllowedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).UserAllowed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_UserAllowed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).UserAllowed(ctx, req.(*UserAllowedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionsService_ServiceDesc is the grpc.ServiceDesc for PermissionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "permissions.PermissionsService",
	HandlerType: (*PermissionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserRoles",
			Handler:    _PermissionsService_GetUserRoles_Handler,
		},
		{
			MethodName: "SetUserRole",
			Handler:    _PermissionsService_SetUserRole_Handler,
		},
		{
			MethodName: "UnsetUserRole",
			Handler:    _PermissionsService_UnsetUserRole_Handler,
		},
		{
			MethodName: "GetGlobalScopes",
			Handler:    _PermissionsService_GetGlobalScopes_Handler,
		},
		{
			MethodName: "GetRoleScopes",
			Handler:    _PermissionsService_GetRoleScopes_Handler,
		},
		{
			MethodName: "GetUserScopes",
			Handler:    _PermissionsService_GetUserScopes_Handler,
		},
		{
			MethodName: "SetGlobalScope",
			Handler:    _PermissionsService_SetGlobalScope_Handler,
		},
		{
			MethodName: "SetRoleScope",
			Handler:    _PermissionsService_SetRoleScope_Handler,
		},
		{
			MethodName: "SetUserScope",
			Handler:    _PermissionsService_SetUserScope_Handler,
		},
		{
			MethodName: "UnsetGlobalScope",
			Handler:    _PermissionsService_UnsetGlobalScope_Handler,
		},
		{
			MethodName: "UnsetRoleScope",
			Handler:    _PermissionsService_UnsetRoleScope_Handler,
		},
		{
			MethodName: "UnsetUserScope",
			Handler:    _PermissionsService_UnsetUserScope_Handler,
		},
		{
			MethodName: "GlobalAllowed",
			Handler:    _PermissionsService_GlobalAllowed_Handler,
		},
		{
			MethodName: "RoleAllowed",
			Handler:    _PermissionsService_RoleAllowed_Handler,
		},
		{
			MethodName: "UserAllowed",
			Handler:    _PermissionsService_UserAllowed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/permissions.proto",
}
