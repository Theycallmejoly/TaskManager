package tasks

import (
	"fmt"
)

type Task struct {
	ID     int
	Title  string
	Status bool
}

type TaskManager struct {
	Tasks []Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (tm *TaskManager) AddTask(title string) {
	taskID := len(tm.Tasks) + 1
	task := Task{
		ID:     taskID,
		Title:  title,
		Status: false,
	}
	tm.Tasks = append(tm.Tasks, task)
	fmt.Printf("Task added: %s (ID: %d)\n", title, taskID)
}

func (tm *TaskManager) MarkTaskComplete(taskID int) {
	for i := range tm.Tasks {
		if tm.Tasks[i].ID == taskID {
			tm.Tasks[i].Status = true
			fmt.Printf("Task Marked as Complete: %s (ID: %d)\n", tm.Tasks[i].Title, taskID)
			return
		}
	}

	fmt.Printf("Task with ID %d not found\n", taskID)
}

func (tm *TaskManager) ListTasks() {
	if len(tm.Tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("Tasks:")
	for _, task := range tm.Tasks {
		fmt.Printf("ID: %d, Title: %s, Status: %t\n", task.ID, task.Title, task.Status)
	}
}

func (tm *TaskManager) RemoveTask(taskID int) {
	index := -1
	for i, task := range tm.Tasks {
		if task.ID == taskID {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Task not found.")
		return
	}

	tm.Tasks = append(tm.Tasks[:index], tm.Tasks[index+1:]...)
	fmt.Println("Task has been Removed.")
}
