package storage

import "task-tracker/models"

type Storage interface {
	SaveInfo(task models.Task)
	LoadInfo() []models.Task
	UpdateInfo([]models.Task)
}
