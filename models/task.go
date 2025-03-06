package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	User_ID     primitive.ObjectID `json:"user_id" bson:"user_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Due_date    time.Time          `json:"duedate"`
	Status      string             `json:"status"`
}