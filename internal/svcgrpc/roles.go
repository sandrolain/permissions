package svcgrpc

import (
	context "context"
	"fmt"

	"github.com/sandrolain/gomsvc/pkg/grpclib"
	"github.com/sandrolain/permissions/internal/dbsvc"
	g "github.com/sandrolain/permissions/pkg/grpc"
	"github.com/sandrolain/permissions/pkg/models"
)

func (m *grpcServer) GetUserRoles(ctx context.Context, request *g.GetUserRolesRequest) (res *g.GetUserRolesResponse, err error) {
	if !validEntity(request.User) {
		err = grpclib.InvalidArgument("Invalid User")
		return
	}

	roles, e := m.db.GetUserRoles(ctx, request.User)
	if e != nil {
		err = grpclib.InternalError("Query Error", e)
		return
	}

	res = &g.GetUserRolesResponse{
		Roles: roles,
	}
	return
}

func (m *grpcServer) SetUserRole(ctx context.Context, req *g.SetUserRoleRequest) (res *g.SetUserRoleResponse, err error) {
	if !validEntity(req.User) || !validEntity(req.Role) {
		err = grpclib.InvalidArgument("Invalid User or Role")
		return
	}

	scope := fmt.Sprintf("%s%s", models.RolePrefix, req.Role)

	affected, e := m.db.SetScope(ctx, req.User, scope, true)
	if e != nil {
		err = grpclib.InternalError("Query Error", e)
		return
	}

	roles, e := m.db.GetUserRoles(ctx, req.User)
	if e != nil {
		err = grpclib.InternalError("Query Error", e)
		return
	}

	res = &g.SetUserRoleResponse{
		Affected: affected,
		Roles:    roles,
	}
	return
}

func (m *grpcServer) UnsetUserRole(ctx context.Context, req *g.UnsetUserRoleRequest) (res *g.UnsetUserRoleResponse, err error) {
	if !validEntity(req.User) || !validEntity(req.Role) {
		err = grpclib.InvalidArgument("Invalid User or Role")
		return
	}

	affected, e := m.db.UnsetScope(ctx, req.User, dbsvc.FormatRoleScope(req.Role))
	if e != nil {
		err = grpclib.InternalError("Query Error", e)
		return
	}

	roles, e := m.db.GetUserRoles(ctx, req.User)
	if e != nil {
		err = grpclib.InternalError("Query Error", e)
		return
	}

	res = &g.UnsetUserRoleResponse{
		Affected: affected,
		Roles:    roles,
	}
	return
}
