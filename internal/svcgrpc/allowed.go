package svcgrpc

import (
	context "context"

	"github.com/sandrolain/gomsvc/pkg/grpclib"
	"github.com/sandrolain/permissions/internal/dbsvc"
	g "github.com/sandrolain/permissions/pkg/grpc"
	"github.com/sandrolain/permissions/pkg/models"
)

func (m *grpcServer) RoleAllowed(ctx context.Context, request *g.RoleAllowedRequest) (res *g.RoleAllowedResponse, err error) {
	if !validEntity(request.Role) || !validScope(request.Scope) {
		err = grpclib.InvalidArgument("Invalid Role or Scope")
		return
	}

	// Check for role
	found, allowed, e := m.db.IsAllowed(ctx, dbsvc.FormatRoleScope(request.Role), request.Scope)
	if e != nil {
		err = grpclib.InternalError("Query Error", e)
		return
	}
	if found {
		res = &g.RoleAllowedResponse{
			Allowed: allowed,
		}
		return
	}

	// Check for global
	_, allowed, e = m.db.IsAllowed(ctx, models.GlobalEntity, request.Scope)
	if e != nil {
		err = grpclib.InternalError("Query Error", e)
		return
	}

	res = &g.RoleAllowedResponse{
		Allowed: allowed,
	}
	return
}

func (m *grpcServer) UserAllowed(ctx context.Context, request *g.UserAllowedRequest) (res *g.UserAllowedResponse, err error) {
	if !validEntity(request.User) || !validScope(request.Scope) {
		err = grpclib.InvalidArgument("Invalid User or Scope")
		return
	}

	// Check for user
	found, allowed, e := m.db.IsAllowed(ctx, request.User, request.Scope)
	if e != nil {
		err = grpclib.InternalError("Query Error", e)
		return
	}
	if found {
		res = &g.UserAllowedResponse{
			Allowed: allowed,
		}
		return
	}

	// Check global
	found, allowed, e = m.db.IsAllowed(ctx, models.GlobalEntity, request.Scope)
	if e != nil {
		err = grpclib.InternalError("Query Error", e)
		return
	}

	scopes, e := m.db.GetUserRolesScopes(ctx, request.User)
	if e != nil {
		err = grpclib.InternalError("Query Error", e)
		return
	}

	// Check for User Roles with negated Scopes
	if len(scopes) > 0 {
		found, neg, e := m.db.IsAllowedNegated(ctx, scopes, request.Scope, allowed)
		if e != nil {
			err = grpclib.InternalError("Query Error", e)
			return
		}
		if found {
			allowed = neg
		}
	}

	res = &g.UserAllowedResponse{
		Allowed: allowed,
	}
	return
}

func (m *grpcServer) GlobalAllowed(ctx context.Context, request *g.GlobalAllowedRequest) (res *g.GlobalAllowedResponse, err error) {
	if !validScope(request.Scope) {
		err = grpclib.InvalidArgument("Invalid Scope")
		return
	}

	// Check global
	_, allowed, e := m.db.IsAllowed(ctx, models.GlobalEntity, request.Scope)
	if e != nil {
		err = grpclib.InternalError("Query Error", e)
		return
	}

	res = &g.GlobalAllowedResponse{
		Allowed: allowed,
	}
	return
}
