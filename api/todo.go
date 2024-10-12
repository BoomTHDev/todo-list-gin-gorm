package api

import (
	// "fmt"
	// "gin-todo/database"
	// "io"
	// "log"
	"gin-todo/database"
	"net/http"
	// "os"

	"github.com/gin-gonic/gin"
)

type CreateTodo struct {
	Username string `json:"username" binding:"required"`
	Title string `json:"title" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func GetAllLists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func CreateTodoList(c *gin.Context) {
	var input CreateTodo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todoList := database.Todo{Username: input.Username, Title: input.Title, Message: input.Message}
	database.DB.Create(&todoList)

	c.JSON(http.StatusOK, gin.H{"Result": todoList})
}