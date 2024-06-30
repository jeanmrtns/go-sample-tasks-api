package config

import (
	"fmt"
	"jeanmrtns/sample-go-api/models"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetSQLite() (*gorm.DB, error) {
	dbPath := "./db/main.db"

	_, err := os.Open(dbPath)

	if err != nil {
		fmt.Println("Database not found. Creating...")
		os.MkdirAll("./db", os.ModePerm)
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error opening database: %v", err)
		return nil, err
	}

	db.AutoMigrate(&models.Task{})

	return db, nil
}
