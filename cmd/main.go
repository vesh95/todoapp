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
	router.DELETE("/todo/:id", todoHandler.Delete)

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		panic(err.Error())
	}
}
