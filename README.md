# Permissions Service

A high-performance gRPC-based service for managing permissions, roles, and scopes in distributed systems.

## Overview

This project provides a scalable and efficient way to manage permissions for users and roles. It utilizes Protocol Buffers (protobuf) for defining the service interface and gRPC for building the service, offering high performance and type safety.

## Features

* User and role management
* Scope management (global, role, and user scopes)
* Permission management (allowed/denied)
* gRPC-based service for efficient communication
* PostgreSQL database support with GORM
* Validation using protobuf validation rules
* Comprehensive test coverage
* Docker support

## Prerequisites

* Go 1.23 or higher
* Docker (for containerized deployment)
* buf (for protocol buffer compilation)
* PostgreSQL 15 or higher

## Project Structure

```
.
├── cmd/                # Main application entry points
├── internal/          # Private application code
│   ├── dbsvc/        # Database service layer
│   └── svcgrpc/      # gRPC service implementation
├── pkg/              # Public API packages
├── .env              # Environment configuration
├── Dockerfile        # Docker build configuration
└── buf.yaml          # Protocol buffer configuration
```

## Getting Started

### Local Development

1. Clone the repository:

   ```bash
   git clone https://github.com/sandrolain/permissions.git
   cd permissions
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Set up PostgreSQL:
   * Install PostgreSQL if not already installed
   * Create a new database for the service
   * Set up database user with appropriate permissions

4. Set up environment variables:

   ```bash
   cp .env.example .env
   # Edit .env with your PostgreSQL configuration:
   # - DB_HOST=localhost
   # - DB_PORT=5432
   # - DB_NAME=permissions
   # - DB_USER=your_user
   # - DB_PASSWORD=your_password
   ```

5. Run the service:

   ```bash
   ./start.sh
   ```

### Docker Deployment

1. Build the Docker image:

   ```bash
   docker build -t permissions:latest .
   ```

2. Run with PostgreSQL:

   ```bash
   docker run -p 9090:9090 \
     -e DB_HOST=your_postgres_host \
     -e DB_PORT=5432 \
     -e DB_NAME=permissions \
     -e DB_USER=your_user \
     -e DB_PASSWORD=your_password \
     permissions:latest
   ```

## Development

### Generate gRPC Code

```bash
./gen-grpc.sh
```

### Running Tests

```bash
./test.sh
```

To view test coverage:

```bash
go tool cover -html=coverage.out
```

### Benchmarking

```bash
./benchmark.sh
```

## API Documentation

The service API is defined in `pkg/grpc/permissions.proto`. The service provides endpoints for:

* User permission management
* Role management
* Scope configuration
* Permission validation

Use tools like `buf` to generate client code for your preferred programming language.

## Configuration

The service can be configured through environment variables:

* `DB_HOST`: PostgreSQL host address
* `DB_PORT`: PostgreSQL port (default: 5432)
* `DB_NAME`: Database name
* `DB_USER`: Database user
* `DB_PASSWORD`: Database password
* `GRPC_PORT`: gRPC server port
* Additional configuration options in `.env`

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

Please ensure to:

* Add tests for new features
* Update documentation as needed
* Follow the existing coding style

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

## Support

For support, please open an issue in the GitHub repository.
