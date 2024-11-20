package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/sandrolain/permissions/pkg/grpc"
)

// This example shows how to manage scopes (permissions) at various levels:
// - Global scopes
// - Role scopes
// - User scopes
// Also includes permission verification examples

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPermissionsServiceClient(conn)

	// 1. Global scope management
	log.Println("=== Global Scope Management ===")
	
	// Set some global scopes
	globalScopes := []struct {
		scope   string
		allowed bool
	}{
		{"documents:read", true},
		{"documents:list", true},
		{"admin:*", false},
	}

	for _, s := range globalScopes {
		_, err := client.SetGlobalScope(ctx, &pb.SetGlobalScopeRequest{
			Scope:   s.scope,
			Allowed: s.allowed,
		})
		if err != nil {
			log.Printf("Error setting global scope %s: %v", s.scope, err)
			continue
		}
		log.Printf("Global scope set: %s = %v", s.scope, s.allowed)
	}

	// Retrieve and verify global scopes
	globalResp, err := client.GetGlobalScopes(ctx, &pb.GetGlobalScopesRequest{
		ScopePattern: "documents:*",
	})
	if err != nil {
		log.Printf("Error retrieving global scopes: %v", err)
	} else {
		log.Printf("Global scopes found for pattern 'documents:*': %v", globalResp.Scopes)
	}

	// 2. Role scope management
	log.Println("\n=== Role Scope Management ===")
	
	roleScopes := []struct {
		role    string
		scope   string
		allowed bool
	}{
		{"editor", "documents:write", true},
		{"editor", "documents:delete", true},
		{"viewer", "documents:read", true},
	}

	for _, rs := range roleScopes {
		_, err := client.SetRoleScope(ctx, &pb.SetRoleScopeRequest{
			Role:    rs.role,
			Scope:   rs.scope,
			Allowed: rs.allowed,
		})
		if err != nil {
			log.Printf("Error setting role scope for %s: %v", rs.role, err)
			continue
		}
		log.Printf("Role scope set: %s.%s = %v", rs.role, rs.scope, rs.allowed)
	}

	// Check scopes for a specific role
	roleResp, err := client.GetRoleScopes(ctx, &pb.GetRoleScopesRequest{
		Role:         "editor",
		ScopePattern: "documents:*",
	})
	if err != nil {
		log.Printf("Error retrieving role scopes: %v", err)
	} else {
		log.Printf("Scopes found for 'editor' role: %v", roleResp.Scopes)
	}

	// 3. User scope management
	log.Println("\n=== User Scope Management ===")
	
	// Set some user-specific scopes
	userScopes := []struct {
		user    string
		scope   string
		allowed bool
	}{
		{"john_doe", "projects:123:edit", true},
		{"john_doe", "projects:456:view", true},
	}

	for _, us := range userScopes {
		_, err := client.SetUserScope(ctx, &pb.SetUserScopeRequest{
			User:    us.user,
			Scope:   us.scope,
			Allowed: us.allowed,
		})
		if err != nil {
			log.Printf("Error setting user scope for %s: %v", us.user, err)
			continue
		}
		log.Printf("User scope set: %s.%s = %v", us.user, us.scope, us.allowed)
	}

	// 4. Permission verification
	log.Println("\n=== Permission Verification ===")
	
	// Check various types of permissions
	checksToPerform := []struct {
		description string
		check      func() (bool, error)
	}{
		{
			"Global permission documents:read",
			func() (bool, error) {
				resp, err := client.GlobalAllowed(ctx, &pb.GlobalAllowedRequest{
					Scope: "documents:read",
				})
				return resp.GetAllowed(), err
			},
		},
		{
			"Editor role permission for documents:write",
			func() (bool, error) {
				resp, err := client.RoleAllowed(ctx, &pb.RoleAllowedRequest{
					Role:  "editor",
					Scope: "documents:write",
				})
				return resp.GetAllowed(), err
			},
		},
		{
			"User john_doe permission for projects:123:edit",
			func() (bool, error) {
				resp, err := client.UserAllowed(ctx, &pb.UserAllowedRequest{
					User:  "john_doe",
					Scope: "projects:123:edit",
				})
				return resp.GetAllowed(), err
			},
		},
	}

	for _, check := range checksToPerform {
		allowed, err := check.check()
		if err != nil {
			log.Printf("Error in verification: %s - %v", check.description, err)
		} else {
			log.Printf("Verification: %s = %v", check.description, allowed)
		}
	}
}
