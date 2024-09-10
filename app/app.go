package app

import (
	"task-tracker/models"
	"task-tracker/storage"
)

type TaskService interface {
	AddTask(task string) int
	UpdateTask(id int, task string) error
	DeleteTask(id int) error
	ListAllTasks() []models.Task
	ListDoneTasks() []models.Task
	ListProgressTasks() []models.Task
	ListToDoTasks() []models.Task
	MarkInProgress(id int) error
	MarkDone(id int) error
	MarkToDo(id int) error
}

type App struct {
	Storage storage.Storage
}

func NewApp(storage storage.Storage) *App {
	return &App{Storage: storage}
}
