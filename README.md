# task-cli

This is a command-line task tracker. You can use it to manage your tasks and track which ones you need to do, which ones you're currently working on, and which ones you've done.

This was created as a solution for the [Task Tracker project on roadmap.sh](https://roadmap.sh/projects/task-tracker), using the Go programming language.

# Installation

```sh
git clone https://github.com/lucassanginetto/task-cli.git

cd task-cli
```

Make sure you have [Go](https://go.dev/) installed for compiling, and then run:

```sh
go install
```

The package will be installed to `$GOPATH/bin`, and you'll be able to run it using `go run task-cli`.

If you want to be able to run it using just `task-cli`, add the `$GOPATH/bin` folder to your `$PATH`.

# Usage

To add tasks you can run

```sh
task-cli add "First task"
task-cli add "Another task"
```

A `tasks.json` file will be created inside the current working directory, if it doesn't exist.

To list the tasks in the file, you can run

```sh
task-cli list
```

The previously created tasks will be shown, each with their number ID, description, status ("todo", "in-progress" or "done") and the creation and last update timestamps.

If you want to change the description of a task you can run

```sh
task-cli update 2 "Second task"
```

To delete a task you can run

```sh
task-cli delete 2
```

To mark a task as "in progress" and later as "done" you can run

```sh
task-cli mark-in-progress 1
task-cli mark-done 1
```

To list only tasks with a specific status you can run

```sh
task-cli list todo
task-cli list in-progress
task-cli list done
```
