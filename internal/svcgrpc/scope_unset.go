package svcgrpc

import (
	context "context"

	"github.com/sandrolain/gomsvc/pkg/grpclib"
	"github.com/sandrolain/permissions/internal/dbsvc"
	"github.com/sandrolain/permissions/internal/models"
	g "github.com/sandrolain/permissions/pkg/grpc"
)

func (m *grpcServer) UnsetUserScope(ctx context.Context, req *g.UnsetUserScopeRequest) (res *g.UnsetUserScopeResponse, err error) {
	if !validEntity(req.User) || !validScope(req.Scope) {
		err = grpclib.InvalidArgument("Invalid User or Scope")
		return
	}

	affected, err := m.db.UnsetScope(ctx, req.User, req.Scope)
	if err != nil {
		err = grpclib.InternalError("Query Error", err)
		return
	}

	res = &g.UnsetUserScopeResponse{
		Affected: affected,
	}
	return
}

func (m *grpcServer) UnsetRoleScope(ctx context.Context, req *g.UnsetRoleScopeRequest) (res *g.UnsetRoleScopeResponse, err error) {
	if !validEntity(req.Role) || !validScope(req.Scope) {
		err = grpclib.InvalidArgument("Invalid Role or Scope")
		return
	}

	affected, err := m.db.UnsetScope(ctx, dbsvc.FormatRoleScope(req.Role), req.Scope)
	if err != nil {
		err = grpclib.InternalError("Query Error", err)
		return
	}

	res = &g.UnsetRoleScopeResponse{
		Affected: affected,
	}
	return
}

func (m *grpcServer) UnsetGlobalScope(ctx context.Context, req *g.UnsetGlobalScopeRequest) (res *g.UnsetGlobalScopeResponse, err error) {
	if !validScope(req.Scope) {
		err = grpclib.InvalidArgument("Invalid Scope")
		return
	}

	affected, err := m.db.UnsetScope(ctx, models.GlobalEntity, req.Scope)
	if err != nil {
		err = grpclib.InternalError("Query Error", err)
		return
	}

	res = &g.UnsetGlobalScopeResponse{
		Affected: affected,
	}
	return
}
