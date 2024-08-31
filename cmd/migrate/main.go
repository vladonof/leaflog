package main

import (
	"log"

	"github.com/vladonof/leaflog/internal/database"
	"github.com/vladonof/leaflog/internal/storage"
	"github.com/vladonof/leaflog/internal/tasting"
	"github.com/vladonof/leaflog/internal/tea"
	"github.com/vladonof/leaflog/internal/user"
)

func main() {
	database.ConnectDatabase()

	err := database.DB.AutoMigrate(
		&user.User{},
		&storage.Storage{},
		&tea.Tea{},
		&tasting.TastingSession{},
		&tasting.TastingNote{},
	)
	if err != nil {
		log.Fatal("failed to run migration: ", err)
	}

	log.Println("Database migration completed successfully")
}
