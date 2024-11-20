package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/sandrolain/permissions/pkg/grpc"
)

// This example shows how to establish a connection with the Permissions service
// and how to perform a simple call to verify the connection

func main() {
	// Set a timeout for the context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a gRPC connection
	// In production, you should use appropriate credentials instead of insecure.NewCredentials()
	conn, err := grpc.Dial("localhost:50051", 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := pb.NewPermissionsServiceClient(conn)

	// Try a simple call to verify the connection
	// Check if a user has a specific role
	resp, err := client.GetUserRoles(ctx, &pb.GetUserRolesRequest{
		User: "test_user",
	})
	if err != nil {
		log.Fatalf("Error in GetUserRoles call: %v", err)
	}

	// Print the user's roles
	log.Printf("Roles for test_user: %v", resp.Roles)
}
