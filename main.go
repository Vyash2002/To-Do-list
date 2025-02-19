package main

//Import necessary files
import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Defining the Task structure
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks []Task

const fileName = "task.json"

// 1. Loading Tasks From file
// Reads tasks.json file
func loadTasks() {
	file, err := os.ReadFile(fileName) //If the file exists, it decodes JSON data into the tasks slice.
	if err == nil {                    //If the file does not exist (first run), tasks remains empty.
		json.Unmarshal(file, &tasks)
	}
}

// 2. Saving Tasks to file
// Converts tasks slice into formatted JSON (MarshalIndent makes it human-readable).
// Saves the JSON data into tasks.json.
// The file permission 0644 means it is readable and writable by the owner and readable by others.
func saveTasks() {
	file, _ := json.MarshalIndent(tasks, "", " ")
	os.WriteFile(fileName, file, 0644)
}

// 3 Adding a New Task
// Assigns a new unique ID (1 more than the current task count).
// Adds the task to tasks slice.
// Saves the updated list to file.
func addTask(title string) {
	id := len(tasks) + 1
	tasks = append(tasks, Task{ID: id, Title: title, Done: false})
	saveTasks()
	fmt.Println("Task added successfully")

}

// 4 Listing All Tasks
// If no tasks exist, prints "No tasks found."
// Otherwise, prints a formatted task list:
// [ ] → Incomplete tasks.
// [✓] → Completed tasks.
func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No Tasks found.")
		return
	}
	fmt.Println("To-Do List:")
	for _, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[✓]"
		}
		fmt.Printf("%d. %s %s\n", task.ID, status, task.Title)

	}
}

// 5 Marking a Task as Done
// Searches for the task by ID.
// If found, sets Done = true and saves changes.
// If not found, prints "Task not found."
func markDone(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			saveTasks()
			fmt.Println("Task marked as done!")
			return
		}
	}
	fmt.Println("Task not found")
}

// 6 Removing a Task
// Finds the task by ID.
// Removes it from the tasks slice using append().
// Saves the updated list.
func removeTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks()
			fmt.Println("Task removed successfully")
			return

		}
	}
	fmt.Println("Task not found.")
}

// 7 Command Line Interface(CLI)
func main() {
	loadTasks()

	if len(os.Args) < 2 {
		fmt.Println("Usage: todo add|lis|done|remove [task_id/task_title]")
		return
	}
	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task title.")
			return
		}
		title := os.Args[2]
		addTask(title)
	case "list":
		listTasks()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID.")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		markDone(id)
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID.")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		removeTask(id)
	default:
		fmt.Println("Unknown command.")
	}

}
