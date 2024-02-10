package bookapp

import (
	"context"
	"log"

	"github.com/sheshanksuresh/reading-room/server/bookapp"
)

type userServiceServer struct {
	bookapp.UnimplementedUserServiceServer
}

func NewUserServiceServer() *userServiceServer {
	return &userServiceServer{}
}

func (s *userServiceServer) CreateUser(ctx context.Context, req *bookapp.CreateUserRequest) (*bookapp.UserResponse, error) {
	// Implement data entry validation and processing logic
	log.Printf("Received: %v", req)
	return &bookapp.UserResponse{Id: "123", Username: req.Username, Email: req.Email, FirstName: req.FirstName, LastName: req.LastName}, nil
}

func (s *userServiceServer) GetUser(ctx context.Context, req *bookapp.GetUserRequest) (*bookapp.UserResponse, error) {
	// TODO: Implement logic
	return &bookapp.UserResponse{Id: req.Id, Username: "testuser", Email: "test@example.com", FirstName: "Test", LastName: "User"}, nil
}
