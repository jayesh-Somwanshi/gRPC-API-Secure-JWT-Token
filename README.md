# SecureAPIWithgrpc

SecureAPIWithgrpc is a Go-based project that implements a secure API using gRPC. It includes authentication, authorization, and CRUD operations while following an MVC structure.

-----------------------------------------------------------------------------------------------------------------------------------------------------

## Features

- Secure authentication and authorization using JWT
- CRUD operations for employee management
- Uses gRPC for efficient communication
- MySQL database integration
- Structured project layout following best practices

-----------------------------------------------------------------------------------------------------------------------------------------------------

## Project Structure

```
SecureAPIWithgrpc/
├── authorization/                # Handles authentication & authorization
│   ├── auth.go                   # Authentication logic
├── config/                       # Configuration files
│   ├── database.go               # MySQL connection handling
├── grpcAPI/                      # gRPC-related files
│   ├── proto/                    # Protocol buffer definitions
│   │   ├── employee.proto        # gRPC service definitions
│   ├── protobuf/                 # Generated protobuf files
│   │   ├── employee.pb.go        # Generated Go file for protobuf messages
│   │   ├── employee_grpc.pb.go   # Generated Go file for gRPC services
├── handlers/                     # Business logic and gRPC handlers
│   ├── grpcHandlers.go           # gRPC request handlers
├── model/                        # Data models
│   ├── employee.go               # Employee model definition
├── main.go                       # Entry point of the application
├── go.mod                        # Go module file
├── go.sum                        # Dependency tracking file
├── README.md                     # Project documentation
```
-----------------------------------------------------------------------------------------------------------------------------------------------------

## Running the Project
### Option 1
1. Navigate to the `SecureAPIWithgrpc/grpcAPI` directory:
   ```bash
   cd SecureAPIWithgrpc/grpcAPI
   ```
2. Run the following command to start the application:
   ```bash
   go run main.go

-----------------------------------------------------------------------------------------------------------------------------------------------------

## Installation & Setup

### Prerequisites

- Go 1.20+
- MySQL
- Protobuf compiler

-----------------------------------------------------------------------------------------------------------------------------------------------------

### Install dependencies
The go mod tidy command is used in Go projects to clean up and optimize the go.mod and go.sum
```sh
go mod tidy
```
-----------------------------------------------------------------------------------------------------------------------------------------------------

#### Option 1: Using `protoc`(Create the Protobuf Folder)
1. Navigate to the `SecureAPIWithgrpc/grpcAPI` directory:
   ```bash
   cd SecureAPIWithgrpc/grpcAPI
   ```
2. Run the following command to manually generate proto files:
   ```bash
   for proto_file in proto/*.proto; do \
       protoc --proto_path=proto \
       --go_out=paths=source_relative:./protobuf \
       --go-grpc_out=paths=source_relative:./protobuf \
       $$proto_file; \
   done
   ```
-----------------------------------------------------------------------------------------------------------------------------------------------------

### Configure MySQL Database

Update `config/database.go` with your MySQL credentials.

-----------------------------------------------------------------------------------------------------------------------------------------------------

### Run the application

```sh  
go run main.go
```
-----------------------------------------------------------------------------------------------------------------------------------------------------

### gRPC Services

- `CreateEmployee`
- `GetAllEmployees`
- `GetEmployeeByID`
- `UpdateEmployee`
- `DeleteEmployee`

-----------------------------------------------------------------------------------------------------------------------------------------------------

## Security

- JWT authentication for API security
- Environment variable-based configuration management

-----------------------------------------------------------------------------------------------------------------------------------------------------
## API Endpoints (gRPC UI Testing)

### 1. Login (Generate Token)
- **Select:** `Login`
- **Enter Credentials:**
  - **Username** (string)
  - **Password** (string)


### 2. Create Employee
- **Select:** `CreateEmployee`
- **Headers:**
  - **Name:** `authorization` (Hardcoded in code)
  - **Value:** `<generated_token>`
- **Enter Employee Details:**
  - **Name** (string)
  - **Age** (int32)
  - **Position** (string)

### 3. Delete Employee
- **Select:** `DeleteEmployee`
- **Headers:**
  - **Name:** `authorization` (Hardcoded in code)
  - **Value:** `<generated_token>`
- **Enter Employee ID:**
  - **ID** (uint32)

### 4. Get All Employees
- **Select:** `GetAllEmployees`
- **Headers:**
  - **Name:** `authorization` (Hardcoded in code)
  - **Value:** `<generated_token>`

### 5. Get Employee By ID
- **Select:** `GetEmployeesByID`
- **Headers:**
  - **Name:** `authorization` (Hardcoded in code)
  - **Value:** `<generated_token>`
- **Enter Employee ID:**
  - **ID** (uint32)

## Security Notes
- Always include the `authorization` header with the JWT token for secured endpoints.
- Ensure that the token is valid and has not expired.
- Store tokens securely and do not expose them in client-side code.


# grpc-API-With-JWT-Token-in-Go-Lang
# grpc-API-With-JWT-Token-in-Go-Lang
