package controllers

import (
	"go-todo/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {

	var todo models.Todo
	todo.Title = "title1"
	todo.Discription = "discription"
	todo.IsActive = true
	todo.Deadline = time.Now()

	c.JSON(http.StatusOK, gin.H{"message": "OK!"})
}
