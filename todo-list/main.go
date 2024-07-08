package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Title     string
	Completed bool
}

var tasks []Task

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Println("Enter your choice:")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addTask(reader)
		case "2":
			viewTask()
		case "3":
			completeTask(reader)
		case "4":
			deleteTask(reader)
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func addTask(reader *bufio.Reader) {
	fmt.Print("Enter task title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	tasks = append(tasks, Task{Title: title, Completed: false})
	fmt.Println("Task added successfully")
}

func viewTask() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	fmt.Println("Tasks:")
	for i, task := range tasks {
		status := "Incomplete"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.Title, status)
	}
}

func completeTask(reader *bufio.Reader) {
	fmt.Print("Enter the task number to mark as completed: ")
	var tasknumber int
	fmt.Scanln(&tasknumber)
	if tasknumber > 0 && tasknumber <= len(tasks) {
		tasks[tasknumber-1].Completed = true
		fmt.Println("Task marked as completed!")
	} else {
		fmt.Println("Invalid task number.")
	}
}

func deleteTask(reader *bufio.Reader) {
	fmt.Print("Enter the task number to delete: ")
	var tasknumber int
	fmt.Scanln(&tasknumber)
	if tasknumber > 0 && tasknumber < len(tasks) {
		tasks = append(tasks[:tasknumber-1], tasks[tasknumber:]...)
		fmt.Println("Task deleted successfully!")
	} else {
		fmt.Println("Invalid task number.")
	}
}
