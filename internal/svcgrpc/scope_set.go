package svcgrpc

import (
	context "context"

	"github.com/sandrolain/gomsvc/pkg/grpclib"
	"github.com/sandrolain/permissions/internal/dbsvc"
	g "github.com/sandrolain/permissions/pkg/grpc"
	"github.com/sandrolain/permissions/pkg/models"
)

// gRPC APIs

func (m *grpcServer) SetUserScope(ctx context.Context, req *g.SetUserScopeRequest) (res *g.SetUserScopeResponse, err error) {
	if !validEntity(req.User) || !validScope(req.Scope) {
		err = grpclib.InvalidArgument("Invalid User or Scope")
		return
	}

	affected, e := m.db.SetScope(ctx, req.User, req.Scope, req.Allowed)
	if e != nil {
		err = grpclib.InternalError("Query Error", e)
		return
	}

	res = &g.SetUserScopeResponse{
		Affected: affected,
	}
	return
}

func (m *grpcServer) SetRoleScope(ctx context.Context, req *g.SetRoleScopeRequest) (res *g.SetRoleScopeResponse, err error) {
	if !validEntity(req.Role) || !validScope(req.Scope) {
		err = grpclib.InvalidArgument("Invalid Role or Scope")
		return
	}

	affected, e := m.db.SetScope(ctx, dbsvc.FormatRoleScope(req.Role), req.Scope, req.Allowed)
	if e != nil {
		err = grpclib.InternalError("Query Error", e)
		return
	}

	res = &g.SetRoleScopeResponse{
		Affected: affected,
	}
	return
}

func (m *grpcServer) SetGlobalScope(ctx context.Context, req *g.SetGlobalScopeRequest) (res *g.SetGlobalScopeResponse, err error) {
	if !validScope(req.Scope) {
		err = grpclib.InvalidArgument("Invalid Scope")
		return
	}

	affected, e := m.db.SetScope(ctx, models.GlobalEntity, req.Scope, req.Allowed)
	if e != nil {
		err = grpclib.InternalError("Query Error", e)
		return
	}

	res = &g.SetGlobalScopeResponse{
		Affected: affected,
	}
	return
}
