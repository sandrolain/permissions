package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/sandrolain/permissions/pkg/grpc"
)

// This example shows advanced patterns for using the Permissions service:
// - Hierarchical scope management
// - Permission inheritance patterns
// - Batch operations
// - Best practices for error handling

// PermissionsClient encapsulates connection logic and provides utility methods
type PermissionsClient struct {
	client pb.PermissionsServiceClient
	conn   *grpc.ClientConn
}

// NewPermissionsClient creates a new client instance
func NewPermissionsClient(address string) (*PermissionsClient, error) {
	conn, err := grpc.Dial(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}

	return &PermissionsClient{
		client: pb.NewPermissionsServiceClient(conn),
		conn:   conn,
	}, nil
}

// Close closes the connection
func (pc *PermissionsClient) Close() error {
	return pc.conn.Close()
}

// SetupHierarchicalRoles creates a hierarchical role structure
func (pc *PermissionsClient) SetupHierarchicalRoles(ctx context.Context) error {
	// Define a role and permission hierarchy
	roleHierarchy := []struct {
		role   string
		scopes []struct {
			scope   string
			allowed bool
		}
	}{
		{
			role: "admin",
			scopes: []struct {
				scope   string
				allowed bool
			}{
				{"*", true}, // Admins have access to everything
			},
		},
		{
			role: "manager",
			scopes: []struct {
				scope   string
				allowed bool
			}{
				{"projects:*", true},
				{"users:view", true},
				{"reports:*", true},
			},
		},
		{
			role: "developer",
			scopes: []struct {
				scope   string
				allowed bool
			}{
				{"projects:view", true},
				{"projects:edit", true},
				{"code:*", true},
			},
		},
	}

	for _, rh := range roleHierarchy {
		for _, s := range rh.scopes {
			_, err := pc.client.SetRoleScope(ctx, &pb.SetRoleScopeRequest{
				Role:    rh.role,
				Scope:   s.scope,
				Allowed: s.allowed,
			})
			if err != nil {
				return err
			}
			log.Printf("Set scope %s for role %s = %v", s.scope, rh.role, s.allowed)
		}
	}

	return nil
}

// SetupProjectPermissions shows how to manage permissions for a specific project
func (pc *PermissionsClient) SetupProjectPermissions(ctx context.Context, projectID string) error {
	// Define project-specific scopes
	projectScopes := []struct {
		role    string
		scope   string
		allowed bool
	}{
		{"project_admin", "projects:" + projectID + ":*", true},
		{"project_member", "projects:" + projectID + ":view", true},
		{"project_member", "projects:" + projectID + ":edit", true},
		{"project_viewer", "projects:" + projectID + ":view", true},
	}

	for _, ps := range projectScopes {
		_, err := pc.client.SetRoleScope(ctx, &pb.SetRoleScopeRequest{
			Role:    ps.role,
			Scope:   ps.scope,
			Allowed: ps.allowed,
		})
		if err != nil {
			return err
		}
		log.Printf("Set scope %s for role %s = %v", ps.scope, ps.role, ps.allowed)
	}

	return nil
}

// CheckUserAccess performs a comprehensive permission check
func (pc *PermissionsClient) CheckUserAccess(ctx context.Context, user, scope string) (bool, error) {
	// First check for user-specific permission
	userAllowed, err := pc.client.UserAllowed(ctx, &pb.UserAllowedRequest{
		User:  user,
		Scope: scope,
	})
	if err != nil {
		return false, err
	}
	if userAllowed.Allowed {
		return true, nil
	}

	// Then check permissions based on user roles
	roles, err := pc.client.GetUserRoles(ctx, &pb.GetUserRolesRequest{
		User: user,
	})
	if err != nil {
		return false, err
	}

	for _, role := range roles.Roles {
		roleAllowed, err := pc.client.RoleAllowed(ctx, &pb.RoleAllowedRequest{
			Role:  role,
			Scope: scope,
		})
		if err != nil {
			return false, err
		}
		if roleAllowed.Allowed {
			return true, nil
		}
	}

	// Finally check global permissions
	globalAllowed, err := pc.client.GlobalAllowed(ctx, &pb.GlobalAllowedRequest{
		Scope: scope,
	})
	if err != nil {
		return false, err
	}

	return globalAllowed.Allowed, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create an instance of our client
	client, err := NewPermissionsClient("localhost:50051")
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer client.Close()

	// Setup role hierarchy
	if err := client.SetupHierarchicalRoles(ctx); err != nil {
		log.Fatalf("Error in role setup: %v", err)
	}

	// Setup permissions for a specific project
	projectID := "proj123"
	if err := client.SetupProjectPermissions(ctx, projectID); err != nil {
		log.Fatalf("Error in project permissions setup: %v", err)
	}

	// Example permission checks for different scenarios
	testCases := []struct {
		user  string
		scope string
	}{
		{"admin_user", "projects:proj123:edit"},
		{"developer_user", "code:commit"},
		{"viewer_user", "projects:proj123:view"},
	}

	for _, tc := range testCases {
		allowed, err := client.CheckUserAccess(ctx, tc.user, tc.scope)
		if err != nil {
			log.Printf("Error checking access for %s to %s: %v", tc.user, tc.scope, err)
			continue
		}
		log.Printf("Access for %s to %s: %v", tc.user, tc.scope, allowed)
	}
}
