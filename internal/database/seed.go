package database

import (
	"context"
	"log"
)

func SeedUsers() {
	users := []struct {
		Name     string
		Email    string
		Password string
	}{
		{"Rabgay", "rabgay@example.com", "password123"},
		{"ByteEater", "byte@example.com", "secret456"},
	}

	for _, user := range users {
		_, err := DB.Exec(context.Background(), `
			INSERT INTO users (name, email, password)
			VALUES ($1, $2, $3)
			ON CONFLICT (email) DO NOTHING
		`, user.Name, user.Email, user.Password)

		if err != nil {
			log.Printf("Error inserting user %s: %v", user.Email, err)
		} else {
			log.Printf("Seeded user: %s", user.Email)
		}
	}
}
