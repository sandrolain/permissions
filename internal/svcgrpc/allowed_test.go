package svcgrpc

import (
	"context"
	"testing"

	"github.com/sandrolain/permissions/internal/models"
	g "github.com/sandrolain/permissions/pkg/grpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock DBService
type mockDBService struct {
	mock.Mock
}

func (m *mockDBService) GetScopes(ctx context.Context, entity string, scopePattern string) ([]models.Permission, error) {
	args := m.Called(ctx, entity, scopePattern)
	return args.Get(0).([]models.Permission), args.Error(1)
}

func (m *mockDBService) SetScope(ctx context.Context, entity string, scope string, allow bool) (bool, error) {
	args := m.Called(ctx, entity, scope, allow)
	return args.Bool(0), args.Error(1)
}

func (m *mockDBService) UnsetScope(ctx context.Context, entity string, scope string) (bool, error) {
	args := m.Called(ctx, entity, scope)
	return args.Bool(0), args.Error(1)
}

func (m *mockDBService) IsAllowed(ctx context.Context, entity string, scope string) (bool, bool, error) {
	args := m.Called(ctx, entity, scope)
	return args.Bool(0), args.Bool(1), args.Error(2)
}

func (m *mockDBService) IsAllowedNegated(ctx context.Context, scopes []string, scope string, negate bool) (bool, bool, error) {
	args := m.Called(ctx, scopes, scope, negate)
	return args.Bool(0), args.Bool(1), args.Error(2)
}

func (m *mockDBService) GetUserRolesPermissions(ctx context.Context, user string) ([]models.Permission, error) {
	args := m.Called(ctx, user)
	return args.Get(0).([]models.Permission), args.Error(1)
}

func (m *mockDBService) GetUserRolesScopes(ctx context.Context, user string) ([]string, error) {
	args := m.Called(ctx, user)
	return args.Get(0).([]string), args.Error(1)
}

func (m *mockDBService) GetUserRoles(ctx context.Context, user string) ([]string, error) {
	args := m.Called(ctx, user)
	return args.Get(0).([]string), args.Error(1)
}

func TestRoleAllowed(t *testing.T) {
	ctx := context.Background()
	db := &mockDBService{}
	server := NewGrpcServer(db)

	tests := []struct {
		name        string
		role        string
		scope       string
		mockSetup   func()
		wantAllowed bool
		wantError   bool
		errorMsg    string
	}{
		{
			name:  "invalid role",
			role:  "",
			scope: "test_scope",
			mockSetup: func() {
				// No mock setup needed
			},
			wantError: true,
			errorMsg:  "Invalid Role or Scope",
		},
		{
			name:  "invalid scope",
			role:  "admin",
			scope: "",
			mockSetup: func() {
				// No mock setup needed
			},
			wantError: true,
			errorMsg:  "Invalid Role or Scope",
		},
		{
			name:  "role allowed",
			role:  "admin",
			scope: "test:scope",
			mockSetup: func() {
				db.On("IsAllowed", ctx, "R$admin", "test:scope").
					Return(true, true, nil).Once()
			},
			wantAllowed: true,
		},
		{
			name:  "role denied",
			role:  "admin",
			scope: "test_scope",
			mockSetup: func() {
				db.On("IsAllowed", ctx, "R$admin", "test_scope").
					Return(true, false, nil).Once()
			},
			wantAllowed: false,
		},
		{
			name:  "role not found, check global allowed",
			role:  "admin",
			scope: "test_scope",
			mockSetup: func() {
				db.On("IsAllowed", ctx, "R$admin", "test_scope").
					Return(false, false, nil).Once()
				db.On("IsAllowed", ctx, models.GlobalEntity, "test_scope").
					Return(true, true, nil).Once()
			},
			wantAllowed: true,
		},
		{
			name:  "role not found, check global denied",
			role:  "admin",
			scope: "test_scope",
			mockSetup: func() {
				db.On("IsAllowed", ctx, "R$admin", "test_scope").
					Return(false, false, nil).Once()
				db.On("IsAllowed", ctx, models.GlobalEntity, "test_scope").
					Return(true, false, nil).Once()
			},
			wantAllowed: false,
		},
		{
			name:  "database error",
			role:  "admin",
			scope: "test_scope",
			mockSetup: func() {
				db.On("IsAllowed", ctx, "R$admin", "test_scope").
					Return(false, false, assert.AnError).Once()
			},
			wantError: true,
			errorMsg:  "Query Error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			res, err := server.RoleAllowed(ctx, &g.RoleAllowedRequest{
				Role:  tt.role,
				Scope: tt.scope,
			})

			if tt.wantError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
				assert.Nil(t, res)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.Equal(t, tt.wantAllowed, res.Allowed)
		})
	}
}

