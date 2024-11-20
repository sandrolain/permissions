# Quick Start Guide

This guide will help you get started with the Permissions Service API quickly.

## Prerequisites

Before you begin, ensure you have:
- Go 1.23 or higher installed
- Docker (for containerized deployment)
- PostgreSQL 15 or higher
- Protocol buffer compiler (protoc)
- buf tool for protocol buffer management

## Installation

1. Clone the repository:
```bash
git clone https://github.com/sandrolain/permissions.git
cd permissions
```

2. Install dependencies:
```bash
go mod download
```

3. Build the service:
```bash
./build.sh
```

## Running the Service

1. Using Docker:
```bash
docker build -t permissions-service .
docker run -p 50051:50051 permissions-service
```

2. Locally:
```bash
./start.sh
```

## Basic Usage Examples

Here are some common use cases to get you started:

### 1. Managing User Roles

```go
package main

import (
    "context"
    "log"
    
    "google.golang.org/grpc"
    pb "github.com/sandrolain/permissions/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    
    client := pb.NewPermissionsServiceClient(conn)
    
    // Assign a role to a user
    setRoleResp, err := client.SetUserRole(context.Background(), &pb.SetUserRoleRequest{
        User: "john_doe",
        Role: "editor",
    })
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Role assigned: %v", setRoleResp.Affected)
    
    // Get user's roles
    getRolesResp, err := client.GetUserRoles(context.Background(), &pb.GetUserRolesRequest{
        User: "john_doe",
    })
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("User roles: %v", getRolesResp.Roles)
}
```

### 2. Managing Permissions

```go
// Set a role-specific permission
setRoleScopeResp, err := client.SetRoleScope(context.Background(), &pb.SetRoleScopeRequest{
    Role:    "editor",
    Scope:   "documents:write",
    Allowed: true,
})

// Check if a user has permission
allowed, err := client.UserAllowed(context.Background(), &pb.UserAllowedRequest{
    User:  "john_doe",
    Scope: "documents:write",
})
```

## Common Patterns

### 1. Hierarchical Scopes

Scopes can be hierarchical using colons:
```
documents:read
documents:write
admin:*
```

### 2. Role-Based Access Control (RBAC)

1. Create roles with specific permissions:
```go
// Set up editor role
client.SetRoleScope(ctx, &pb.SetRoleScopeRequest{
    Role:    "editor",
    Scope:   "documents:write",
    Allowed: true,
})

// Set up viewer role
client.SetRoleScope(ctx, &pb.SetRoleScopeRequest{
    Role:    "viewer",
    Scope:   "documents:read",
    Allowed: true,
})
```

2. Assign roles to users:
```go
client.SetUserRole(ctx, &pb.SetUserRoleRequest{
    User: "john_doe",
    Role: "editor",
})
```

## Best Practices

1. **Scope Naming**
   - Use lowercase letters, numbers, and underscores
   - Use colons for hierarchy
   - Be specific but not too granular

2. **Role Management**
   - Create roles based on job functions
   - Assign minimum required permissions
   - Regularly audit role assignments

3. **Error Handling**
   - Always check for errors in responses
   - Implement proper logging
   - Handle permission denied scenarios gracefully

## Next Steps

- Read the full [API Reference](api-reference.md) for detailed endpoint documentation
- Check out the integration tests in `internal/integration` for more usage examples
- Join our community for support and discussions
