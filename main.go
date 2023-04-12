package main

import (
	"go-todo/routes"
)

func main() {
	router := routes.SetupRouter()

	router.Run("localhost:8080")
}
