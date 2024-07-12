package main

import (
	"time"
	"github.com/google/uuid"
)

// ToDoData represents a ToDo item.
type ToDoData struct {
	ID          string     `gorm:"type:text;primaryKey" json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime:false"`
	FinishedAt  *time.Time `json:"finished_at,omitempty"`
}

// NewToDoData creates a new instance of ToDoData.
func NewToDoData(title, description string) ToDoData {
	return ToDoData{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		CreatedAt:   time.Now().UTC(),
	}
}
