package bookapp

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type userServiceServer struct {
	db *pgxpool.Pool
	UnimplementedUserServiceServer
}

func NewUserServiceServer(db *pgxpool.Pool) *userServiceServer {
	return &userServiceServer{
		db: db,
	}
}

func (s *userServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*UserResponse, error) {
	if req.Username == "" || req.Email == "" || req.Password == "" || req.FirstName == "" || req.LastName == "" {
		return nil, fmt.Errorf("Missing user form field.")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password.")
		return nil, fmt.Errorf("Failed due to internal error.")
	}

	var userId int
	err = s.db.QueryRow(ctx, "INSERT INTO users (username, email, password_hash, first_name, last_name, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id", req.Username, req.Email, hashedPassword, req.FirstName, req.LastName, time.Now()).Scan(&userId)

	if err != nil {
		log.Printf("Failed to insert user into db.")
		return nil, fmt.Errorf("Failed to create user: %v", err)
	}

	log.Printf("Created user ID: %v", userId)

	return &UserResponse{Id: fmt.Sprintf("%d", userId), Username: req.Username, Email: req.Email, FirstName: req.FirstName, LastName: req.LastName}, nil
}

func (s *userServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*UserResponse, error) {
	// TODO: Implement logic
	return &UserResponse{Id: req.Id, Username: "testuser", Email: "test@example.com", FirstName: "Test", LastName: "User"}, nil
}
