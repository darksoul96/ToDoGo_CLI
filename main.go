package main

import (
	"fmt"
)

func main() {
	var task_list TaskList

	startupMessage()
	loadTasks(&task_list)

	for {
		var input string
		fmt.Scanln(&input)
		switch input {
		case "help":
			helpMessage()
		case "clear":
			fmt.Print("\033[H\033[2J")
		case "add":
			fmt.Println("Enter task name")
			var title string
			fmt.Scanln(&title)
			addTask(&task_list, title)
		case "list":
			listTasks(task_list)
		case "done":
			fmt.Println("Enter task id")
			var id int
			fmt.Scanln(&id)
			markAsDone(&task_list, id)
		case "delete":
			fmt.Println("Enter task id")
			var id int
			fmt.Scanln(&id)
			deleteTask(&task_list, id)
		case "savefile":
			fmt.Println("Saving tasks to tasks.json")
			saveToFile(task_list)
		case "exit":
			fmt.Println("Goodbye!")
			return
		}
	}
}
