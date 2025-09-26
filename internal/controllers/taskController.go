package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mbient/todo-api/internal/initializers"
	"github.com/mbient/todo-api/internal/models"
)

type CreateTaskInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UpdateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetTask(c *gin.Context) {
	var task []models.Task
	initializers.DB.Find(&task)
	c.JSON(http.StatusOK, gin.H{"data": task})
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := initializers.DB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": task})
}

func AddTask(c *gin.Context) {
	// validate Input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
	}
	initializers.DB.Create(&task)
	c.JSON(http.StatusCreated, gin.H{"data": task})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	//var updatedTask models.Task
	var task models.Task
	if err := initializers.DB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task not found"})
		return
	}
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	initializers.DB.Model(&task).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": task})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := initializers.DB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task not found"})
		return
	}
	initializers.DB.Delete(&task, id)
	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
