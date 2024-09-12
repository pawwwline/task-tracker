package app

import (
	"errors"
	"task-tracker/lib/e"
	"task-tracker/models"
	"time"
)

const IDNotFound = "id of task is not found"

func counter(fileTasks []models.Task) int {
	var id int
	for i, _ := range fileTasks {
		if i == len(fileTasks)-1 {
			id = fileTasks[i].Id
		}
	}
	return id
}

func (a *App) AddTask(task string) (int, error) {
	fileTasks, err := a.Storage.LoadInfo()
	if err != nil {
		return -1, err
	}
	id := counter(fileTasks)
	data := models.Task{
		Id:          id + 1,
		Description: task,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   "Not updated",
	}
	err = a.Storage.SaveInfo(data)
	if err != nil {
		return -1, err
	}
	return id + 1, nil
}

func (a *App) UpdateTask(id int, task string) error {
	fileTasks, err := a.Storage.LoadInfo()
	if err != nil {
		return err
	}
	for i, _ := range fileTasks {
		if fileTasks[i].Id == id {
			fileTasks[i].Description = task
			fileTasks[i].UpdatedAt = time.Now()
			break
		} else {
			err := e.WrapError("ID error", errors.New(IDNotFound))
			if err != nil {
				return err
			}
		}
	}
	err = a.Storage.UpdateInfo(fileTasks)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) DeleteTask(id int) error {
	fileTasks, err := a.Storage.LoadInfo()
	if err != nil {
		return err
	}
	for i, _ := range fileTasks {
		if fileTasks[i].Id == id {
			fileTasks = append(fileTasks[:i], fileTasks[i+1:]...)
			break
		} else {
			err := e.WrapError("ID error", errors.New(IDNotFound))
			if err != nil {
				return err
			}
		}
	}
	err = a.Storage.UpdateInfo(fileTasks)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) ListAllTasks() ([]models.Task, error) {
	fileTasks, err := a.Storage.LoadInfo()
	if err != nil {
		return nil, err
	}
	return fileTasks, nil
}

func (a *App) ListDoneTasks() ([]models.Task, error) {
	fileTasks, err := a.Storage.LoadInfo()
	if err != nil {
		return nil, err
	}
	var doneTasks []models.Task
	for i, _ := range fileTasks {
		if fileTasks[i].Status == "done" {
			doneTasks = append(doneTasks, fileTasks[i])
		}
	}
	return doneTasks, nil
}

func (a *App) ListToDoTasks() ([]models.Task, error) {
	fileTasks, err := a.Storage.LoadInfo()
	if err != nil {
		return nil, err
	}
	var toDoTasks []models.Task
	for i, _ := range fileTasks {
		if fileTasks[i].Status == "todo" {
			toDoTasks = append(toDoTasks, fileTasks[i])
		}
	}
	return toDoTasks, nil
}

func (a *App) ListProgressTasks() ([]models.Task, error) {
	fileTasks, err := a.Storage.LoadInfo()
	if err != nil {
		return nil, err
	}
	var ProgressTasks []models.Task
	for i, _ := range fileTasks {
		if fileTasks[i].Status == "in-progress" {
			ProgressTasks = append(ProgressTasks, fileTasks[i])
		}
	}
	return ProgressTasks, nil
}

func (a *App) MarkInProgress(id int) error {
	fileTasks, err := a.Storage.LoadInfo()
	if err != nil {
		return err
	}
	for i, _ := range fileTasks {
		if fileTasks[i].Id == id {
			fileTasks[i].Status = "in-progress"
			fileTasks[i].UpdatedAt = time.Now()
			break
		} else {
			err := e.WrapError("ID error", errors.New(IDNotFound))
			if err != nil {
				return err
			}
		}

	}
	err = a.Storage.UpdateInfo(fileTasks)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) MarkDone(id int) error {
	fileTasks, err := a.Storage.LoadInfo()
	if err != nil {
		return err
	}
	for i, _ := range fileTasks {
		if fileTasks[i].Id == id {
			fileTasks[i].Status = "done"
			fileTasks[i].UpdatedAt = time.Now()
			break
		} else {
			err := e.WrapError("ID error", errors.New(IDNotFound))
			if err != nil {
				return err
			}
		}
	}
	err = a.Storage.UpdateInfo(fileTasks)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) MarkToDo(id int) error {
	fileTasks, err := a.Storage.LoadInfo()
	if err != nil {
		return err
	}
	for i, _ := range fileTasks {
		if fileTasks[i].Id == id {
			fileTasks[i].Status = "todo"
			fileTasks[i].UpdatedAt = time.Now()
			break
		} else {
			err := e.WrapError("ID error", errors.New(IDNotFound))
			if err != nil {
				return err
			}
		}
	}
	err = a.Storage.UpdateInfo(fileTasks)
	if err != nil {
		return err
	}
	return nil
}
