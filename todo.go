package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

// var Todos = make([]Todo,0)
type Todos []Todo
type Todo struct {
	Title     string
	Completed bool
	CreatedAt time.Time
	//vi time.Time khong in ra nil duoc nen se dung poniter
	CompletedAt *time.Time
}

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)

}

func (todos *Todos) checkIsValid(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// delete
func (todos *Todos) deleteElement(index int) error {
	err := todos.checkIsValid(index)
	//todos no dang la 1 slice
	fmt.Println(&todos)
	if err != nil {
		return err
	}
	//toi ca tao 1 cai slice Todos moi voi kieu gia tri la Todo
	newTodos := Todos{}
	for i := 0; i < len(*todos); i++ {
		if i != index {
			newTodos = append(newTodos, (*todos)[i])
		}
	}
	*todos = newTodos

	return nil
}
func (todos *Todos) toggle(index int) error {
	err := todos.checkIsValid(index)
	if err != nil {
		return err
	}
	isCompleted := (*todos)[index].Completed
	if isCompleted == false {
		completionTime := time.Now()
		(*todos)[index].CompletedAt = &completionTime
		// fmt.Println(&completionTime)
	}
	(*todos)[index].Completed = !isCompleted
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	err := (*todos).checkIsValid(index)
	if err != nil {
		return err
	}
	(*todos)[index].Title = title
	return nil
}

func (todos *Todos) printFunc() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "CreatedAt", "CompletedAt")

	for i := 0; i < len(*todos); i++ {
		completed := "❌"
		completedAt := ""
		if (*todos)[i].Completed {
			completed = "✅"
			if (*todos)[i].CompletedAt != nil {
				completedAt = (*todos)[i].CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(i), (*todos)[i].Title, completed, (*todos)[i].CreatedAt.Format(time.RFC1123), completedAt)

	}
	table.Render()
}
