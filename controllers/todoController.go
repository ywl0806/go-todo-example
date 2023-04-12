package controllers

import (
	"go-todo/database"
	"go-todo/models"
	"net/http"
	"time"

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

func (repository *TodoRepo) CreateTodo(c *gin.Context) {

	var todo models.Todo
	todo.Title = "title1"
	todo.Discription = "discription"
	todo.IsActive = true
	todo.Deadline = time.Now()

	models.CreateTodo(repository.Db, &todo)
	c.JSON(http.StatusOK, gin.H{"message": "OK!"})
}
