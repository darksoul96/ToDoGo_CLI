package main

import (
	"bufio"
	"fmt"
	"os"
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
			in := bufio.NewScanner(os.Stdin)
			in.Scan()
			title = in.Text()
			addTask(&task_list, title)
			fmt.Println("Task added successfully!")
		case "list":
			listTasks(task_list)
		case "done":
			fmt.Println("Enter task id")
			var id int
			fmt.Scanln(&id)
			if markAsDone(&task_list, id) {
				fmt.Println("Task marked as done!")
			} else {
				fmt.Println("Invalid task id")
			}
		case "delete":
			fmt.Println("Enter task id")
			var id int
			fmt.Scanln(&id)
			if deleteTask(&task_list, id) {
				fmt.Println("Task deleted successfully!")
			} else {
				fmt.Println("Invalid task id")
			}
		case "savefile":
			fmt.Println("Saving tasks to tasks.json")
			saveToFile(task_list)
		case "exit":
			fmt.Print("\033[H\033[2J")
			fmt.Println("Goodbye!")
			return
		}
		fmt.Println()
	}
}
