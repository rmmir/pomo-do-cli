package database

import (
	"log"

	m "github.com/rmmir/pomo-do/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("pomo-do.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = DB.AutoMigrate(&m.Task{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	err = DB.AutoMigrate(&m.Category{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
