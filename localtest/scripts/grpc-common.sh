#!/bin/bash

GRPC_HOST="localhost:9090"
PROTO_PATH="../pkg/grpc/permissions.proto"
SERVICE="permissions.PermissionsService"

# Function to check if grpcurl is installed
check_grpcurl() {
    if ! command -v grpcurl &> /dev/null; then
        echo "grpcurl is not installed. Please install it first:"
        echo "brew install grpcurl"
        exit 1
    fi
}

# Function to wait for the gRPC server to be ready
wait_for_grpc() {
    echo "Waiting for gRPC server to be ready..."
    until grpcurl -plaintext "$GRPC_HOST" list > /dev/null 2>&1; do
        sleep 1
    done
    echo "gRPC server is ready!"
}

# Function to format JSON output
format_json() {
    if command -v jq &> /dev/null; then
        jq '.'
    else
        cat
    fi
}

# Check if grpcurl is installed at script start
check_grpcurl
