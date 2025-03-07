package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Email string  `json:"email"`
	Password string  `json:"password"`
	Role     string  `json:"role"` //"User" or "Admin"
}
