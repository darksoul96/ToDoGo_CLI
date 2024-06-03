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

func StartupMessage() {
	fmt.Println("Welcome back to To Do List CLI!")
	fmt.Println("Type 'help' for more information")
}

func HelpMessage() {
	fmt.Println("Type 'clear' to clear the console")
	fmt.Println("Type 'add' to add a new task")
	fmt.Println("Type 'list' to list all tasks")
	fmt.Println("Type 'done' to mark a task as done")
	fmt.Println("Type 'delete' to delete a task")
	fmt.Println("Type 'savefile' to save tasks to tasks.json")
	fmt.Println("Type 'exit' to exit the program")
}

func LoadTasks(task_list *TaskList) {
	file, err := os.OpenFile("tasks.json", os.O_RDONLY, 0644)

	if err != nil {
		//if file not found, then create the file
		_, create_error := os.Create("tasks.json")
		if create_error != nil {
			panic(create_error)
		}
		fmt.Println("tasks.json file created")
	} else {
		defer file.Close()
		//unmarshall json file
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&task_list)
		if err != nil {
			fmt.Println("tasks.json file not found")
		} else {
			fmt.Println("Loaded tasks from tasks.json")
		}
	}
}

func ListTasks(task_list TaskList) {
	if len(task_list.Tasks) == 0 {
		fmt.Println("No tasks found")
	} else {
		for i, task := range task_list.Tasks {
			var status string
			if task.Done {
				status = "done"
			} else {
				status = "pending"
			}
			fmt.Printf("%d. %s - %s\n", i+1, task.Title, status)
		}
	}
}

func AddTask(task_list *TaskList, title string) {
	task := Task{ID: int64(len(task_list.Tasks) + 1), Title: title, Done: false}
	task_list.Tasks = append(task_list.Tasks, task)
}

func DeleteTask(task_list *TaskList, id int) bool {
	//if id not found, return false
	if id < 1 || id > len(task_list.Tasks) {
		return false
	}
	task_list.Tasks = append(task_list.Tasks[:id-1], task_list.Tasks[id:]...)
	return true
}

func SaveToFile(task_list TaskList) {
	file, err := os.OpenFile("tasks.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	} else {
		defer file.Close()
		encoder := json.NewEncoder(file)
		err = encoder.Encode(&task_list)
		if err != nil {
			panic(err)
		}
		fmt.Println("Saved tasks to tasks.json")
	}
}

func MarkAsDone(task_list *TaskList, id int) bool {
	if id < 1 || id > len(task_list.Tasks) {
		return false
	}
	task_list.Tasks[id-1].Done = true
	return true
}
