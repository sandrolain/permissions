#!/bin/bash
set -e

# Build the test image
docker build -t permissions-service:test .

# Run the integration tests
go test ./internal/integration -v
