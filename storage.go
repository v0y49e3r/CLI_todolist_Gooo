package main

import (
	"encoding/json"
	"os"
	"time"
)

type Todo_json struct {
	Title       string     `json:"title"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at","omitempty"`
}
type Todos_json []Todo_json

func SaveTodos(filename string, todos Todos) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func LoadTodos(filename string) (Todos, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var todos Todos
	err = json.Unmarshal(data, &todos)
	return todos, err
}
