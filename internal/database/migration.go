package database

import (
	"context"
	"log"
)

func Migrate() {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			password TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`

	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to run migration: %v", err)
	}

	log.Println("Database migration completed ✅")
}
