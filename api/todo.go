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
	var todoLists []database.Todo
	database.DB.Find(&todoLists)
	c.JSON(http.StatusOK, gin.H{"Result": todoLists})
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

func GetTodoList(c *gin.Context) {
	var todoLists []database.Todo
	if err := database.DB.Where("username = ?", c.Query("username")).Find(&todoLists).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"todoList": todoLists})
}

func DeleteTodo(c *gin.Context) {
	var todoList database.Todo
	if err := database.DB.Where("id = ?", c.Param("id")).First(&todoList).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.DB.Delete(&todoList)
	c.JSON(http.StatusOK, gin.H{"message": "delete success"})
}

func GetList(c *gin.Context) {
	var todolist database.Todo
	if err := database.DB.Where("id = ?", c.Param("id")).First(&todolist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"list": todolist})
}