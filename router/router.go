package router

import (
	"TM/controllers"
	"TM/data"
	"TM/db"
	"TM/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRouter() *gin.Engine {

	client := db.DbInitalize()

	// create a new gin router
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // No content for OPTIONS requests
			return
		}
	
		c.Next()
	})

	taskService := data.NewTaskService(client,"taskmanagment","tasks")
	TaskController := controllers.NewTaskController(taskService)

	userService  := data.NewUserService(client,"taskmanagment","users")
	UserController := controllers.NewUserController(userService)

	router.POST("/Register",UserController.RegisterUserController)
	router.POST("/Login",UserController.LoginUserController)

	// Middleware to secure routes
	secureRoutes := router.Group("/").Use(middleware.AuthMiddleware())

	secureRoutes.GET("/tasks", TaskController.GetTaskController)
	secureRoutes.GET("/tasks/:id", TaskController.GetTaskByIDController)
	secureRoutes.PUT("/tasks/:id", TaskController.UpdateTaskByIDController)
	secureRoutes.DELETE("/tasks/:id", TaskController.DeleteTaskByIDController)
	secureRoutes.POST("/tasks", TaskController.CreateTaskController)

	return router
	
}