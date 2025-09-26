package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mbient/todo-api/internal/controllers"
	"github.com/mbient/todo-api/internal/middleware"
)

func TaskRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		// public routes
		v1.POST("/register", controllers.SignUp)
		v1.POST("/login", controllers.LogIn)
		v1.POST("/logout", controllers.LogOut)

		// protected routes with jwt-auth
		v1.Use(middleware.JWTAuthMiddleware())
		{
			v1.GET("/tasks", controllers.GetTask)
			v1.GET("/tasks/:id", controllers.GetTaskByID)
			v1.POST("/tasks", controllers.AddTask)
			v1.PUT("/tasks/:id", controllers.UpdateTask)
			v1.DELETE("/tasks/:id", controllers.DeleteTask)
			v1.POST("/protected", controllers.Protected)
		}
	}
	return router
}
