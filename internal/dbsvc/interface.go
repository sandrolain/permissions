package dbsvc

import (
	"context"

	"github.com/sandrolain/permissions/internal/models"
)

type DBServiceInterface interface {
	GetScopes(ctx context.Context, entity string, scopePattern string) (res []models.Permission, err error)
	SetScope(ctx context.Context, entity string, scope string, allow bool) (affected bool, err error)
	UnsetScope(ctx context.Context, entity string, scope string) (affected bool, err error)
	IsAllowed(ctx context.Context, entity string, scope string) (found bool, res bool, err error)
	IsAllowedNegated(ctx context.Context, scopes []string, scope string, negate bool) (found bool, res bool, err error)
	GetUserRolesPermissions(ctx context.Context, user string) (res []models.Permission, err error)
	GetUserRolesScopes(ctx context.Context, user string) (res []string, err error)
	GetUserRoles(ctx context.Context, user string) (res []string, err error)
}
