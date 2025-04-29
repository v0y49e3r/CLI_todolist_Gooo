package main

import "fmt"

func main() {
	todos := Todos{}
	todos.add("hello")
	todos.add("hi")
	fmt.Println(todos)
	// todos.deleteElement(0)

	// fmt.Println(todos)
	todos.toggle(0)
	todos.printFunc()
	storage := NewStorage("todos.json")
	err := storage.Load(&todos)
	if err != nil {
		fmt.Println(err)
	}
	storage.Save(todos)
}
