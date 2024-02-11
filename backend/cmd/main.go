package main

import (
	"log"
	"net"

	"github.com/sheshanksuresh/reading-room/backend/db"
	"github.com/sheshanksuresh/reading-room/backend/server/bookapp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	dbPool := db.ConnectDB()
	defer dbPool.Close()

	userServiceServer := bookapp.NewUserServiceServer(dbPool)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	bookapp.RegisterUserServiceServer(s, userServiceServer)
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
