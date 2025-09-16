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
	//TODO parse request
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//TODO check if user exists
	var existingUser models.User
	initializers.DB.Where("email = ?", user.Email).First(&existingUser)
	if existingUser.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong user or password"})
		return
	}
	//TODO compare password
	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)
	if !errHash {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong user or password"})
	}
	//TODO Generate JWT
	//
	//TODO Return token

	c.JSON(http.StatusOK, gin.H{"success": "user logged in"})
}

func LogOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "LogOut Called"})
}
