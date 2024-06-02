package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func startupMessage() {
	fmt.Println("Welcome back to To Do List CLI!")
	fmt.Println("Type 'help' for more information")
}

func helpMessage() {
	fmt.Println("Type 'clear' to clear the console")
	fmt.Println("Type 'add' to add a new task")
	fmt.Println("Type 'list' to list all tasks")
	fmt.Println("Type 'done' to mark a task as done")
	fmt.Println("Type 'delete' to delete a task")
	fmt.Println("Type 'exit' to exit the program")
}

func loadTasks(task_list *TaskList) {
	file, err := os.OpenFile("tasks.json", os.O_RDONLY, 0644)

	if err != nil {
		panic(err)
	} else {
		defer file.Close()
		//unmarshall json file
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&task_list)
		if err != nil {
			panic(err)
		}
		fmt.Println("Loaded tasks from tasks.json")
	}
}

func listTasks(task_list TaskList) {
	for i, task := range task_list.Tasks {
		fmt.Printf("%d. %s\n", i+1, task.Title)
	}
}
