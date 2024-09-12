package storage

import "task-tracker/models"

type Storage interface {
	SaveInfo(task models.Task) error
	LoadInfo() ([]models.Task, error)
	UpdateInfo([]models.Task) error
}
