package main

import (
	"fmt"
)

type Task struct {
	id      int
	content string
	status  bool
}

type List struct {
	title string
	tasks []Task
}

var choices = []string{
	"show all tasks",
	"show only uncompleted tasks",
	"done a task",
	"add new task",
	"delete a task",
	"quit",
}

func main() {
	var ToDO List
	ToDO.title = "My To-Do List"
	ToDO.tasks = []Task{
		{1, "the first task", false},
		{2, "the second task", true},
		{3, "the third task", false},
		{4, "the fourh task", true},
	}

	fmt.Printf("Welcome To %s App\n", ToDO.title)
main_loop:
	for {
		for idx, val := range choices {
			fmt.Printf("%d- %s\n", idx+1, val)
		}
		var choice int
		print("choose an option: ")
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			showTasks(ToDO, false)
		case 2:
			showTasks(ToDO, true)
		case 3:
		case 4:
		case 5:
			continue
		case 6:
			println("END..")
			break main_loop
		default:
			println("your choice is wrong, try again")
		}
	}
}

func showTasks(list List, uncompleted bool) {
	for idx, val := range list.tasks {
		status := "NO"
		if val.status {
			status = "DONE"
		}

		if uncompleted && !val.status {
			fmt.Printf("%d: %s (%s).\n", idx+1, val.content, status)
		} else if !uncompleted {
			fmt.Printf("%d: %s (%s).\n", idx+1, val.content, status)
		} else {
			continue
		}
	}
	println("--------------")
}

func addTask(new_task Task) {

}

func delTask(task_id int) {

}

func doneTask(task_id int) {

}
