package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

type Task struct {
	Id          uint      `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func tasksFromFile() ([]Task, error) {
	tasks := []Task{}

	jsonFile, err := os.Open("tasks.json")
	if err == nil {
		byteValue, _ := io.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &tasks)
		jsonFile.Close()
	}

	return tasks, err
}

func main() {
	if len(os.Args) < 2 {
		// print usage
		os.Exit(1)
	}

	action := os.Args[1]
	switch action {
	case "add":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "A description is necessary for the new task")
			os.Exit(1)
		}

		tasks, err := tasksFromFile()
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}

		description := os.Args[2]
		tasksLen := len(tasks)
		now := time.Now()
		if tasksLen > 0 {
			tasks = append(
				tasks,
				Task{
					Id:          tasks[tasksLen-1].Id + 1,
					Description: description,
					Status:      "todo",
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			)
		} else {
			tasks = append(
				tasks,
				Task{
					Id:          1,
					Description: description,
					Status:      "todo",
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			)
		}

		byteValue, err := json.Marshal(tasks)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
		err = os.WriteFile("tasks.json", byteValue, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}

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

		tasks, err := tasksFromFile()
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				fmt.Fprintln(os.Stderr, "No \"tasks.json\" file was found in the current directory")
				os.Exit(1)
			} else {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		}

		tasksLen := len(tasks)
		for i, t := range tasks {
			fmt.Printf(
				"Task %d: %s\nStatus: %s\nCreated at: %s\nUpdated at: %s\n",
				t.Id,
				t.Description,
				t.Status,
				t.CreatedAt.String()[:19],
				t.UpdatedAt.String()[:19],
			)
			if i < tasksLen-1 {
				fmt.Println("")
			}
		}

	default:
		fmt.Fprintf(os.Stderr, "Invalid action \"%s\"\n", action)
		os.Exit(1)
	}
}
