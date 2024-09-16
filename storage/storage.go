package storage

import "task-tracker/models"

type Repo interface {
	AddTask(task string) (int, error)
	DeleteTask(id int) error
	UpdateTask(id int, task string) error
	GetAll() ([]models.Task, error)
	GetByStatus(status models.TaskStatus) ([]models.Task, error)
	MarkTask(id int, status models.TaskStatus) error
}
