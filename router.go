package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"my-todo-app/controllers"
	"my-todo-app/repositories"
	"my-todo-app/services"
)

// SetupRouter sets up the Gin router with routes and their handlers.
func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	todoRepo := repositories.ToDoRepository{DB: db}
	todoService := services.ToDoService{Repository: &todoRepo}
	todoController := controllers.ToDoController{Service: todoService}

	router.GET("/api", todoController.GetToDos)
	router.POST("/api", todoController.CreateToDo)
	router.DELETE("/api/:id", todoController.DeleteToDo)
	router.PUT("/api/:id", todoController.UpdateToDo)
	router.GET("/api/:id", todoController.GetToDoByID)
	router.POST("/api/:id/finish", todoController.FinishToDo)
	router.GET("/api/search", todoController.SearchToDos)

	return router
}
