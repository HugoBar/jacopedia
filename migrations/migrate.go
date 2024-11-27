package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get database credentials from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Create the connection string
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Create a migration instance
	m, err := migrate.New(
		"file://migrations", // Path to the migrations folder
		connStr,
	)
	if err != nil {
		log.Fatalf("Error creating migration instance: %v", err)
	}

	if len(os.Args) > 1 && os.Args[1] == "down" {
		// Rollback migration
		err = m.Down()
		if err != nil {
			if err == migrate.ErrNoChange {
				fmt.Println("No migrations to roll back.")
			} else {
				log.Fatalf("Error rolling back migrations: %v", err)
			}
		} else {
			fmt.Println("Migration rolled back successfully!")
		}
	} else {
		// Apply all pending migrations (default action)
		err = m.Up()
		if err != nil {
			if err == migrate.ErrNoChange {
				fmt.Println("No migrations to apply.")
			} else {
				log.Fatalf("Error running migrations: %v", err)
			}
		} else {
			fmt.Println("Migrations applied successfully!")
		}
	}
}
