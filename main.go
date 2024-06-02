package main

import (
	"fmt"
)

func main() {
	var task_list TaskList

	loadTasks(&task_list)
	startupMessage()

	for {
		var input string
		fmt.Scanln(&input)
		switch input {
		case "help":
			helpMessage()
		case "clear":
			fmt.Print("\033[H\033[2J")
		case "add":
			fmt.Println("Add task")
		case "list":
			listTasks(task_list)
		case "done":
			fmt.Println("Done tasks")
		case "delete":
			fmt.Println("Delete tasks")
		case "exit":
			fmt.Println("Goodbye!")
			return
		}
	}
}
