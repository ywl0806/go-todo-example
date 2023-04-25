package routes

import (
	"go-todo/controllers"
	"go-todo/database"
	"go-todo/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	db := database.InitDb()
	db.AutoMigrate(&models.Todo{}, &models.User{})

	todoRepo := controllers.NewTodoRepo(db)
	r.GET("/todos", todoRepo.GetTodos)
	r.POST("/todo", todoRepo.CreateTodo)

	userRepo := controllers.NewUserRepo(db)
	r.GET("/users", userRepo.GetUsers)
	return r
}
