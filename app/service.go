package app

import (
	"task-tracker/models"
	"time"
)

//var counter int

func counter(fileTasks []models.Task) int {
	var id int
	for i, _ := range fileTasks {
		if i == len(fileTasks)-1 {
			id = fileTasks[i].Id
		}
	}
	return id
}

func (a *App) AddTask(task string) int {
	fileTasks := a.Storage.LoadInfo()
	id := counter(fileTasks)
	data := models.Task{
		Id:          id + 1,
		Description: task,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   "Not updated",
	}
	a.Storage.SaveInfo(data)
	return id + 1
}

func (a *App) UpdateTask(id int, task string) {
	fileTasks := a.Storage.LoadInfo()
	for i, _ := range fileTasks {
		if fileTasks[i].Id == id {
			fileTasks[i].Description = task
			fileTasks[i].UpdatedAt = time.Now()
			break
		}
	}
	a.Storage.UpdateInfo(fileTasks)
}

func (a *App) DeleteTask(id int) {
	fileTasks := a.Storage.LoadInfo()
	for i, _ := range fileTasks {
		if fileTasks[i].Id == id {
			fileTasks = append(fileTasks[:i], fileTasks[i+1:]...)
			break
		}
	}
	a.Storage.UpdateInfo(fileTasks)
}

func (a *App) ListAllTasks() []models.Task {
	fileTasks := a.Storage.LoadInfo()
	return fileTasks
}

func (a *App) ListDoneTasks() []models.Task {
	fileTasks := a.Storage.LoadInfo()
	var doneTasks []models.Task
	for i, _ := range fileTasks {
		if fileTasks[i].Status == "done" {
			doneTasks = append(doneTasks, fileTasks[i])
		}
	}
	return doneTasks
}

func (a *App) ListToDoTasks() []models.Task {
	fileTasks := a.Storage.LoadInfo()
	var toDoTasks []models.Task
	for i, _ := range fileTasks {
		if fileTasks[i].Status == "todo" {
			toDoTasks = append(toDoTasks, fileTasks[i])
		}
	}
	return toDoTasks
}

func (a *App) ListProgressTasks() []models.Task {
	fileTasks := a.Storage.LoadInfo()
	var ProgressTasks []models.Task
	for i, _ := range fileTasks {
		if fileTasks[i].Status == "in-progress" {
			ProgressTasks = append(ProgressTasks, fileTasks[i])
		}
	}
	return ProgressTasks
}

func (a *App) MarkInProgress(id int) {
	fileTasks := a.Storage.LoadInfo()
	for i, _ := range fileTasks {
		if fileTasks[i].Id == id {
			fileTasks[i].Status = "in-progress"
			fileTasks[i].UpdatedAt = time.Now()
			break
		}
	}
	a.Storage.UpdateInfo(fileTasks)
}

func (a *App) MarkDone(id int) {
	fileTasks := a.Storage.LoadInfo()
	for i, _ := range fileTasks {
		if fileTasks[i].Id == id {
			fileTasks[i].Status = "done"
			fileTasks[i].UpdatedAt = time.Now()
			break
		}
	}
	a.Storage.UpdateInfo(fileTasks)
}

func (a *App) MarkToDo(id int) {
	fileTasks := a.Storage.LoadInfo()
	for i, _ := range fileTasks {
		if fileTasks[i].Id == id {
			fileTasks[i].Status = "todo"
			fileTasks[i].UpdatedAt = time.Now()
			break
		}
	}
	a.Storage.UpdateInfo(fileTasks)
}
