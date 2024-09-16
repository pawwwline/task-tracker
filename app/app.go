package app

import (
	"task-tracker/models"
	"task-tracker/storage"
)

type TaskService interface {
	AddTask(task string) (int, error)
	UpdateTask(id int, task string) error
	DeleteTask(id int) error
	ListAllTasks() ([]models.Task, error)
	ListByStatus(status models.TaskStatus) ([]models.Task, error)
	MarkTask(id int, status models.TaskStatus) error
}

type App struct {
	Storage storage.Repo
}

func NewApp(storage storage.Repo) *App {
	return &App{Storage: storage}
}
