package main

import (
	"encoding/json"
	"fmt"
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

func (s *Storage) Save(todos Todos) bool {
	fileData, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		fmt.Println(err)
		return false
	}
	e := os.WriteFile(s.FileName, fileData, 0644)
	if e != nil {
		fmt.Println(e)
		return false
	}
	return true
}

func (s *Storage) Load(todos *Todos) bool {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		fmt.Println(err)
		return false
	}
	e := json.Unmarshal(fileData, todos)
	if e != nil {
		fmt.Println(e)
		return false
	}
	return true
}
