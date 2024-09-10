package files

import (
	"encoding/json"
	"log"
	"os"
	"task-tracker/models"
)

type FileStorage struct {
	Filename string
	taskDB   []models.Task `json:"tasks"`
}

func NewFileStorage(filename string) *FileStorage {
	return &FileStorage{Filename: filename, taskDB: make([]models.Task, 0)}
}

const (
	perm = 0754
)

func (fs *FileStorage) SaveInfo(data models.Task) {
	fileData, err := os.ReadFile(fs.Filename)
	if err != nil {
		_, err := os.OpenFile(fs.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perm)
		if err != nil {
			log.Fatalf("open file failed: %v", err)
		}
	}

	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &fs.taskDB)
		if err != nil {
			log.Fatalf("parse file failed: %v", err)
		}
	}

	fs.taskDB = append(fs.taskDB, data)

	jsonData, _ := json.MarshalIndent(fs.taskDB, "", "  ")

	err = os.WriteFile(fs.Filename, jsonData, perm)
	if err != nil {
		log.Fatalf("write file failed: %v", err)
	}
}

func (fs *FileStorage) LoadInfo() []models.Task {
	fileData, err := os.ReadFile(fs.Filename)
	if err != nil {
		log.Fatalf("read file failed: %v", err)
	}

	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &fs.taskDB)
		if err != nil {
			log.Fatalf("parse file failed: %v", err)
		}
	}

	return fs.taskDB
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
