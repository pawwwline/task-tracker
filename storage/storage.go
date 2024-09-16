package storage

import "task-tracker/models"

type Storage interface {
	AddTask(task string) error
	DeleteTask(id string) error
	GetAll() ([]models.Task, error)
	GetByStatus(status string) ([]models.Task, error)
	MarkTask(id string) error
}
