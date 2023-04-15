package controllers

import (
	"fmt"
	"go-todo/database"
	"go-todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoRepo struct {
	Db *gorm.DB
}

func New() *TodoRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Todo{})
	return &TodoRepo{Db: db}
}

func (repository *TodoRepo) GetTodos(c *gin.Context) {
	var todos []models.Todo
	err := models.GetTodos(repository.Db, &todos)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	print(todos)
	c.JSON(http.StatusOK, todos)
}

type CreateTodoDto struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (repository *TodoRepo) CreateTodo(c *gin.Context) {
	var dto CreateTodoDto
	err := c.BindJSON(&dto)
	fmt.Println(err)
	fmt.Println(dto)
	// var todo models.Todo
	// models.CreateTodo(repository.Db, &todo)
	c.JSON(http.StatusOK, gin.H{"message": "OK!"})
}
