package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"todo/handlers"
	"todo/storage"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", 0)
	todoHandler := &handlers.TodoHandler{
		Repository: storage.NewRedisStorage(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0}, logger),
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
