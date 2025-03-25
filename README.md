# SecureAPIWithgrpc

SecureAPIWithgrpc is a Go-based project that implements a secure API using gRPC. It includes authentication, authorization, and CRUD operations while following an MVC structure.

## Features

- Secure authentication and authorization using JWT
- CRUD operations for employee management
- Uses gRPC for efficient communication
- MySQL database integration
- Structured project layout following best practices

## Project Structure

```
SecureAPIWithgrpc/
├── authorization/        # Handles authentication & authorization
│   ├── auth.go          # Authentication logic
├── config/              # Configuration files
│   ├── config.go        # Application configuration
├── database/            # Database connection setup
│   ├── database.go      # MySQL connection handling
├── grpcAPI/             # gRPC-related files
│   ├── proto/           # Protocol buffer definitions
│   │   ├── employee.proto # gRPC service definitions
│   ├── protobuf/        # Generated protobuf files
│   ├── grpcHandlers.go  # gRPC request handlers
├── handlers/            # Business logic and gRPC handlers
│   ├── employeeHandler.go # Employee service handlers
├── model/               # Data models
│   ├── employee.go      # Employee model definition
├── main.go              # Entry point of the application
├── go.mod               # Go module file
├── go.sum               # Dependency tracking file
├── README.md            # Project documentation
```

## Installation & Setup

### Prerequisites

- Go 1.20+
- MySQL
- Protobuf compiler

### Clone the repository

```sh
git clone https://github.com/yourusername/SecureAPIWithgrpc.git
cd SecureAPIWithgrpc
```

### Install dependencies

```sh
go mod tidy
```

### Generate gRPC code

```sh
protoc --proto_path=proto --go_out=protobuf --go_opt=paths=source_relative --go-grpc_out=protobuf --go-grpc_opt=paths=source_relative proto/employee.proto
```

### Configure MySQL Database

Update `database/database.go` with your MySQL credentials.

### Run the application

```sh
go run main.go
```

## API Endpoints

### gRPC Services

- `CreateEmployee`
- `GetEmployeeByID`
- `UpdateEmployee`
- `DeleteEmployee`

## Security

- JWT authentication for API security
- Environment variable-based configuration management



