package app

import (
	"errors"
	"task-tracker/models"
)

func (a *App) AddTask(task string) (int, error) {
	return a.Storage.AddTask(task)
}

func (a *App) UpdateTask(id int, task string) error {
	return a.Storage.UpdateTask(id, task)
}

func (a *App) DeleteTask(id int) error {
	return a.Storage.DeleteTask(id)
}

func (a *App) ListAllTasks() ([]models.Task, error) {
	tasks, _ := a.Storage.GetAll()
	if len(tasks) == 0 {
		return nil, errors.New("task list is empty")
	}
	return tasks, nil
}

func (a *App) ListByStatus(status models.TaskStatus) ([]models.Task, error) {
	tasks, _ := a.Storage.GetByStatus(status)
	if len(tasks) == 0 {
		return nil, errors.New("task list is empty")
	}
	return tasks, nil
}

func (a *App) MarkTask(id int, status models.TaskStatus) error {
	return a.Storage.MarkTask(id, status)
}
