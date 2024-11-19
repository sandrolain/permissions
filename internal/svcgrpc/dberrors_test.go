package svcgrpc

import (
	"context"
	"errors"
	"testing"

	"github.com/sandrolain/permissions/internal/dbsvc"
	g "github.com/sandrolain/permissions/pkg/grpc"
	"github.com/sandrolain/permissions/pkg/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorDBService is a mock DB service that always returns errors
type ErrorDBService struct {
	dbsvc.DBServiceInterface
}

func (m *ErrorDBService) GetScopes(ctx context.Context, entity string, scopePattern string) (res []models.Permission, err error) {
	return nil, errors.New("database error")
}

func (m *ErrorDBService) SetScope(ctx context.Context, entity string, scope string, allow bool) (affected bool, err error) {
	return false, errors.New("database error")
}

func (m *ErrorDBService) UnsetScope(ctx context.Context, entity string, scope string) (affected bool, err error) {
	return false, errors.New("database error")
}

func (m *ErrorDBService) IsAllowed(ctx context.Context, entity string, scope string) (found bool, res bool, err error) {
	return false, false, errors.New("database error")
}

func (m *ErrorDBService) IsAllowedNegated(ctx context.Context, scopes []string, scope string, negate bool) (found bool, res bool, err error) {
	return false, false, errors.New("database error")
}

func (m *ErrorDBService) GetUserRolesPermissions(ctx context.Context, user string) (res []models.Permission, err error) {
	return nil, errors.New("database error")
}

func (m *ErrorDBService) GetUserRolesScopes(ctx context.Context, user string) (res []string, err error) {
	return nil, errors.New("database error")
}

func (m *ErrorDBService) GetUserRoles(ctx context.Context, user string) (res []string, err error) {
	return nil, errors.New("database error")
}

func TestDatabaseErrors(t *testing.T) {
	db := &ErrorDBService{}
	srv := NewGrpcServer(db)

	t.Run("GetUserRoles database error", func(t *testing.T) {
		req := &g.GetUserRolesRequest{
			User: "user123",
		}
		_, err := srv.GetUserRoles(context.Background(), req)
		if err == nil {
			t.Error("expected error, got nil")
			return
		}
		if status.Code(err) != codes.Internal {
			t.Errorf("expected Internal error, got %v", err)
		}
	})

	t.Run("SetUserRole database error", func(t *testing.T) {
		req := &g.SetUserRoleRequest{
			User: "user123",
			Role: "admin",
		}
		_, err := srv.SetUserRole(context.Background(), req)
		if err == nil {
			t.Error("expected error, got nil")
			return
		}
		if status.Code(err) != codes.Internal {
			t.Errorf("expected Internal error, got %v", err)
		}
	})

	t.Run("UnsetUserRole database error", func(t *testing.T) {
		req := &g.UnsetUserRoleRequest{
			User: "user123",
			Role: "admin",
		}
		_, err := srv.UnsetUserRole(context.Background(), req)
		if err == nil {
			t.Error("expected error, got nil")
			return
		}
		if status.Code(err) != codes.Internal {
			t.Errorf("expected Internal error, got %v", err)
		}
	})

	t.Run("GetUserScopes database error", func(t *testing.T) {
		req := &g.GetUserScopesRequest{
			User: "user123",
		}
		_, err := srv.GetUserScopes(context.Background(), req)
		if err == nil {
			t.Error("expected error, got nil")
			return
		}
		if status.Code(err) != codes.Internal {
			t.Errorf("expected Internal error, got %v", err)
		}
	})

	t.Run("SetUserScope database error", func(t *testing.T) {
		req := &g.SetUserScopeRequest{
			User:    "user123",
			Scope:   "read",
			Allowed: true,
		}
		_, err := srv.SetUserScope(context.Background(), req)
		if err == nil {
			t.Error("expected error, got nil")
			return
		}
		if status.Code(err) != codes.Internal {
			t.Errorf("expected Internal error, got %v", err)
		}
	})

	t.Run("UnsetUserScope database error", func(t *testing.T) {
		req := &g.UnsetUserScopeRequest{
			User:  "user123",
			Scope: "read",
		}
		_, err := srv.UnsetUserScope(context.Background(), req)
		if err == nil {
			t.Error("expected error, got nil")
			return
		}
		if status.Code(err) != codes.Internal {
			t.Errorf("expected Internal error, got %v", err)
		}
	})

	t.Run("UserAllowed database error", func(t *testing.T) {
		req := &g.UserAllowedRequest{
			User:  "user123",
			Scope: "read",
		}
		_, err := srv.UserAllowed(context.Background(), req)
		if err == nil {
			t.Error("expected error, got nil")
			return
		}
		if status.Code(err) != codes.Internal {
			t.Errorf("expected Internal error, got %v", err)
		}
	})
}
