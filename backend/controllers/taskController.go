package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetTask Called"})
}

func GetTaskByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetTaskByID Called"})
}

func AddTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "AddTask Called"})
}

func UpdateTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdateTask Called"})
}

func DeleteTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteTask Called"})
}
