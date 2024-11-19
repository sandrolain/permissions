#!/bin/bash

set -e

# Run all tests with coverage
go test -v -race -coverprofile=coverage.out ./...

# Display coverage report
go tool cover -func=coverage.out

# Generate an HTML report
go tool cover -html=coverage.out -o coverage.html

# Open the HTML report in the default web browser
open coverage.html
