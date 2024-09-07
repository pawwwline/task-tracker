package files

import (
	"log"
	"os"
)

type FileStorage struct {
	Filename string
}

func NewFileStorage(filename string) *FileStorage {
	return &FileStorage{Filename: filename}
}

const (
	perm = 0754
)

func (fs *FileStorage) SaveInfo(data []byte) {
	f, err := os.OpenFile(fs.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perm)
	if err != nil {
		log.Fatalf("open file failed: %v", err)
	}
	n, err := f.Write(data)
	log.Printf("writed %d bytes", n)

	if err != nil {
		log.Fatalf("write file failed: %v", err)
	}
}

func (fs *FileStorage) LoadInfo() []byte {
	data, err := os.ReadFile(fs.Filename)
	if err != nil {
		log.Fatalf("read file failed: %v", err)
	}
	return data
}
