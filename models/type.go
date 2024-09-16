package models

import "time"

type TaskStatus string

const (
	StatusInProgress TaskStatus = "in-progress"
	StatusTodo       TaskStatus = "todo"
	StatusDone       TaskStatus = "done"
)

type Task struct {
	Id          int         `json:"id"`
	Description string      `json:"description"`
	Status      TaskStatus  `json:"status"`
	CreatedAt   time.Time   `json:"created"`
	UpdatedAt   interface{} `json:"updated"`
}
