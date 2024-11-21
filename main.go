package main

import (
	"fmt"
)

type Task struct {
	id       int
	content string
	status   bool
}

type List struct {
	title string
	tasks []Task
}

func main() {
	var ToDO List
	ToDO.title = "My To-Do List"

	fmt.Printf("Welcome To %s App\n", ToDO.title)
}
