package config
 
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // SQLite dialect
 
	"./models"
)
 
// SetupDatabase sets up the database connection and initializes the schema
func SetupDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./test.db")
	if err != nil {
		panic("Failed to connect to the database")
	}
 
	// Drop the existing table if it exists and re-migrate
	db.DropTableIfExists(&models.ToDoData{})
	db.AutoMigrate(&models.ToDoData{})
 
	return db
}
