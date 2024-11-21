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

var choices = []string {
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

	fmt.Printf("Welcome To %s App\n", ToDO.title)
    main_loop:
    for {
        for idx, val := range choices {
            fmt.Printf("%d- %s\n", idx + 1, val)
        }
        var choice int
        print("choose an option: ")
        fmt.Scanf("%d\n", &choice)

        switch choice {
        case 1:
        case 2:
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

}

func addTask(new_task Task) {
    
}

func delTask(task_id int) {

}


func doneTask(task_id int) {
    
}