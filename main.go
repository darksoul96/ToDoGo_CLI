package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var task_list TaskList

	StartupMessage()
	LoadTasks(&task_list)

	for {
		var input string
		fmt.Scanln(&input)
		switch input {
		case "help":
			HelpMessage()
		case "clear":
			fmt.Print("\033[H\033[2J")
		case "add":
			fmt.Println("Enter task name")
			var title string
			in := bufio.NewScanner(os.Stdin)
			in.Scan()
			title = in.Text()
			AddTask(&task_list, title)
			fmt.Println("Task added successfully!")
		case "list":
			ListTasks(task_list)
		case "done":
			fmt.Println("Enter task id")
			var id int
			fmt.Scanln(&id)
			if MarkAsDone(&task_list, id) {
				fmt.Println("Task marked as done!")
			} else {
				fmt.Println("Invalid task id")
			}
		case "delete":
			fmt.Println("Enter task id")
			var id int
			fmt.Scanln(&id)
			if DeleteTask(&task_list, id) {
				fmt.Println("Task deleted successfully!")
			} else {
				fmt.Println("Invalid task id")
			}
		case "savefile":
			fmt.Println("Saving tasks to tasks.json")
			SaveToFile(task_list)
		case "exit":
			fmt.Print("\033[H\033[2J")
			fmt.Println("Goodbye!")
			return
		}
		fmt.Println()
	}
}
