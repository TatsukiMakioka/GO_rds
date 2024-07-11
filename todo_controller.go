package controllers

import (
	"net/http"
	"time"

	"my-todo-app/models"
	"my-todo-app/services"

	"github.com/gin-gonic/gin"
)

// ToDoController handles HTTP requests for ToDo operations.
type ToDoController struct {
	Service services.ToDoService
}

// GetToDos handles GET requests to fetch all ToDo items.
func (ctrl *ToDoController) GetToDos(c *gin.Context) {
	todos, err := ctrl.Service.GetToDos()
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, todos)
}

// CreateToDo handles POST requests to add a new ToDo item.
func (ctrl *ToDoController) CreateToDo(c *gin.Context) {
	var input struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.NewToDoData(input.Title, input.Description)
	createdToDoUUID, err := ctrl.Service.CreateToDo(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, createdToDoUUID)
}

// DeleteToDo handles DELETE requests to remove a ToDo item by ID.
func (ctrl *ToDoController) DeleteToDo(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.Service.DeleteToDo(id); err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.Status(http.StatusOK)
}

// UpdateToDo handles PUT requests to update a ToDo item by ID.
func (ctrl *ToDoController) UpdateToDo(c *gin.Context) {
	id := c.Param("id")
	var input models.ToDoData
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedToDo, err := ctrl.Service.UpdateToDo(id, input)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, updatedToDo)
}

// GetToDoByID handles GET requests to fetch a single ToDo item by ID.
func (ctrl *ToDoController) GetToDoByID(c *gin.Context) {
	id := c.Param("id")
	todo, err := ctrl.Service.GetToDoByID(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, todo)
}

// FinishToDo handles POST requests to mark a ToDo item as finished.
func (ctrl *ToDoController) FinishToDo(c *gin.Context) {
	id := c.Param("id")
	now := time.Now().UTC()
	finishedToDo, err := ctrl.Service.FinishToDo(id, &now)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, finishedToDo)
}

// SearchToDos handles GET requests to search for ToDo items by keyword.
func (ctrl *ToDoController) SearchToDos(c *gin.Context) {
	keyword := c.Query("keyword")
	todos, err := ctrl.Service.SearchToDos(keyword)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, todos)
}
