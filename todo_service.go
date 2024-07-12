package main

import (
	"my-todo-app/models"
	"my-todo-app/repositories"
	"time"
)

// ToDoService provides ToDo-related services.
type ToDoService struct {
	Repository *repositories.ToDoRepository
}

// GetToDos retrieves all ToDo items.
func (service *ToDoService) GetToDos() ([]models.ToDoData, error) {
	return service.Repository.GetToDos()
}

// CreateToDo adds a new ToDo item and returns its ID.
func (service *ToDoService) CreateToDo(todo models.ToDoData) (string, error) {
	createdToDo, err := service.Repository.CreateToDo(todo)
	if err != nil {
		return "", err
	}
	return createdToDo.ID, nil
}

// DeleteToDo removes a ToDo item by ID.
func (service *ToDoService) DeleteToDo(id string) error {
	return service.Repository.DeleteToDo(id)
}

// UpdateToDo updates an existing ToDo item by ID.
func (service *ToDoService) UpdateToDo(id string, todo models.ToDoData) (models.ToDoData, error) {
	return service.Repository.UpdateToDo(id, todo)
}

// GetToDoByID retrieves a single ToDo item by ID.
func (service *ToDoService) GetToDoByID(id string) (models.ToDoData, error) {
	return service.Repository.GetToDoByID(id)
}

// FinishToDo marks a ToDo item as finished.
func (service *ToDoService) FinishToDo(id string, finishedAt *time.Time) (models.ToDoData, error) {
	return service.Repository.FinishToDo(id, finishedAt)
}

// SearchToDos searches for ToDo items by keyword.
func (service *ToDoService) SearchToDos(keyword string) ([]models.ToDoData, error) {
	return service.Repository.SearchToDos(keyword)
}
