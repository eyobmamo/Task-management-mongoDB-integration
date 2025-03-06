package controllers

import (
	"TM/data"
	"TM/models"
	"net/http"
	// "log"

	
	
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	TaskService  data.TaskManager
}

func NewTaskController(TaskService data.TaskManager) *TaskController {
	return &TaskController{
		TaskService:  TaskService,
	}
}

func (tc *TaskController) GetTaskController(c *gin.Context) {
	Tasks,err := tc.TaskService.GetTask()
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
	}
	if len(Tasks) == 0{
		c.JSON(http.StatusNotFound,gin.H{"massage":"no task recorded yet"})
		return
	}
	c.IndentedJSON(http.StatusOK,Tasks)
}

func (tc *TaskController) GetTaskByIDController( c *gin.Context){
	taskID := c.Param("id")
	// // id,err := strconv.Atoi(taskID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest,"can not convet given ID to int")
	// 	return
	// }

	task,err := tc.TaskService.GetTaskByID(taskID)
	if err != nil {
		c.JSON(http.StatusNotFound,err)
		return
	}
	c.JSON(http.StatusOK,task)
}

func (tc *TaskController) UpdateTaskByIDController(c *gin.Context) {
	taskID := c.Param("id")
	// id, err := strconv.Atoi(taskID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
	// 	return
	// }
	var updateTask models.Task
	if err := c.ShouldBindJSON(&updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	error := tc.TaskService.UpdateTaskByID(taskID, updateTask)

	if error != nil {
		c.JSON(http.StatusBadRequest, error)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task updated successfully"})
}

func (tc *TaskController) DeleteTaskByIDController(c *gin.Context) {
	taskid := c.Param("id")


	
	error := tc.TaskService.DeleteTaskByID(taskid)
	if error != nil {
		// Log the error for debugging
		// log.Printf("Error deleting task with ID %s: %v", taskid, error)
		c.JSON(http.StatusNotFound, gin.H{"message": "not task found with a given id"})
		return
	}

	
	c.JSON(http.StatusOK,gin.H{"message":"Task deleted successfully"})
}

func (tc *TaskController) CreateTaskController(c *gin.Context)  {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return 
	}

	userID,exists :=c.Get("userID")

	if !exists{
		c.JSON(http.StatusUnauthorized,gin.H{"message":"User ID not found in token"})
	}
	newTask.ID= primitive.NewObjectID()
	objectID,error := primitive.ObjectIDFromHex(userID.(string))
	if error != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message":"invalid id payload in token"})
		return
	}
	newTask.User_ID = objectID

	// Log the new task for debugging
	// log.Printf("Creating new task: %+v", newTask)

	err := tc.TaskService.CreateTask(newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		
		return 
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task Created successfully"})

}

