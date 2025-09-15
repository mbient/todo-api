package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mbient/todo-api/models"
)

var tasks = []models.Task{
	{ID: "1", Title: "Buy groceries", Description: "Buy milk, eggs, and bread"},
	{ID: "2", Title: "Clean the house", Description: "Vacuum the living room and dust the shelves"},
	{ID: "3", Title: "Finish project report", Description: "Complete the final draft of the project report and submit it"},
	{ID: "4", Title: "Schedule dentist appointment", Description: "Call the dentist's office to schedule a check-up"},
	{ID: "5", Title: "Exercise", Description: "Go for a 30-minute run in the park"},
	{ID: "6", Title: "Read a book", Description: "Read at least two chapters of the current book"},
	{ID: "7", Title: "Prepare dinner", Description: "Cook spaghetti with marinara sauce and a side salad"},
	{ID: "8", Title: "Water the plants", Description: "Water all indoor and outdoor plants"},
	{ID: "9", Title: "Organize files", Description: "Sort and organize digital files on the computer"},
	{ID: "10", Title: "Call Mom", Description: "Check in with Mom and see how she's doing"},
}

type CreateTaskInput struct {
	ID          string `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

func GetTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	for _, task := range tasks {
		if task.ID == id {
			c.JSON(http.StatusOK, gin.H{"data": task})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func AddTask(c *gin.Context) {
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := models.Task{
		ID:          input.ID,
		Title:       input.Title,
		Description: input.Description,
	}
	tasks = append(tasks, task)
	c.JSON(http.StatusCreated, gin.H{"data": task})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	//var updatedTask models.Task
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := models.Task{
		ID:          input.ID,
		Title:       input.Title,
		Description: input.Description,
	}

	for i, item := range tasks {
		if item.ID == id {
			tasks[i] = task
			c.JSON(http.StatusOK, gin.H{"data": task})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	for i, item := range tasks {
		if item.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "task deleted!"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
