package main

import (
	"my-todo-app/models"
	"time"
	"github.com/jinzhu/gorm"
)

// ToDoRepository provides access to the ToDo data storage.
type ToDoRepository struct {
	DB *gorm.DB
}

// GetToDos retrieves all ToDo items.
func (repo *ToDoRepository) GetToDos() ([]models.ToDoData, error) {
	var todos []models.ToDoData
	if err := repo.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// CreateToDo adds a new ToDo item.
func (repo *ToDoRepository) CreateToDo(todo models.ToDoData) (models.ToDoData, error) {
	if err := repo.DB.Create(&todo).Error; err != nil {
		return models.ToDoData{}, err
	}
	return todo, nil
}

// DeleteToDo removes a ToDo item by ID.
func (repo *ToDoRepository) DeleteToDo(id string) error {
	var todo models.ToDoData
	if err := repo.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		return err
	}
	return repo.DB.Delete(&todo).Error
}

// UpdateToDo updates an existing ToDo item by ID.
func (repo *ToDoRepository) UpdateToDo(id string, input models.ToDoData) (models.ToDoData, error) {
	var todo models.ToDoData
	if err := repo.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		return models.ToDoData{}, err
	}
	now := time.Now().UTC()
	input.UpdatedAt = &now
	if err := repo.DB.Model(&todo).Updates(input).Error; err != nil {
		return models.ToDoData{}, err
	}
	return todo, nil
}

// GetToDoByID retrieves a single ToDo item by ID.
func (repo *ToDoRepository) GetToDoByID(id string) (models.ToDoData, error) {
	var todo models.ToDoData
	if err := repo.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		return models.ToDoData{}, err
	}
	return todo, nil
}

// FinishToDo marks a ToDo item as finished.
func (repo *ToDoRepository) FinishToDo(id string, finishedAt *time.Time) (models.ToDoData, error) {
	var todo models.ToDoData
	if err := repo.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		return models.ToDoData{}, err
	}
	todo.FinishedAt = finishedAt
	if err := repo.DB.Model(&todo).Omit("updated_at").Update("finished_at", finishedAt).Error; err != nil {
		return models.ToDoData{}, err
	}
	return todo, nil
}

// SearchToDos searches for ToDo items by keyword.
func (repo *ToDoRepository) SearchToDos(keyword string) ([]models.ToDoData, error) {
	var todos []models.ToDoData
	if err := repo.DB.Where("title LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
