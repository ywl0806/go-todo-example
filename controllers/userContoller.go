package controllers

import (
	"go-todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	db.AutoMigrate(&models.User{})
	return &UserRepo{Db: db}
}

func (repository *UserRepo) GetUsers(c *gin.Context) {
	var todos []models.Todo
	err := models.GetTodos(repository.Db, &todos)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	print(todos)
	c.JSON(http.StatusOK, todos)
}
