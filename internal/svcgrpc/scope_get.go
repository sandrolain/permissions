package svcgrpc

import (
	context "context"

	"github.com/sandrolain/gomsvc/pkg/grpclib"
	"github.com/sandrolain/permissions/internal/dbsvc"
	"github.com/sandrolain/permissions/internal/models"
	g "github.com/sandrolain/permissions/pkg/grpc"
)

func (m *grpcServer) GetUserScopes(ctx context.Context, req *g.GetUserScopesRequest) (res *g.GetUserScopesResponse, err error) {
	if !validEntity(req.User) {
		err = grpclib.InvalidArgument("Invalid User")
		return
	}

	var pattern string
	if len(req.ScopePattern) > 0 && req.ScopePattern != "*" {
		if !validPattern(req.ScopePattern) {
			err = grpclib.InvalidArgument("Invalid Scope Pattern")
			return
		}
		pattern = req.ScopePattern
	}

	scopes, e := m.db.GetScopes(ctx, req.User, pattern)
	if e != nil {
		err = grpclib.InternalError("Query Error", e)
		return
	}

	res = &g.GetUserScopesResponse{
		Scopes: dbsvc.GetScopeItems(scopes),
	}

	return
}

func (m *grpcServer) GetRoleScopes(ctx context.Context, req *g.GetRoleScopesRequest) (res *g.GetRoleScopesResponse, err error) {
	if !validEntity(req.Role) {
		err = grpclib.InvalidArgument("Invalid Role")
		return
	}

	var pattern string
	if len(req.ScopePattern) > 0 && req.ScopePattern != "*" {
		if !validPattern(req.ScopePattern) {
			err = grpclib.InvalidArgument("Invalid Scope Pattern")
			return
		}
		pattern = req.ScopePattern
	}

	scopes, e := m.db.GetScopes(ctx, dbsvc.FormatRoleScope(req.Role), pattern)
	if e != nil {
		err = grpclib.InternalError("Query Error", e)
		return
	}

	res = &g.GetRoleScopesResponse{
		Scopes: dbsvc.GetScopeItems(scopes),
	}

	return
}

func (m *grpcServer) GetGlobalScopes(ctx context.Context, req *g.GetGlobalScopesRequest) (res *g.GetGlobalScopesResponse, err error) {
	var pattern string
	if len(req.ScopePattern) > 0 && req.ScopePattern != "*" {
		if !validPattern(req.ScopePattern) {
			err = grpclib.InvalidArgument("Invalid Scope Pattern")
			return
		}
		pattern = req.ScopePattern
	}

	scopes, e := m.db.GetScopes(ctx, models.GlobalEntity, pattern)
	if e != nil {
		err = grpclib.InternalError("Query Error", e)
		return
	}

	res = &g.GetGlobalScopesResponse{
		Scopes: dbsvc.GetScopeItems(scopes),
	}

	return
}
