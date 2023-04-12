package routes

import (
	"go-todo/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	todoRepo := controllers.New()
	r.GET("/todos", todoRepo.GetTodos)
	r.POST("/todo", todoRepo.CreateTodo)

	return r
}
