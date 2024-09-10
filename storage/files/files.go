package files

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task-tracker/models"
)

type FileStorage struct {
	Filename string
	TaskDB   []models.Task
}

func NewFileStorage(filename string) *FileStorage {
	return &FileStorage{Filename: filename, TaskDB: make([]models.Task, 0)}
}

const (
	perm = 0754
)

func (fs *FileStorage) SaveInfo(data models.Task) {
	fileData, err := os.ReadFile(fs.Filename)
	if os.IsNotExist(err) {
		f, err := os.OpenFile(fs.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perm)
		if err != nil {
			log.Fatalf("open file failed: %v", err)
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
			log.Fatalf("parse file failed: %v", err)
		}
	}

	fs.TaskDB = append(fs.TaskDB, data)

	jsonData, _ := json.MarshalIndent(fs.TaskDB, "", "  ")

	err = os.WriteFile(fs.Filename, jsonData, perm)
	if err != nil {
		log.Fatalf("write file failed: %v", err)
	}
}

func (fs *FileStorage) LoadInfo() []models.Task {
	fileData, err := os.ReadFile(fs.Filename)
	if os.IsNotExist(err) {
		f, err := os.OpenFile(fs.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perm)
		if err != nil {
			log.Fatalf("open file failed: %v", err)
		}
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Println("Error closing file:", err)
			}
		}()
	}

	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &fs.TaskDB)
		if err != nil {
			log.Fatalf("parse file failed: %v", err)
		}
	}

	return fs.TaskDB
}

func (fs *FileStorage) UpdateInfo(data []models.Task) {
	_, err := os.ReadFile(fs.Filename)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("file %s not found", fs.Filename)
		}
	}
	jsonData, _ := json.MarshalIndent(data, "", "  ")
	err = os.WriteFile(fs.Filename, jsonData, perm)

}
