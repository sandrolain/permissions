# Permission Service gRPC Test Scripts

This directory contains helper scripts to interact with the Permission Service gRPC API during local testing.

## Prerequisites

- [grpcurl](https://github.com/fullstorydev/grpcurl) - A command-line tool for interacting with gRPC servers
  ```bash
  brew install grpcurl
  ```
- [jq](https://stedolan.github.io/jq/) (optional) - For better JSON formatting
  ```bash
  brew install jq
  ```

## Getting Started

1. Start the local test environment from the project root:
   ```bash
   ./scripts/local-test.sh start
   ```

2. Wait for the services to be ready. The scripts will automatically wait for the gRPC server.

3. Navigate to the scripts directory:
   ```bash
   cd localtest/scripts
   ```

## Available Scripts

### Managing Roles (`roles.sh`)

```bash
# Get roles for a user
./roles.sh get-roles <user>

# Assign a role to a user
./roles.sh set-role <user> <role>

# Remove a role from a user
./roles.sh unset-role <user> <role>

# Examples:
./roles.sh set-role "john" "admin"
./roles.sh get-roles "john"
./roles.sh unset-role "john" "admin"
```

### Managing Scopes (`scopes.sh`)

```bash
# Get scopes
./scopes.sh get-global [pattern]              # Get global scopes
./scopes.sh get-role-scopes <role> [pattern]  # Get scopes for a role
./scopes.sh get-user-scopes <user> [pattern]  # Get scopes for a user

# Set scopes (allowed defaults to true)
./scopes.sh set-global <scope> [allowed]      # Set global scope
./scopes.sh set-role <role> <scope> [allowed] # Set role scope
./scopes.sh set-user <user> <scope> [allowed] # Set user scope

# Check permissions
./scopes.sh check-global <scope>              # Check if scope is globally allowed
./scopes.sh check-role <role> <scope>         # Check if scope is allowed for role
./scopes.sh check-user <user> <scope>         # Check if scope is allowed for user

# Examples:
./scopes.sh set-global "public:read" true
./scopes.sh set-role "admin" "users:write" true
./scopes.sh check-user "john" "users:write"
```

## Common Usage Patterns

### Setting Up an Admin User

```bash
# Create admin role with permissions
./roles.sh set-role "john" "admin"
./scopes.sh set-role "admin" "users:*" true
./scopes.sh set-role "admin" "roles:*" true
./scopes.sh set-role "admin" "scopes:*" true

# Verify permissions
./roles.sh get-roles "john"
./scopes.sh get-role-scopes "admin"
./scopes.sh check-user "john" "users:write"
```

### Setting Up a Regular User

```bash
# Create user with limited permissions
./roles.sh set-role "jane" "user"
./scopes.sh set-role "user" "profile:read" true
./scopes.sh set-role "user" "profile:write" true

# Verify permissions
./scopes.sh check-user "jane" "profile:read"
./scopes.sh check-user "jane" "users:write"  # Should return false
```

### Managing Global Permissions

```bash
# Set up public access
./scopes.sh set-global "public:read" true
./scopes.sh set-global "public:write" false

# Check global permissions
./scopes.sh get-global
./scopes.sh check-global "public:read"
```

## Troubleshooting

1. If grpcurl is not installed:
   ```bash
   brew install grpcurl
   ```

2. If you see connection errors:
   - Make sure the local test environment is running (`./scripts/local-test.sh start`)
   - Check if the service is healthy (`./scripts/local-test.sh status`)
   - Try restarting the environment (`./scripts/local-test.sh restart`)

3. If you see "invalid JSON" errors:
   - Make sure your scope and role names follow the pattern: `^(?:[A-Za-z0-9_-]+(?:[:/][A-Za-z0-9_-]+)*)$`
   - Examples of valid names: `admin`, `users:read`, `api:v1:write`

## Notes

- All scripts use the gRPC server at `localhost:9090`
- JSON output is automatically formatted if `jq` is installed
- Scope and role names must match the pattern: `^(?:[A-Za-z0-9_-]+(?:[:/][A-Za-z0-9_-]+)*)$`
- Boolean values for `allowed` parameter can be `true` or `false`
