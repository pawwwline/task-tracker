package files

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"task-tracker/lib/e"
	"task-tracker/models"
	"time"
)

type FileStorage struct {
	filename string
	taskDB   map[int]models.Task
}

func NewFileStorage(filename string) *FileStorage {
	return &FileStorage{filename: filename, taskDB: make(map[int]models.Task)}
}

const (
	perm = 0754
)

func (fs *FileStorage) LoadFile() error {
	fileData, err := os.ReadFile(fs.filename)
	if os.IsNotExist(err) {
		f, err := os.OpenFile(fs.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perm)
		if err != nil {
			log.Fatalf("open file failed: %v", err)
			return err
		}
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatalf("close file failed: %v", err)
			}
		}()
	}
	fs.taskDB = make(map[int]models.Task)
	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &fs.taskDB)
		if err != nil {
			return e.WrapError("parse file failed", err)
		}
		fmt.Println(&fs.taskDB)
	}
	return nil
}

func (fs *FileStorage) SaveFile(data map[int]models.Task) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return e.WrapError("json marshal failed", err)
	}
	err = os.WriteFile(fs.filename, jsonData, perm)
	if err != nil {
		log.Fatalf("write file failed: %v", err)
		return err
	}
	return nil
}

func (fs *FileStorage) AddTask(task string) (int, error) {
	err := fs.LoadFile()
	if err != nil {
		return 0, err
	}
	data := models.Task{
		Id:          len(fs.taskDB) + 1,
		Description: task,
		Status:      models.StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   "Not updated",
	}
	fs.taskDB[data.Id] = data
	fmt.Println(fs.taskDB)
	if err := fs.SaveFile(fs.taskDB); err != nil {
		return -1, e.WrapError("save task failed", err)
	}
	return data.Id, nil
}

func (fs *FileStorage) DeleteTask(id int) error {
	err := fs.LoadFile()
	if err != nil {
		return err
	}
	if _, ok := fs.taskDB[id]; ok {
		delete(fs.taskDB, id)
		if err := fs.SaveFile(fs.taskDB); err != nil {
			return e.WrapError("delete task failed", err)
		}
	} else {
		return errors.New("ID not found")
	}
	return nil
}

func (fs *FileStorage) UpdateTask(id int, task string) error {
	err := fs.LoadFile()
	if err != nil {
		return err
	}
	if curTask, ok := fs.taskDB[id]; ok {
		curTask.Description = task
		curTask.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		fs.taskDB[id] = curTask
		if err := fs.SaveFile(fs.taskDB); err != nil {
			return e.WrapError("update task failed", err)
		}
	} else {
		return errors.New("ID not found")
	}
	return nil
}

func (fs *FileStorage) GetAll() ([]models.Task, error) {
	if err := fs.LoadFile(); err != nil {
		return nil, err
	}
	var tasks []models.Task
	for _, task := range fs.taskDB {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (fs *FileStorage) GetByStatus(status models.TaskStatus) ([]models.Task, error) {
	if err := fs.LoadFile(); err != nil {
		return nil, err
	}
	var tasks []models.Task
	for _, task := range fs.taskDB {
		if task.Status == status {
			tasks = append(tasks, task)
		}
	}
	return tasks, nil
}

func (fs *FileStorage) MarkTask(id int, status models.TaskStatus) error {
	err := fs.LoadFile()
	if err != nil {
		return err
	}
	if curTask, ok := fs.taskDB[id]; ok {
		curTask.Status = status
		curTask.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		fs.taskDB[id] = curTask
		if err := fs.SaveFile(fs.taskDB); err != nil {
			return e.WrapError("update task failed", err)
		}
	} else {
		return errors.New("ID not found")
	}
	return nil
}
