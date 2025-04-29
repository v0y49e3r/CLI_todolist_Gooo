package main

import (
	"encoding/json"
	"os"
	"time"
)

// Todo represents an individual to-do item
type Todo_json struct {
	Title       string     `json:"title"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type Todos_json []Todo_json

type Storage struct {
	FileName string
}

func NewStorage(fileName string) *Storage {
	return &Storage{FileName: fileName}
}

func (s *Storage) Save(todos Todos) error {
	fileData, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storage) Load(todos *Todos) error {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileData, todos)
}
