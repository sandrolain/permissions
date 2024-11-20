#!/bin/bash

# Source common functions
source "$(dirname "$0")/grpc-common.sh"

# Function to get global scopes
get_global_scopes() {
    local pattern="${1:-*}"
    grpcurl -plaintext -d "{\"scope_pattern\": \"$pattern\"}" \
        "$GRPC_HOST" \
        "$SERVICE/GetGlobalScopes" | format_json
}

# Function to get role scopes
get_role_scopes() {
    local role="$1"
    local pattern="${2:-*}"
    if [ -z "$role" ]; then
        echo "Usage: $0 get-role-scopes <role> [pattern]"
        exit 1
    fi
    
    grpcurl -plaintext -d "{\"role\": \"$role\", \"scope_pattern\": \"$pattern\"}" \
        "$GRPC_HOST" \
        "$SERVICE/GetRoleScopes" | format_json
}

# Function to get user scopes
get_user_scopes() {
    local user="$1"
    local pattern="${2:-*}"
    if [ -z "$user" ]; then
        echo "Usage: $0 get-user-scopes <user> [pattern]"
        exit 1
    fi
    
    grpcurl -plaintext -d "{\"user\": \"$user\", \"scope_pattern\": \"$pattern\"}" \
        "$GRPC_HOST" \
        "$SERVICE/GetUserScopes" | format_json
}

# Function to set global scope
set_global_scope() {
    local scope="$1"
    local allowed="${2:-true}"
    if [ -z "$scope" ]; then
        echo "Usage: $0 set-global <scope> [allowed]"
        exit 1
    fi
    
    grpcurl -plaintext -d "{\"scope\": \"$scope\", \"allowed\": $allowed}" \
        "$GRPC_HOST" \
        "$SERVICE/SetGlobalScope" | format_json
}

# Function to set role scope
set_role_scope() {
    local role="$1"
    local scope="$2"
    local allowed="${3:-true}"
    if [ -z "$role" ] || [ -z "$scope" ]; then
        echo "Usage: $0 set-role <role> <scope> [allowed]"
        exit 1
    fi
    
    grpcurl -plaintext -d "{\"role\": \"$role\", \"scope\": \"$scope\", \"allowed\": $allowed}" \
        "$GRPC_HOST" \
        "$SERVICE/SetRoleScope" | format_json
}

# Function to set user scope
set_user_scope() {
    local user="$1"
    local scope="$2"
    local allowed="${3:-true}"
    if [ -z "$user" ] || [ -z "$scope" ]; then
        echo "Usage: $0 set-user <user> <scope> [allowed]"
        exit 1
    fi
    
    grpcurl -plaintext -d "{\"user\": \"$user\", \"scope\": \"$scope\", \"allowed\": $allowed}" \
        "$GRPC_HOST" \
        "$SERVICE/SetUserScope" | format_json
}

# Function to check global allowed
check_global_allowed() {
    local scope="$1"
    if [ -z "$scope" ]; then
        echo "Usage: $0 check-global <scope>"
        exit 1
    fi
    
    grpcurl -plaintext -d "{\"scope\": \"$scope\"}" \
        "$GRPC_HOST" \
        "$SERVICE/GlobalAllowed" | format_json
}

# Function to check role allowed
check_role_allowed() {
    local role="$1"
    local scope="$2"
    if [ -z "$role" ] || [ -z "$scope" ]; then
        echo "Usage: $0 check-role <role> <scope>"
        exit 1
    fi
    
    grpcurl -plaintext -d "{\"role\": \"$role\", \"scope\": \"$scope\"}" \
        "$GRPC_HOST" \
        "$SERVICE/RoleAllowed" | format_json
}

# Function to check user allowed
check_user_allowed() {
    local user="$1"
    local scope="$2"
    if [ -z "$user" ] || [ -z "$scope" ]; then
        echo "Usage: $0 check-user <user> <scope>"
        exit 1
    fi
    
    grpcurl -plaintext -d "{\"user\": \"$user\", \"scope\": \"$scope\"}" \
        "$GRPC_HOST" \
        "$SERVICE/UserAllowed" | format_json
}

# Main script logic
case "$1" in
    "get-global")
        get_global_scopes "$2"
        ;;
    "get-role-scopes")
        get_role_scopes "$2" "$3"
        ;;
    "get-user-scopes")
        get_user_scopes "$2" "$3"
        ;;
    "set-global")
        set_global_scope "$2" "$3"
        ;;
    "set-role")
        set_role_scope "$2" "$3" "$4"
        ;;
    "set-user")
        set_user_scope "$2" "$3" "$4"
        ;;
    "check-global")
        check_global_allowed "$2"
        ;;
    "check-role")
        check_role_allowed "$2" "$3"
        ;;
    "check-user")
        check_user_allowed "$2" "$3"
        ;;
    *)
        echo "Usage: $0 <command> [arguments]"
        echo "Commands:"
        echo "  get-global [pattern]                - Get global scopes"
        echo "  get-role-scopes <role> [pattern]    - Get scopes for a role"
        echo "  get-user-scopes <user> [pattern]    - Get scopes for a user"
        echo "  set-global <scope> [allowed]        - Set global scope"
        echo "  set-role <role> <scope> [allowed]   - Set role scope"
        echo "  set-user <user> <scope> [allowed]   - Set user scope"
        echo "  check-global <scope>                - Check if scope is globally allowed"
        echo "  check-role <role> <scope>           - Check if scope is allowed for role"
        echo "  check-user <user> <scope>           - Check if scope is allowed for user"
        exit 1
        ;;
esac
