package config

import (
    "fmt"
    "my-todo-app/models"
    "os"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres" // PostgreSQL dialect
)

var DB *gorm.DB

func SetupDatabase() *gorm.DB {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbName := os.Getenv("DB_NAME")
    dbPassword := os.Getenv("DB_PASSWORD")

    dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)
    
    db, err := gorm.Open("postgres", dsn)
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to the database: %s", err.Error()))
    }
    fmt.Println("Connected to the database successfully")

    err = db.AutoMigrate(&models.ToDoData{}).Error
    if err != nil {
        panic(fmt.Sprintf("Failed to auto migrate: %s", err.Error()))
    }
    fmt.Println("Database migrated successfully")

    return db
}
