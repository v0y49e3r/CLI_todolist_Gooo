package main

import "fmt"

func main() {
	// todos := Todos{}
	// // todos.add("hello")
	// // todos.add("hi")
	// // fmt.Println(todos)
	// // todos.deleteElement(0)

	// // fmt.Println(todos)
	// // todos.toggle(0)
	// todos.printFunc()
	// storage := NewStorage("todos.json")
	// err := storage.Load(&todos)
	// // if err != nil {
	// // 	fmt.Println(err)
	// // }
	// storage.Save(todos)

	// todos := Todos{}
	// storage := NewStorage("todos.json")
	// storage.Load(&todos)
	// cmdFlags := newCmdFlags()
	// cmdFlags.excute(&todos)
	// storage.Save(todos)

	todos := Todos{}
	storage := NewStorage("todos.json")

	// Load dữ liệu từ file
	if !storage.Load(&todos) {
		fmt.Println("Không thể load dữ liệu từ file. Danh sách công việc sẽ bắt đầu trống.")
	}

	// Phân tích và thực hiện các lệnh từ terminal
	cmdFlags := newCmdFlags()
	cmdFlags.excute(&todos)

	// Lưu lại dữ liệu sau khi thay đổi
	if !storage.Save(todos) {
		fmt.Println("Không thể lưu dữ liệu vào file.")
	}
}
