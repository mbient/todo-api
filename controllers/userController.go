package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mbient/todo-api/initializers"
	"github.com/mbient/todo-api/models"
	"github.com/mbient/todo-api/utils"

	"fmt"
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
	//parse request
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//check if user exists
	var existingUser models.User
	initializers.DB.Where("email = ?", user.Email).First(&existingUser)
	if existingUser.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong user or password"})
		return
	}
	//compare password
	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)
	if !errHash {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong user or password"})
	}
	//generate JWT
	tokenString, err := utils.CreateJWTToken(existingUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}
	//return token
	//c.JSON(http.StatusOK, gin.H{"token": tokenString})
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", tokenString, 600, "", "", false, true) // 1o min
	c.JSON(http.StatusOK, gin.H{"success": "user logged in"})
}

func LogOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "LogOut Called"})
}

func Protected(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "claims not found"})
		return
	}
	userClaims, ok := claims.(*models.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello, %s!", userClaims.Subject)})
}
