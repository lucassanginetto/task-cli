package main

import (
	"fmt"
	"os"
	"time"
)

type Task struct {
	Id          uint
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func main() {
	if len(os.Args) < 2 {
		// print usage
		os.Exit(1)
	}

	action := os.Args[1]
	switch action {
	case "add":
		fmt.Println("Adding a new task")
		// if JSON file exists
		//   get tasks array from file
		// else
		//   create empty tasks array
		// add task with provided description to array
		// write tasks array to JSON file
		// add task to JSON file
	case "update":
		fmt.Println("Updating a task")
		// get tasks array from JSON file
		// get task with provided ID from array
		// update task with provided description
		// write tasks array to JSON file
	case "delete":
		fmt.Println("Deleting a task")
		// get tasks array from JSON file
		// get task with provided ID from array
		// delete task with provided ID
		// write tasks array to JSON file
	case "mark-in-progress":
		fmt.Println("Marking a task as in progress")
		// get tasks array from JSON file
		// get task with provided ID from array
		// set status of task to "in-progress"
		// write tasks array to JSON file
	case "mark-done":
		fmt.Println("Marking a task as done")
		// get tasks array from JSON file
		// get task with provided ID from array
		// set status of task to "done"
		// write tasks array to JSON file
	case "list":
		fmt.Println("Listing tasks")
		// set only to ""
		// if status filter was provided
		//   switch status
		//     case "done"
		//       set only to "done"
		//     case "todo"
		//       set only to "todo"
		//     case "in-progress"
		//       set only to "in-progress"
		//     default
		//       warn about invalid status filter
		// get tasks array from JSON file
		// for each task
		//   if only set and only equal to status of task
		//     print task
	default:
		fmt.Fprintf(os.Stderr, "Invalid action \"%s\"\n", action)
		os.Exit(1)
	}
}
