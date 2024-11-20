# Permissions Service API Reference

This document provides detailed information about the gRPC-based Permissions Service API.

## Overview

The Permissions Service provides a comprehensive API for managing user roles, scopes, and permissions in a distributed system. The service is built using gRPC and Protocol Buffers, offering high performance and type safety.

## Authentication

The service expects authentication details to be provided via gRPC metadata. Specific authentication requirements should be configured based on your deployment environment.

## API Endpoints

### User Role Management

#### GetUserRoles

Retrieves all roles assigned to a specific user.

- **Request** (`GetUserRolesRequest`):
  - `user` (string): User identifier (must match pattern: `^[A-Za-z0-9_-]+$`)
- **Response** (`GetUserRolesResponse`):
  - `roles` (array of string): List of roles assigned to the user

#### SetUserRole

Assigns a role to a user.

- **Request** (`SetUserRoleRequest`):
  - `user` (string): User identifier
  - `role` (string): Role to assign
- **Response** (`SetUserRoleResponse`):
  - `affected` (boolean): Whether the operation changed anything
  - `roles` (array of string): Updated list of user's roles

#### UnsetUserRole

Removes a role from a user.

- **Request** (`UnsetUserRoleRequest`):
  - `user` (string): User identifier
  - `role` (string): Role to remove
- **Response** (`UnsetUserRoleResponse`):
  - `affected` (boolean): Whether the operation changed anything
  - `roles` (array of string): Updated list of user's roles

### Scope Management

#### Global Scopes

1. **GetGlobalScopes**
   - **Request** (`GetGlobalScopesRequest`):
     - `scope_pattern` (string): Pattern to filter scopes (supports wildcards)
   - **Response** (`GetGlobalScopesResponse`):
     - `scopes` (array of ScopeItem): List of matching scope items

2. **SetGlobalScope**
   - **Request** (`SetGlobalScopeRequest`):
     - `scope` (string): Scope identifier
     - `allowed` (boolean): Permission status
   - **Response** (`SetGlobalScopeResponse`):
     - `affected` (boolean): Whether the operation changed anything

3. **UnsetGlobalScope**
   - **Request** (`UnsetGlobalScopeRequest`):
     - `scope` (string): Scope to remove
   - **Response** (`UnsetGlobalScopeResponse`):
     - `affected` (boolean): Whether the operation changed anything

#### Role Scopes

1. **GetRoleScopes**
   - **Request** (`GetRoleScopesRequest`):
     - `role` (string): Role identifier
     - `scope_pattern` (string): Pattern to filter scopes
   - **Response** (`GetRoleScopesResponse`):
     - `scopes` (array of ScopeItem): List of matching scope items

2. **SetRoleScope**
   - **Request** (`SetRoleScopeRequest`):
     - `role` (string): Role identifier
     - `scope` (string): Scope identifier
     - `allowed` (boolean): Permission status
   - **Response** (`SetRoleScopeResponse`):
     - `affected` (boolean): Whether the operation changed anything

3. **UnsetRoleScope**
   - **Request** (`UnsetRoleScopeRequest`):
     - `role` (string): Role identifier
     - `scope` (string): Scope to remove
   - **Response** (`UnsetRoleScopeResponse`):
     - `affected` (boolean): Whether the operation changed anything

#### User Scopes

1. **GetUserScopes**
   - **Request** (`GetUserScopesRequest`):
     - `user` (string): User identifier
     - `scope_pattern` (string): Pattern to filter scopes
   - **Response** (`GetUserScopesResponse`):
     - `scopes` (array of ScopeItem): List of matching scope items

2. **SetUserScope**
   - **Request** (`SetUserScopeRequest`):
     - `user` (string): User identifier
     - `scope` (string): Scope identifier
     - `allowed` (boolean): Permission status
   - **Response** (`SetUserScopeResponse`):
     - `affected` (boolean): Whether the operation changed anything

3. **UnsetUserScope**
   - **Request** (`UnsetUserScopeRequest`):
     - `user` (string): User identifier
     - `scope` (string): Scope to remove
   - **Response** (`UnsetUserScopeResponse`):
     - `affected` (boolean): Whether the operation changed anything

### Permission Checking

#### GlobalAllowed

Check if a scope is allowed globally.

- **Request** (`GlobalAllowedRequest`):
  - `scope` (string): Scope to check
- **Response** (`GlobalAllowedResponse`):
  - `allowed` (boolean): Whether the scope is allowed

#### RoleAllowed

Check if a scope is allowed for a specific role.

- **Request** (`RoleAllowedRequest`):
  - `role` (string): Role identifier
  - `scope` (string): Scope to check
- **Response** (`RoleAllowedResponse`):
  - `allowed` (boolean): Whether the scope is allowed

#### UserAllowed

Check if a scope is allowed for a specific user.

- **Request** (`UserAllowedRequest`):
  - `user` (string): User identifier
  - `scope` (string): Scope to check
- **Response** (`UserAllowedResponse`):
  - `allowed` (boolean): Whether the scope is allowed

## Data Types

### ScopeItem

Represents a scope and its permission status:
- `scope` (string): Scope identifier (must match pattern: `^[A-Za-z0-9_-]+$`)
- `allowed` (boolean): Whether the scope is allowed

## Input Validation

All string inputs must follow these validation rules:
- User and role identifiers: `^[A-Za-z0-9_-]+$`
- Scope patterns: `^[A-Za-z0-9:/*_-]+$`
- All string fields must have a minimum length of 1

## Error Handling

The service follows standard gRPC error handling practices. Common error codes:

- `INVALID_ARGUMENT`: Input validation failed
- `NOT_FOUND`: Requested resource doesn't exist
- `PERMISSION_DENIED`: Authentication or authorization failed
- `INTERNAL`: Internal server error

## Usage Examples

Here are some example gRPC calls using different programming languages:

### Go Example
```go
client := permissionsGrpc.NewPermissionsServiceClient(conn)

// Check if user has permission
resp, err := client.UserAllowed(context.Background(), &permissionsGrpc.UserAllowedRequest{
    User:  "john_doe",
    Scope: "documents:read",
})
if err != nil {
    log.Fatal(err)
}
if resp.Allowed {
    // User has permission
}
```

### Python Example
```python
stub = permissions_pb2_grpc.PermissionsServiceStub(channel)

# Assign role to user
response = stub.SetUserRole(
    permissions_pb2.SetUserRoleRequest(
        user="john_doe",
        role="editor"
    )
)
if response.affected:
    print("Role assigned successfully")
```
