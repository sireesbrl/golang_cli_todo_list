package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	tasks := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Todo List!")

	for {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Add a task")
		fmt.Println("2. View tasks")
		fmt.Println("3. Delete a task")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var task string
			fmt.Print("Enter the task: ")
			scanner.Scan()
			task = scanner.Text()
			tasks = append(tasks, task)
			fmt.Println("Task added successfully!")
		case 2:
			if len(tasks) == 0 {
				fmt.Println("No tasks found.")
			} else {
				fmt.Println("Tasks:")
				for i, task := range tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}
		case 3:
			if len(tasks) == 0 {
				fmt.Println("No tasks found.")
			} else {
				fmt.Println("Select a task to delete:")
				for i, task := range tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
				var taskIndex int
				fmt.Print("Enter the task number to delete: ")
				fmt.Scan(&taskIndex)
				if taskIndex > 0 && taskIndex <= len(tasks) {
					tasks = append(tasks[:taskIndex-1], tasks[taskIndex:]...)
					fmt.Println("Task deleted successfully!")
				} else {
					fmt.Println("Invalid task number.")
				}
			}
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
