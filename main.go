package main

import (
	"go-todo/database"
	"go-todo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := routes.SetupRouter()
	database.InitDb()
	router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "OK!"}) })
	router.Run("localhost:8080")
}
