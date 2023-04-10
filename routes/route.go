package routes

import (
	"go-todo/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", controllers.GetTodos)

	return router
}
