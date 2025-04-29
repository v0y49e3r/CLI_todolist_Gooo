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
	storage := Todo_json{}
}
