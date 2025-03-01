package data

import (
	"TM/models"
	"context"
	// "fmt"
	// "time"

	// "log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type  TaskManager interface {
	GetTask() ([]models.Task,error)
	GetTaskByID(TaskID string) (models.Task,error)
	UpdateTaskByID(TaskID string,updateTask models.Task) error
	DeleteTaskByID(TaskID string) error
	CreateTask(newTask models.Task) error 
}

// type TaskService struct {
// 	Tasks map[int]models.Task
// }
type TaskService struct {
	collection *mongo.Collection
}




// func NewTaskService() *TaskService {
// 	ts := &TaskService{
// 		Tasks: make(map[int]models.Task),
// 	}

func NewTaskService(client *mongo.Client,dbName,collectionName string) *TaskService{
	return &TaskService{
		collection : client.Database(dbName).Collection(collectionName),
	}
} 



// func (tc *TaskService) GetTask () []models.Task {
// 	var tasks []models.Task 
// 	for _,task := range tc.Tasks{
// 		tasks= append(tasks, task)
// 	}
// 	return tasks
// }

func (ts *TaskService) GetTask() ([]models.Task,error){
	var tasks []models.Task
	cursor,err := ts.collection.Find(context.TODO(),bson.M{})
	if err != nil {
		return nil,err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()){
		var task models.Task
		if err := cursor.Decode(&task);err != nil {
			return nil,err
		}
		tasks = append(tasks,task)
	}
	return tasks,nil
}



// func (tc *TaskService) GetTaskByID(TaskID int) (*models.Task, error) {
// 	task, exist := tc.Tasks[TaskID]

// 	if !exist {
// 		return nil, fmt.Errorf("Task not exist with given ID")
// 	}
// 	return &task, nil
// }

func (ts *TaskService) GetTaskByID(taskID string) (models.Task,error){
	var task models.Task

	objectID,err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return task,err
	}

	filter := bson.M{"_id":objectID}

	error := ts.collection.FindOne(context.TODO(),filter).Decode(&task)
	return task,error
}


// func (tc *TaskService) UpdateTaskByID(TaskID int, UpdateTask models.Task) error {
// 	task, Exist := tc.Tasks[TaskID]
// 	if !Exist {
// 		return fmt.Errorf("NO task with given ID")
// 	}

// 	if UpdateTask.Title != "" {
// 		task.Title = UpdateTask.Title
// 	}

// 	if UpdateTask.Description != "" {
// 		task.Description = UpdateTask.Description
// 	}

// 	tc.Tasks[TaskID] = task // Update the task in the map

// 	return nil
// }

func (ts *TaskService) UpdateTaskByID(id string,updatedTask models.Task) error {
	objectID ,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id":objectID}
	update := bson.M{"$set": updatedTask}

	_,error := ts.collection.UpdateOne(context.TODO(),filter,update)
	return error
}
// func (tc *TaskService) DeleteTaskByID(TaskID int) error {
// 	_,exist := tc.Tasks[TaskID]
// 	if !exist {
// 		return fmt.Errorf("Task not exist with a given ID")
// 	}
// 	delete(tc.Tasks,TaskID)

// 	return nil
// }

// func (ts *TaskService) DeleteTaskByID(TaskID string) error {
// 	objectID, err := primitive.ObjectIDFromHex(TaskID)

// 	if err != nil {
// 		return err
// 	}
// 	filter := bson.M{"_id": objectID}

// 	_, err = ts.collection.DeleteOne(context.TODO(), filter)
// 	return err
// }

func (ts *TaskService) DeleteTaskByID(TaskID string) error {
	objectID,err := primitive.ObjectIDFromHex(TaskID)

	if err != nil {
		return err
	}
	filter := bson.M{"_id":objectID}

	_,err = ts.collection.DeleteOne(context.TODO(),filter)
	return err
}

// func (tc *TaskService) CreateTask(newTask *models.Task) error {
// 	if _, exists := tc.Tasks[newTask.ID]; exists {
// 		return fmt.Errorf("Task with ID %d already exists", newTask.ID)
// 	}
// 	tc.Tasks[newTask.ID] = *newTask
// 	return nil
// }

func (ts *TaskService) CreateTask (task models.Task) error {
	_,err := ts.collection.InsertOne(context.TODO(),task)
	return err
}