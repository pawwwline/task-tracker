package models

import "time"

type TaskStatus string

const (
	STATUS_IN_PRROGRESS TaskStatus = "in-progress"
	STATUS_TODO         TaskStatus = "to-do"
	STATUS_DONE         TaskStatus = "done"
)

type Task struct {
	Id          int         `json:"id"`
	Description string      `json:"description"`
	Status      TaskStatus  `json:"status"`
	CreatedAt   time.Time   `json:"created"`
	UpdatedAt   interface{} `json:"updated"`
}
