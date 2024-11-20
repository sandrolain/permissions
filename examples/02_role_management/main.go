package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/sandrolain/permissions/pkg/grpc"
)

// This example shows how to manage user roles:
// - Assign roles to users
// - Remove roles from users
// - Verify assigned roles

func main() {
	// Set a timeout for the context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create a gRPC connection
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPermissionsServiceClient(conn)

	// 1. Assign roles to a user
	user := "john_doe"
	roles := []string{"editor", "viewer"}

	for _, role := range roles {
		resp, err := client.SetUserRole(ctx, &pb.SetUserRoleRequest{
			User: user,
			Role: role,
		})
		if err != nil {
			log.Fatalf("Error assigning role %s: %v", role, err)
		}
		if resp.Affected {
			log.Printf("Role %s successfully assigned to %s", role, user)
		}
	}

	// 2. Check assigned roles
	getRolesResp, err := client.GetUserRoles(ctx, &pb.GetUserRolesRequest{
		User: user,
	})
	if err != nil {
		log.Fatalf("Error retrieving roles: %v", err)
	}
	log.Printf("Current roles for %s: %v", user, getRolesResp.Roles)

	// 3. Remove a role
	unsetResp, err := client.UnsetUserRole(ctx, &pb.UnsetUserRoleRequest{
		User: user,
		Role: "viewer",
	})
	if err != nil {
		log.Fatalf("Error removing role: %v", err)
	}
	if unsetResp.Affected {
		log.Printf("Role 'viewer' successfully removed from %s", user)
	}

	// 4. Check roles again after removal
	finalRolesResp, err := client.GetUserRoles(ctx, &pb.GetUserRolesRequest{
		User: user,
	})
	if err != nil {
		log.Fatalf("Error retrieving final roles: %v", err)
	}
	log.Printf("Final roles for %s: %v", user, finalRolesResp.Roles)

	// Error handling example
	_, err = client.SetUserRole(ctx, &pb.SetUserRoleRequest{
		User: "", // Empty username, should generate an error
		Role: "editor",
	})
	if err != nil {
		log.Printf("Expected error for invalid input: %v", err)
	}
}
