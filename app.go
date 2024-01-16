package main

import (
	"errors"
	"fmt"
	"strconv"

	tasks "example.com/git/user"
)

func main() {
	taskManager := tasks.NewTaskManager()

	for {
		operation()
		action, err := getUserData("What can I do for you? (Type the corresponding number or 'exit' to quit)")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		switch action {
		case "0":
			taskManager.ListTasks()
		case "1":
			title, err := getUserData("Enter task title:")
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			taskManager.AddTask(title)
		case "2":
			taskID, err := getUserData("Enter task ID to delete:")
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			id, err := parseID(taskID)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			taskManager.RemoveTask(id)
		case "3":
			taskID, err := getUserData("Enter task ID to mark as complete:")
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			id, err := parseID(taskID)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			// Mark the task as complete
			taskManager.MarkTaskComplete(id)
		case "exit":
			fmt.Println("Exiting Task Manager.")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}

		// Prompt for another action
		anotherAction, err := getUserData("Do you want to perform another action? (yes/no): ")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if anotherAction != "yes" {
			fmt.Println("Exiting Task Manager.")
			return
		}
	}
}

func parseID(input string) (int, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("invalid task ID")
	}
	return id, nil
}

func getUserData(input string) (string, error) {
	fmt.Print(input)
	var value string
	fmt.Scan(&value)
	if value == "" {
		return "", errors.New("you have to type something")
	}
	return value, nil
}

func operation() {
	fmt.Println("Welcome to the Task manager!")
	fmt.Println("0. My tasks")
	fmt.Println("1. Create new task")
	fmt.Println("2. Delete task")
	fmt.Println("3. Mark as complete the task")
}
