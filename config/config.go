package config

import (
	"fmt"
	"my-todo-app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // PostgreSQL Dialect
)

var DB *gorm.DB

func SetupDatabase() *gorm.DB {
	// PostgreSQL接続情報
	host := "database-2.cxy8s2oaayga.ap-northeast-1.rds.amazonaws.com"
	port := 5432
	user := "postgres"
	password := "DB14mky10"
	dbname := "todogodb"

	// PostgreSQL接続文字列
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %s", err.Error()))
	}
	fmt.Println("Connected to the PostgreSQL database successfully")

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
