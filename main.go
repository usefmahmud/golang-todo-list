package main

import (
	"fmt"
)

type Task struct {
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
		{"the first task", false},
		{"the second task", true},
		{"the third task", false},
		{"the fourh task", true},
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
            println("what the task you want to add?")
            var task_content string
            fmt.Scanf("%s", &task_content)

            addTask(ToDO, createTask(task_content))
            
            println("Your task has been added successfully!")
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
	completedPercentage := getCompletedPercentage(list)
	fmt.Printf("Your progress is %.0f%%\n", completedPercentage*100)

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

func getCompletedPercentage(list List) float64 {
	var completedCount float64 = 0
	for _, val := range list.tasks {
		if val.status {
			completedCount++
		}
	}
	return completedCount / float64(len(list.tasks))
}

func addTask(list List, new_task Task) {
    list.tasks = append(list.tasks, new_task)
}

func createTask(content string) Task {
    return Task{content, false}
}

func delTask(task_id int) {

}

func doneTask(task_id int) {

}
