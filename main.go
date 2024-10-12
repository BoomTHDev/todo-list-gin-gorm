package main

import (
	"gin-todo/api"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/todos", api.GetAllLists)
	r.POST("/todo", api.CreateTodoList)

	r.Run()
}