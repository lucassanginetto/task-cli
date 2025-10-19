package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
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

func writeTasksToFile(tasks []Task) error {
	byteValue, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile("tasks.json", byteValue, 0644)
	if err != nil {
		return err
	}
	return nil
}

func idFromArgs() uint {
	id, err := strconv.ParseUint(os.Args[2], 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return uint(id)
}

func markTask(status string) {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "The ID of the task is necessary for marking it as \"in progress\"")
		os.Exit(1)
	}

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

	id := idFromArgs()

	found := false
	for i := range tasks {
		if tasks[i].Id == id {
			tasks[i].Status = status
			found = true
			break
		}
	}
	if !found {
		fmt.Fprintln(os.Stderr, "Task not found")
		os.Exit(1)
	}

	writeTasksToFile(tasks)
}

func printHelp(w io.Writer) {
	fmt.Fprintf(
		w,
		"usage: %s <command> [<args>]\n"+
			"commands:\n"+
			"   add <description>           Create a new task with given description\n"+
			"   update <id> <description>   Change the description of a task\n"+
			"   delete <id>                 Delete a task\n"+
			"   mark-in-progress <id>       Mark a task as \"in progress\"\n"+
			"   mark-done <id>              Mark a task as \"done\"\n"+
			"   list [<status>]             List tasks\n"+
			"   help                        Display this\n",
		os.Args[0],
	)
}

func main() {
	if len(os.Args) < 2 {
		printHelp(os.Stderr)
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

		if err := writeTasksToFile(tasks); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}

	case "update":
		if len(os.Args) < 4 {
			fmt.Fprintln(os.Stderr, "The ID of the task and a new description are necessary for updating a task")
			os.Exit(1)
		}

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

		id := idFromArgs()
		description := os.Args[3]

		found := false
		for i := range tasks {
			if tasks[i].Id == id {
				tasks[i].Description = description
				tasks[i].UpdatedAt = time.Now()
				found = true
				break
			}
		}
		if !found {
			fmt.Fprintln(os.Stderr, "Task not found")
			os.Exit(1)
		}

		writeTasksToFile(tasks)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "The ID of the task is necessary for deleting a task")
			os.Exit(1)
		}

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

		id := idFromArgs()

		found := false
		for i := range tasks {
			if tasks[i].Id == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			fmt.Fprintln(os.Stderr, "Task not found")
			os.Exit(1)
		}

		writeTasksToFile(tasks)

	case "mark-in-progress":
		markTask("in-progress")

	case "mark-done":
		markTask("done")

	case "list":
		status := ""
		if len(os.Args) > 2 {
			switch os.Args[2] {
			case "todo", "in-progress", "done":
				status = os.Args[2]
			default:
				fmt.Fprintf(os.Stderr, "Invalid status \"%s\"\n", os.Args[2])
				os.Exit(1)
			}
		}

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

		fmt.Println("")
		for _, t := range tasks {
			if status == "" || t.Status == status {
				fmt.Printf(
					"Task %d: %s\nStatus: %s\nCreated at: %s\nUpdated at: %s\n\n",
					t.Id,
					t.Description,
					t.Status,
					t.CreatedAt.String()[:19],
					t.UpdatedAt.String()[:19],
				)
			}
		}

	case "help":
		printHelp(os.Stdout)

	default:
		fmt.Fprintf(os.Stderr, "Invalid action \"%s\"\n", action)
		os.Exit(1)
	}
}
