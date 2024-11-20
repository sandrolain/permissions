#!/bin/bash

# Set the working directory to the localtest directory
cd "$(dirname "$0")/../localtest" || exit

# Function to display usage
show_usage() {
    echo "Usage: $0 [command]"
    echo "Commands:"
    echo "  start    - Start the local test environment"
    echo "  stop     - Stop the local test environment"
    echo "  restart  - Restart the local test environment"
    echo "  logs     - Show logs from all services"
    echo "  clean    - Stop and remove all containers, networks, and volumes"
    echo "  status   - Show status of all services"
}

# Function to start services
start_services() {
    echo "Starting local test environment..."
    docker compose up -d --build
    echo "Waiting for services to be ready..."
    sleep 5
    docker compose ps
}

# Function to stop services
stop_services() {
    echo "Stopping local test environment..."
    docker compose down
}

# Function to show logs
show_logs() {
    docker compose logs -f
}

# Function to clean everything
clean_environment() {
    echo "Cleaning up local test environment..."
    docker compose down -v
    echo "Removing any leftover volumes..."
    docker volume rm permissions-pgdata 2>/dev/null || true
}

# Function to show status
show_status() {
    echo "Current status of services:"
    docker compose ps
}

# Main script logic
case "$1" in
    "start")
        start_services
        ;;
    "stop")
        stop_services
        ;;
    "restart")
        stop_services
        start_services
        ;;
    "logs")
        show_logs
        ;;
    "clean")
        clean_environment
        ;;
    "status")
        show_status
        ;;
    *)
        show_usage
        exit 1
        ;;
esac
