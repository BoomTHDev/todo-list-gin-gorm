package main

import (
	"gin-todo/api"
	"gin-todo/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	database.ConnectDatabase()

	r.GET("/todos", api.GetAllLists)
	r.POST("/todo", api.CreateTodoList)

	r.Run()
}