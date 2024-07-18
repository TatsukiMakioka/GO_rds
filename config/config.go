package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"my-todo-app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // SQLite dialect
)

var DB *gorm.DB

func SetupDatabase() *gorm.DB {
	tmpDir := "/tmp"
	originalDbPath := GetDatabasePath()
	tmpDbPath := filepath.Join(tmpDir, "test.db")

	fmt.Printf("Original DB Path: %s\n", originalDbPath)
	fmt.Printf("Temporary DB Path: %s\n", tmpDbPath)

	// コピー
	err := copyFile(originalDbPath, tmpDbPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to copy database: %s", err.Error()))
	}
	fmt.Println("Database file copied successfully")

	db, err := gorm.Open("sqlite3", tmpDbPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %s", err.Error()))
	}
	fmt.Println("Connected to the database successfully")

	// Drop the existing table if it exists and re-migrate
	err = db.DropTableIfExists(&models.ToDoData{}).Error
	if err != nil {
		panic(fmt.Sprintf("Failed to drop table: %s", err.Error()))
	}

	err = db.AutoMigrate(&models.ToDoData{}).Error
	if err != nil {
		panic(fmt.Sprintf("Failed to auto migrate: %s", err.Error()))
	}
	fmt.Println("Database migrated successfully")

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
		fmt.Printf("Failed to read source file: %s\n", err.Error())
		return err
	}
	fmt.Println("Source file read successfully")

	err = ioutil.WriteFile(dst, data, 0644)
	if err != nil {
		fmt.Printf("Failed to write to destination file: %s\n", err.Error())
	}
	return err
}
