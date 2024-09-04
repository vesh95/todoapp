package main

import (
	"github.com/gin-gonic/gin"
	"todo/handlers"
	"todo/storage"
)

func main() {
	todoHandler := &handlers.TodoHandler{
		Repository: storage.NewTodoListStorage(),
	}

	router := gin.Default()
	router.GET("/todo", todoHandler.Index)
	router.POST("/todo", todoHandler.Store)
	router.GET("/todo/:id", todoHandler.Show)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err.Error())
	}
}
