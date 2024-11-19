package svcgrpc

import (
	"context"
	"testing"

	g "github.com/sandrolain/permissions/pkg/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestInvalidUserRole(t *testing.T) {
	db := NewTestDBService()
	srv := NewGrpcServer(db)

	tests := []struct {
		name    string
		user    string
		role    string
		wantErr bool
	}{
		{
			name:    "invalid user with special chars",
			user:    "user@123",
			role:    "admin",
			wantErr: true,
		},
		{
			name:    "invalid role with special chars",
			user:    "user123",
			role:    "admin@123",
			wantErr: true,
		},
		{
			name:    "empty user",
			user:    "",
			role:    "admin",
			wantErr: true,
		},
		{
			name:    "empty role",
			user:    "user123",
			role:    "",
			wantErr: true,
		},
		{
			name:    "valid user and role",
			user:    "user123",
			role:    "admin",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test SetUserRole
			setReq := &g.SetUserRoleRequest{
				User: tt.user,
				Role: tt.role,
			}
			_, err := srv.SetUserRole(context.Background(), setReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetUserRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && status.Code(err) != codes.InvalidArgument {
				t.Errorf("SetUserRole() expected InvalidArgument error, got %v", err)
				return
			}

			// Test GetUserRoles only when the user is invalid
			if !validEntity(tt.user) {
				getRolesReq := &g.GetUserRolesRequest{
					User: tt.user,
				}
				_, err = srv.GetUserRoles(context.Background(), getRolesReq)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetUserRoles() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if err != nil && status.Code(err) != codes.InvalidArgument {
					t.Errorf("GetUserRoles() expected InvalidArgument error, got %v", err)
					return
				}
			}

			// Test UnsetUserRole
			unsetReq := &g.UnsetUserRoleRequest{
				User: tt.user,
				Role: tt.role,
			}
			_, err = srv.UnsetUserRole(context.Background(), unsetReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnsetUserRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && status.Code(err) != codes.InvalidArgument {
				t.Errorf("UnsetUserRole() expected InvalidArgument error, got %v", err)
				return
			}
		})
	}
}

func TestInvalidScopes(t *testing.T) {
	db := NewTestDBService()
	srv := NewGrpcServer(db)

	tests := []struct {
		name         string
		entity       string
		scope        string
		pattern      string
		wantSetErr   bool
		wantGetErr   bool
		wantUnsetErr bool
	}{
		{
			name:         "invalid scope with special chars",
			entity:       "user123",
			scope:        "read@write",
			pattern:      "",
			wantSetErr:   true,
			wantGetErr:   false,
			wantUnsetErr: true,
		},
		{
			name:         "empty scope",
			entity:       "user123",
			scope:        "",
			pattern:      "",
			wantSetErr:   true,
			wantGetErr:   false,
			wantUnsetErr: true,
		},
		{
			name:         "invalid pattern",
			entity:       "user123",
			scope:        "read",
			pattern:      "@*",
			wantSetErr:   false,
			wantGetErr:   true,
			wantUnsetErr: false,
		},
		{
			name:         "valid inputs",
			entity:       "user123",
			scope:        "read",
			pattern:      "read*",
			wantSetErr:   false,
			wantGetErr:   false,
			wantUnsetErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test SetUserScope
			setReq := &g.SetUserScopeRequest{
				User:    tt.entity,
				Scope:   tt.scope,
				Allowed: true,
			}
			_, err := srv.SetUserScope(context.Background(), setReq)
			if (err != nil) != tt.wantSetErr {
				t.Errorf("SetUserScope() error = %v, wantSetErr %v", err, tt.wantSetErr)
				return
			}
			if err != nil && status.Code(err) != codes.InvalidArgument {
				t.Errorf("SetUserScope() expected InvalidArgument error, got %v", err)
				return
			}

			// Test GetUserScopes
			getScopesReq := &g.GetUserScopesRequest{
				User:         tt.entity,
				ScopePattern: tt.pattern,
			}
			_, err = srv.GetUserScopes(context.Background(), getScopesReq)
			if (err != nil) != tt.wantGetErr {
				t.Errorf("GetUserScopes() error = %v, wantGetErr %v", err, tt.wantGetErr)
				return
			}
			if err != nil && status.Code(err) != codes.InvalidArgument {
				t.Errorf("GetUserScopes() expected InvalidArgument error, got %v", err)
				return
			}

			// Test UnsetUserScope
			unsetReq := &g.UnsetUserScopeRequest{
				User:  tt.entity,
				Scope: tt.scope,
			}
			_, err = srv.UnsetUserScope(context.Background(), unsetReq)
			if (err != nil) != tt.wantUnsetErr {
				t.Errorf("UnsetUserScope() error = %v, wantUnsetErr %v", err, tt.wantUnsetErr)
				return
			}
			if err != nil && status.Code(err) != codes.InvalidArgument {
				t.Errorf("UnsetUserScope() expected InvalidArgument error, got %v", err)
				return
			}
		})
	}
}
