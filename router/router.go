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