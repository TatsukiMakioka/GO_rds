package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"my-todo-app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // SQLite dialect
)

var DB *gorm.DB

func SetupDatabase() *gorm.DB {
	tmpDir := "/tmp"
	originalDbPath := GetDatabasePath()
	tmpDbPath := fmt.Sprintf("%s/%s", tmpDir, "test.db")

	// コピー
	err := copyFile(originalDbPath, tmpDbPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to copy database: %s", err.Error()))
	}

	db, err := gorm.Open("sqlite3", tmpDbPath)
	if err != nil {
		panic("Failed to connect to the database")
	}

	// Drop the existing table if it exists and re-migrate
	db.DropTableIfExists(&models.ToDoData{})
	db.AutoMigrate(&models.ToDoData{})

	return db
}

func GetDatabasePath() string {
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "test.db"
	}
	return dbPath
}

func copyFile(src, dst string) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dst, data, 0644)
	return err
}
