#!/bin/bash

# Source common functions
source "$(dirname "$0")/grpc-common.sh"

# Function to get user roles
get_user_roles() {
    local user="$1"
    if [ -z "$user" ]; then
        echo "Usage: $0 get-roles <user>"
        exit 1
    fi
    
    grpcurl -plaintext -d "{\"user\": \"$user\"}" \
        "$GRPC_HOST" \
        "$SERVICE/GetUserRoles" | format_json
}

# Function to set user role
set_user_role() {
    local user="$1"
    local role="$2"
    if [ -z "$user" ] || [ -z "$role" ]; then
        echo "Usage: $0 set-role <user> <role>"
        exit 1
    fi
    
    grpcurl -plaintext -d "{\"user\": \"$user\", \"role\": \"$role\"}" \
        "$GRPC_HOST" \
        "$SERVICE/SetUserRole" | format_json
}

# Function to unset user role
unset_user_role() {
    local user="$1"
    local role="$2"
    if [ -z "$user" ] || [ -z "$role" ]; then
        echo "Usage: $0 unset-role <user> <role>"
        exit 1
    fi
    
    grpcurl -plaintext -d "{\"user\": \"$user\", \"role\": \"$role\"}" \
        "$GRPC_HOST" \
        "$SERVICE/UnsetUserRole" | format_json
}

# Main script logic
case "$1" in
    "get-roles")
        get_user_roles "$2"
        ;;
    "set-role")
        set_user_role "$2" "$3"
        ;;
    "unset-role")
        unset_user_role "$2" "$3"
        ;;
    *)
        echo "Usage: $0 <command> [arguments]"
        echo "Commands:"
        echo "  get-roles <user>            - Get roles for a user"
        echo "  set-role <user> <role>      - Assign a role to a user"
        echo "  unset-role <user> <role>    - Remove a role from a user"
        exit 1
        ;;
esac
