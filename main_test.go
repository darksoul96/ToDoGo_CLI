package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTask(t *testing.T) {
	var task_list TaskList
	task_list.Tasks = make([]Task, 0)
	AddTask(&task_list, "test")

	if len(task_list.Tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(task_list.Tasks))
	}

	if task_list.Tasks[0].Title != "test" {
		t.Errorf("Expected task title to be 'test', got %s", task_list.Tasks[0].Title)
	}

	if task_list.Tasks[0].Done != false {
		t.Errorf("Expected task done to be false, got %t", task_list.Tasks[0].Done)
	}

	if task_list.Tasks[0].ID != 1 {
		t.Errorf("Expected task ID to be 1, got %d", task_list.Tasks[0].ID)
	}
}

func TestDeleteTask(t *testing.T) {
	var task_list TaskList
	task_list.Tasks = make([]Task, 0)

	task_list.Tasks = append(task_list.Tasks, Task{ID: 1, Title: "test", Done: false})

	assert.Equal(t, DeleteTask(&task_list, 2), false)
	assert.Equal(t, DeleteTask(&task_list, 1), true)

	if len(task_list.Tasks) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(task_list.Tasks))
	}

}

func TestMarkAsDone(t *testing.T) {
	var task_list TaskList
	task_list.Tasks = make([]Task, 0)

	task_list.Tasks = append(task_list.Tasks, Task{ID: 1, Title: "test", Done: false})

	assert.Equal(t, MarkAsDone(&task_list, 2), false)
	assert.Equal(t, MarkAsDone(&task_list, 1), true)

	if task_list.Tasks[0].Done != true {
		t.Errorf("Expected task done to be true, got %t", task_list.Tasks[0].Done)
	}
}

func RunTests(t *testing.T) {
	t.Run("AddTask", TestAddTask)
	t.Run("DeleteTask", TestDeleteTask)
	t.Run("MarkAsDone", TestMarkAsDone)
}
