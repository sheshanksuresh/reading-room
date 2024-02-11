package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ConnectDB() *pgxpool.Pool {
	// Change this to get from environment variables instead probably
	dbURL := "postgres://sheshankpersonal@localhost:5432/reading_room"

	pool, err := pgxpool.Connect(context.Background(), dbURL)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	return pool
}
