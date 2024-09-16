package files

import (
	"encoding/json"
	"log"
	"os"
	"task-tracker/lib/e"
	"task-tracker/models"
)

type FileStorage struct {
	Filename string
	TaskDB   map[int]models.Task
	Tasks    []models.Task
}

func NewFileStorage(filename string) *FileStorage {
	return &FileStorage{Filename: filename, TaskDB: make(map[int]models.Task)}
}

const (
	perm = 0754
)

func (fs *FileStorage) Save(data models.Task) error {
	fileData, err := os.ReadFile(fs.Filename)
	if os.IsNotExist(err) {
		f, err := os.OpenFile(fs.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perm)
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

	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &fs.TaskDB)
		if err != nil {
			return e.WrapError("parse file failed", err)
		}
	}

	fs.TaskDB = append(fs.TaskDB, data)

	jsonData, err := json.MarshalIndent(fs.TaskDB, "", "  ")
	if err != nil {
		return e.WrapError("json marshal failed", err)
	}

	err = os.WriteFile(fs.Filename, jsonData, perm)
	if err != nil {
		log.Fatalf("write file failed: %v", err)
		return err
	}
	return nil
}

func (fs *FileStorage) GetAll() ([]models.Task, error) {
	fileData, err := os.ReadFile(fs.Filename)
	if os.IsNotExist(err) {
		f, err := os.OpenFile(fs.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perm)
		if err != nil {
			log.Fatalf("open file failed: %v", err)
			return nil, err
		}
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatalf("close file failed: %v", err)
			}
		}()
	}

	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &fs.TaskDB)
		if err != nil {
			return nil, e.WrapError("parse file failed", err)
		}
	}

	return fs.TaskDB, nil
}
