package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a global variable that will hold the connection pool
var DB *gorm.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase() {
	// Get the database connection string from an environment variable
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatal("DATABASE_DSN environment variable not set")
	}

	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = db

	log.Println("Database connection established")
}
