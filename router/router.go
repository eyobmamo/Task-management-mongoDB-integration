package router

import (
	"github.com/gin-gonic/gin"
	"TM/data"
	"TM/controllers"
	"TM/db"
	
)

func InitializeRouter() *gin.Engine {

	client := db.DbInitalize()

	// create a new gin router
	router := gin.Default()

	taskService := data.NewTaskService(client,"taskmanagment","tasks")
	TaskController := controllers.NewTaskController(taskService)

	router.GET("/tasks",TaskController.GetTaskController)

	router.GET("/tasks/:id",TaskController.GetTaskByIDController)

	router.PUT("/tasks/:id",TaskController.UpdateTaskByIDController)

	router.DELETE("/tasks/:id",TaskController.DeleteTaskByIDController)

	router.POST("/tasks",TaskController.CreateTaskController)

	return router
	
}