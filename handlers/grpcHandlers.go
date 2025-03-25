package handlers

import (
	"SecureAPIWithgrpc/authorization"
	"SecureAPIWithgrpc/config"
	"SecureAPIWithgrpc/model"
	"context"
	"log"
	"net"
	"os"
	"os/exec"

	pb "SecureAPIWithgrpc/grpcAPI/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type EmployeeServiceServer struct {
	pb.UnimplementedEmployeeServiceServer
	db *gorm.DB
}

// Middleware to authenticate JWT
func authenticate(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "Missing metadata")
	}
	tokens := md["authorization"]
	if len(tokens) == 0 {
		return status.Errorf(codes.Unauthenticated, "Missing token")
	}

	_, err := authorization.ValidateToken(tokens[0])
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "Invalid token: %v", err)
	}
	return nil
}

func (s *EmployeeServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := authorization.GenerateToken(req.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate token: %v", err)
	}
	return &pb.LoginResponse{Token: token}, nil
}

func (s *EmployeeServiceServer) CreateEmployee(ctx context.Context, req *pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error) {
	if err := authenticate(ctx); err != nil {
		return nil, err
	}

	emp := model.Employee{
		Name:     req.Name,
		Age:      int(req.Age),
		Position: req.Position,
	}

	if err := config.DB.Create(&emp).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create employee: %v", err)
	}

	return &pb.CreateEmployeeResponse{Id: uint32(emp.ID)}, nil
}

func (s *EmployeeServiceServer) DeleteEmployee(ctx context.Context, req *pb.DeleteEmployeeRequest) (*pb.EmployeeResponse, error) {
	if err := authenticate(ctx); err != nil {
		return nil, err
	}

	var emp model.Employee
	if err := config.DB.First(&emp, req.Id).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "Employee not found")
	}

	if err := config.DB.Delete(&emp).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete employee")
	}

	return &pb.EmployeeResponse{Employee: &pb.Employee{
		Id:       uint32(emp.ID),
		Name:     emp.Name,
		Age:      int32(emp.Age),
		Position: emp.Position,
	}}, nil
}

func (s *EmployeeServiceServer) GetAllEmployees(ctx context.Context, req *pb.EmployeeRequest) (*pb.EmployeeList, error) {
	if err := authenticate(ctx); err != nil {
		return nil, err
	}

	var employees []model.Employee
	config.DB.Find(&employees)

	var empList []*pb.Employee
	for _, emp := range employees {
		empList = append(empList, &pb.Employee{
			Id:       uint32(emp.ID),
			Name:     emp.Name,
			Age:      int32(emp.Age),
			Position: emp.Position,
		})
	}

	return &pb.EmployeeList{Employees: empList}, nil
}

func (s *EmployeeServiceServer) GetEmployeeByID(ctx context.Context, req *pb.GetEmployeeByIDRequest) (*pb.GetEmployeeByIDResponse, error) {
	if err := authenticate(ctx); err != nil {
		return nil, err
	}

	var emp model.Employee
	if err := config.DB.First(&emp, req.Id).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "Employee not found: %v", err)
	}

	return &pb.GetEmployeeByIDResponse{
		Employee: &pb.Employee{
			Id:       uint32(emp.ID),
			Name:     emp.Name,
			Age:      int32(emp.Age),
			Position: emp.Position,
		},
	}, nil
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	empService := &EmployeeServiceServer{}
	pb.RegisterEmployeeServiceServer(grpcServer, empService)

	go func() {
		log.Println("Starting gRPC server on port 50051...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %s", err)
		}
	}()

	go func() {
		grpcuiCmd := exec.Command("grpcui", "-plaintext", "localhost:50051")
		grpcuiCmd.Stdout = os.Stdout
		grpcuiCmd.Stderr = os.Stderr
		if err := grpcuiCmd.Start(); err != nil {
			log.Fatalf("Failed to start grpcui: %v", err)
		}
		if err := grpcuiCmd.Wait(); err != nil {
			log.Printf("grpcui process exited: %v", err)
		}
	}()
}
