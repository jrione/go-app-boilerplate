# Go App Boilerplate

This is a boilerplate for Go application with REST API (Gin), gRPC, and PostgreSQL (GORM).

## Features

- REST API with Gin framework
- gRPC server with protobuf
- Database integration with GORM and PostgreSQL
- Configuration with Viper
- Logging with Logrus
- Clean architecture with routes, controllers, repositories, helpers, plugins

## Structure

- `main.go`: Entry point
- `routes/`: API routes
- `controller/`: API handlers
- `helper/`: Utility functions
- `repository/`: Data access layer
- `plugin/`: Plugins like config and logger
- `proto/`: gRPC proto files

## Dependencies

- Gin: Web framework
- gRPC: RPC framework
- Logrus: Logging
- Viper: Configuration
- GORM: ORM for database
- PostgreSQL driver: For database connection

## Database Setup

Ensure PostgreSQL is running and update `config.yaml` with your database credentials:

```yaml
database:
  host: "localhost"
  port: "5432"
  user: "postgres"
  password: "password"
  dbname: "go_app_db"
  sslmode: "disable"
```

The application will auto-migrate the database schema.

## Running

1. `go mod tidy`
2. `go run main.go`

REST API on :8080
gRPC on :50051

## Config

Edit `config.yaml` for ports, etc.

## Example Endpoints

- GET /api/health
- GET /api/users/:id
- POST /api/users

## gRPC

gRPC server runs on :50051 with ExampleService.SayHello method.

Use tools like grpcurl or grpcui to test:

```bash
grpcurl -plaintext localhost:50051 list
grpcurl -plaintext -d '{"name": "World"}' localhost:50051 example.ExampleService/SayHello
```