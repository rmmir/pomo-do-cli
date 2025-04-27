package database

import (
	"log"
	"os"

	m "github.com/rmmir/pomo-do/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
const dbFile = "pomo-do.db"

func ConnectDB() {
	var err error
	if _, err = os.Stat(dbFile); err != nil {
		os.Create(dbFile)
	}

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