func TestUserAllowed(t *testing.T) {
	ctx := context.Background()
	db := &mockDBService{}
	server := NewGrpcServer(db)

	tests := []struct {
		name        string
		user        string
		scope       string
		mockSetup   func()
		wantAllowed bool
		wantError   bool
		errorMsg    string
	}{
		{
			name:  "invalid user",
			user:  "",
			scope: "test_scope",
			mockSetup: func() {
				// No mock setup needed
			},
			wantError: true,
			errorMsg:  "Invalid User or Scope",
		},
		{
			name:  "invalid scope",
			user:  "user1",
			scope: "",
			mockSetup: func() {
				// No mock setup needed
			},
			wantError: true,
			errorMsg:  "Invalid User or Scope",
		},
		{
			name:  "user allowed",
			user:  "user1",
			scope: "test_scope",
			mockSetup: func() {
				db.On("IsAllowed", ctx, "user1", "test_scope").
					Return(true, true, nil).Once()
			},
			wantAllowed: true,
		},
		{
			name:  "user denied",
			user:  "user1",
			scope: "test_scope",
			mockSetup: func() {
				db.On("IsAllowed", ctx, "user1", "test_scope").
					Return(true, false, nil).Once()
			},
			wantAllowed: false,
		},
		{
			name:  "user not found, check global and roles",
			user:  "user1",
			scope: "test_scope",
			mockSetup: func() {
				// User permission not found
				db.On("IsAllowed", ctx, "user1", "test_scope").
					Return(false, false, nil).Once()
				// Global permission allows
				db.On("IsAllowed", ctx, models.GlobalEntity, "test_scope").
					Return(true, true, nil).Once()
				// Get user roles
				db.On("GetUserRolesScopes", ctx, "user1").
					Return([]string{"R$admin"}, nil).Once()
				// Check role permission (negated)
				db.On("IsAllowedNegated", ctx, []string{"R$admin"}, "test_scope", true).
					Return(true, false, nil).Once()
			},
			wantAllowed: false,
		},
		{
			name:  "database error",
			user:  "user1",
			scope: "test_scope",
			mockSetup: func() {
				db.On("IsAllowed", ctx, "user1", "test_scope").
					Return(false, false, assert.AnError).Once()
			},
			wantError: true,
			errorMsg:  "Query Error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			res, err := server.UserAllowed(ctx, &g.UserAllowedRequest{
				User:  tt.user,
				Scope: tt.scope,
			})

			if tt.wantError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
				assert.Nil(t, res)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.Equal(t, tt.wantAllowed, res.Allowed)
		})
	}
}

func TestGlobalAllowed(t *testing.T) {
	ctx := context.Background()
	db := &mockDBService{}
	server := NewGrpcServer(db)

	tests := []struct {
		name        string
		scope       string
		mockSetup   func()
		wantAllowed bool
		wantError   bool
		errorMsg    string
	}{
		{
			name:  "invalid scope",
			scope: "",
			mockSetup: func() {
				// No mock setup needed
			},
			wantError: true,
			errorMsg:  "Invalid Scope",
		},
		{
			name:  "allowed",
			scope: "test_scope",
			mockSetup: func() {
				db.On("IsAllowed", ctx, models.GlobalEntity, "test_scope").
					Return(true, true, nil).Once()
			},
			wantAllowed: true,
		},
		{
			name:  "denied",
			scope: "test_scope",
			mockSetup: func() {
				db.On("IsAllowed", ctx, models.GlobalEntity, "test_scope").
					Return(true, false, nil).Once()
			},
			wantAllowed: false,
		},
		{
			name:  "database error",
			scope: "test_scope",
			mockSetup: func() {
				db.On("IsAllowed", ctx, models.GlobalEntity, "test_scope").
					Return(false, false, assert.AnError).Once()
			},
			wantError: true,
			errorMsg:  "Query Error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			res, err := server.GlobalAllowed(ctx, &g.GlobalAllowedRequest{
				Scope: tt.scope,
			})

			if tt.wantError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
				assert.Nil(t, res)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.Equal(t, tt.wantAllowed, res.Allowed)
		})
	}
}
