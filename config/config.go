package config

import (
	"my-todo-app/models"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // SQLite dialect
)

// SetupDatabase sets up the database connection and initializes the schema
func SetupDatabase() *gorm.DB {
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "./test.db"
	}

	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		panic("Failed to connect to the database")
	}

	// Drop the existing table if it exists and re-migrate
	db.DropTableIfExists(&models.ToDoData{})
	db.AutoMigrate(&models.ToDoData{})

	return db
}
