package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Name      string
	Completed bool
}

func main() {
	tasks := []Task{}

	for {
		showMenu()
		option := getUserInput("Enter option: ")

		switch option {
		case "1":
			showTasks(tasks)
		case "2":
			addTask(&tasks)
		case "3":
			saveTasks(tasks)
		case "4":
			completeTask(&tasks)
		case "5":
			fmt.Println("exiting app")
			return
		default:
			fmt.Println("invalid")
		}
	}
}

func saveTasks(tasks []Task) {
	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "X"
		}
		file.WriteString(fmt.Sprintf("[%s] %s\n", status, task.Name))
	}

	fmt.Println("Tasks saved to file 'file.txt'.")
}

func completeTask(tasks *[]Task) {
	showTasks(*tasks)
	taskIndexString := getUserInput("Enter the task number to mark as complete: ")
	taskIndex, err := strconv.Atoi(taskIndexString)
	if err != nil || taskIndex < 1 || len(*tasks) < taskIndex {
		fmt.Println("Invalid task number. Please try again.")
		return
	}
	(*tasks)[taskIndex-1].Completed = true
	fmt.Println("task has been marked as completed")
}

func addTask(tasks *[]Task) {
	taskName := getUserInput("Enter task: ")
	*tasks = append(*tasks, Task{Name: taskName})
	fmt.Println("task added")
}

func showTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("no tasks available to display")
		return
	}
	fmt.Println("Tasks: ")

	for i, task := range tasks {
		status := ""
		if task.Completed {
			status = "X"
		}
		fmt.Printf("%d. [%s] %s \n", i+1, status, task.Name)
	}
}

func showMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. Show Tasks")
	fmt.Println("2. Add Task")
	fmt.Println("3. Save Tasks to File")
	fmt.Println("4. Mark Task as Completed")
	fmt.Println("5. Exit")
}

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
