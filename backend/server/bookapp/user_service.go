package bookapp

import (
	"context"
	"log"
)

type userServiceServer struct {
	UnimplementedUserServiceServer
}

func NewUserServiceServer() *userServiceServer {
	return &userServiceServer{}
}

func (s *userServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*UserResponse, error) {
	// Implement data entry validation and processing logic
	log.Printf("Received: %v", req)
	return &UserResponse{Id: "123", Username: req.Username, Email: req.Email, FirstName: req.FirstName, LastName: req.LastName}, nil
}

func (s *userServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*UserResponse, error) {
	// TODO: Implement logic
	return &UserResponse{Id: req.Id, Username: "testuser", Email: "test@example.com", FirstName: "Test", LastName: "User"}, nil
}
