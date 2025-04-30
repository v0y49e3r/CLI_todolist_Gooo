package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Comdflag struct {
	Add     string
	Dellete int
	Edit    string
	Toggle  int
	List    bool
}

func newCmdFlags() *Comdflag {
	commandFlag := Comdflag{}

	flag.StringVar(&commandFlag.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&commandFlag.Edit, "edit", "", "Edit a todo by index & specify a new title. id:new_title")
	flag.IntVar(&commandFlag.Dellete, "del", -1, "Specify todo by index to delete")
	flag.IntVar(&commandFlag.Toggle, "toggle", -1, "Specify todo by index to toggle complete true/false")
	flag.BoolVar(&commandFlag.List, "list", false, "List all todos")

	flag.Parse()
	return &commandFlag
}

func (cmdFlag *Comdflag) excute(todos *Todos) {
	if cmdFlag.List == true {
		todos.printFunc()
		return
	}
	if cmdFlag.Add != "" {
		todos.add(cmdFlag.Add)
		return
	}
	if cmdFlag.Toggle != -1 {
		todos.toggle(cmdFlag.Toggle)
		return
	}
	if cmdFlag.Dellete != -1 {
		todos.deleteElement(cmdFlag.Dellete)
		return
	}
	if cmdFlag.Edit != "" {
		handleEdit(cmdFlag.Edit, todos)
		return
	}

	fmt.Println("Invalid command.")
	fmt.Println("Usage:")
	fmt.Println("  --list                     List all todos")
	fmt.Println("  --add \"title\"             Add a new todo")
	fmt.Println("  --edit index:new_title     Edit a todo")
	fmt.Println("  --toggle index             Toggle completion")
	fmt.Println("  --del index                Delete a todo")

}

func handleEdit(editArg string, todos *Todos) {
	parts := strings.SplitN(editArg, ":", 2)
	if len(parts) != 2 {
		log.Fatal("Error: Invalid format for edit. Please use index:new_title")
	}
	index, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal("Error: Invalid index for edit.")
	}
	todos.edit(index, parts[1])
}
