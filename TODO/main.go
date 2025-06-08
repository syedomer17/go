package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Title string
	Done  bool
}

const fileName = "tasks.json"

func loadTasks() ([]Task, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func listTasks() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("❌ Error loading tasks:", err)
		return
	}
	if len(tasks) == 0 {
		fmt.Println("📭 No tasks yet.")
		return
	}
	for i, task := range tasks {
		status := "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Title)
	}
}

func addTask(title string) {
	tasks, _ := loadTasks()
	tasks = append(tasks, Task{Title: title, Done: false})
	saveTasks(tasks)
	fmt.Println("➕ Task added:", title)
}

func markDone(index int) {
	tasks, err := loadTasks()
	if err != nil || index < 1 || index > len(tasks) {
		fmt.Println("❌ Invalid task number")
		return
	}
	tasks[index-1].Done = true
	saveTasks(tasks)
	fmt.Println("✅ Marked as done:", tasks[index-1].Title)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [add|list|done] [task]")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("❗ Provide a task to add")
			return
		}
		addTask(os.Args[2])
	case "list":
		listTasks()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("❗ Provide task number to mark as done")
			return
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("❌ Invalid number")
			return
		}
		markDone(index)
	default:
		fmt.Println("❓ Unknown command. Use add, list, or done.")
	}
}
