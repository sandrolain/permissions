#!/bin/bash
set -e

# Clean test before execution
go clean -testcache

# Build the test image
docker build -t permissions-service:test .

# Run the integration tests
go test ./internal/integration -v
