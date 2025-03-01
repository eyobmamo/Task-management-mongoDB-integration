package models

import "time"

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Due_date    time.Time `json:"duedate"`
	Status      string     `json:"status"`
}