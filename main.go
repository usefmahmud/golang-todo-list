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
		case 1: // show all tasks
			showTasks(ToDO, false)
		case 2: // show only uncompleted tasks
			showTasks(ToDO, true)
		case 3: // mark task as done
            println("enter the id number of the task you need to done")

            var target_id int
            fmt.Scanf("%d", &target_id)

            doneTask(&ToDO, target_id)
		case 4: // add new task
			println("what the task you want to add?")
			var task_content string
			fmt.Scanf("%s", &task_content)

			addTask(&ToDO, createTask(task_content))

			println("Your task has been added successfully!")
		case 5: // delete a task
			continue
		case 6: // quit the app
			println("END..")
			break main_loop
		default: // wrong choice
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

func addTask(list *List, new_task Task) {
    list.tasks = append(list.tasks, new_task)
}

func createTask(content string) Task {
    return Task{content, false}
}

func delTask(task_id int) {

}

func doneTask(list *List, task_id int) {
	if task_id >= len(list.tasks) {
		return
	}

	list.tasks[task_id-1].status = true
}
