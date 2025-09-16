package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mbient/todo-api/initializers"
	"github.com/mbient/todo-api/models"
	"github.com/mbient/todo-api/utils"

	"net/http"
)

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `gorm:"unique" json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignUp(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var existingUser models.User
	initializers.DB.Where("email = ?", input.Email).First(&existingUser)
	if existingUser.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		return
	}
	var errHash error
	input.Password, errHash = utils.GenerateHashPassword(input.Password)
	if errHash != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate password hash"})
		return
	}
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	initializers.DB.Create(&user)

	// return token
	c.JSON(http.StatusOK, gin.H{"success": "new user created"})
}

func LogIn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login Called"})
}

func LogOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "LogOut Called"})
}
