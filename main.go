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
	r.GET("/todo", api.GetTodoList)
	r.GET("/list/:id", api.GetList)
	r.POST("/todo", api.CreateTodoList)
	r.DELETE("/todo/:id", api.DeleteTodo)

	r.Run()
}