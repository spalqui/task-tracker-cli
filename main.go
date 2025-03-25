package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spalqui/task-tracker-cli/repositories"
	"github.com/spalqui/task-tracker-cli/services"
	"github.com/spalqui/task-tracker-cli/types"
)

const (
	AddCommand     = "add"
	UpdateCommand  = "update"
	MarkInProgress = "mark-in-progress"
	MarkDone       = "mark-done"
	DeleteCommand  = "delete"
	ListCommand    = "list"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("no command provided")
	}

	var err error

	// Setup repositories
	taskRepository, err := repositories.NewTaskRepository()
	if err != nil {
		log.Fatalf("failed to create task repository: %s", err)
	}

	// Setup services
	taskService := services.NewTaskService(taskRepository)

	command := args[0]
	commandArgs := make([]string, 0)
	if len(args) > 1 {
		commandArgs = args[1:]
	}

	switch command {
	case AddCommand:
		if len(commandArgs) < 1 {
			log.Fatalf("not enough arguments provided (description)")
		}

		description := commandArgs[0]

		task, err := taskService.Create(description)
		if err != nil {
			log.Fatalf("failed to create task: %s", err)
		}

		log.Printf("task added successfully (ID: %d)", task.ID)
	case UpdateCommand:
		if len(commandArgs) < 2 {
			log.Fatalf("not enough arguments provided (id, description)")
		}

		taskIDValue := commandArgs[0]
		description := commandArgs[1]

		taskID, err := strconv.Atoi(taskIDValue)
		if err != nil {
			log.Fatalf("failed to parse task ID: %s", err)
		}

		err = taskService.Update(taskID, description)
		if err != nil {
			log.Fatalf("failed to update task: %s", err)
		}

		log.Printf("task updated successfully (ID: %d)", taskID)

	case DeleteCommand:
		if len(commandArgs) < 1 {
			log.Fatalf("not enough arguments provided (id)")
		}

		taskIDValue := commandArgs[0]

		taskID, err := strconv.Atoi(taskIDValue)
		if err != nil {
			log.Fatalf("failed to parse task ID: %s", err)
		}

		err = taskService.Delete(taskID)
		if err != nil {
			log.Fatalf("failed to delete task: %s", err)
		}

		log.Print("task deleted successfully")

	case ListCommand:
		if len(commandArgs) > 1 {
			log.Fatalf("too many arguments provided (status [todo, done, in-progress])")
		}

		status := ""
		if len(commandArgs) == 1 {
			status = commandArgs[0]
		}

		tasks, err := taskService.List(status)
		if err != nil {
			log.Fatalf("failed to list tasks: %s", err)
		}

		title := "all tasks"
		switch status {
		case types.TaskStatusTodo:
			title = "all todo tasks"
		case types.TaskStatusInProgress:
			title = "all in-progress tasks"
		case types.TaskStatusDone:
			title = "all done tasks"
		}
		title = fmt.Sprintf("%s (%d)", title, len(tasks))

		log.Println(title)
		for _, task := range tasks {
			log.Printf("%d\t%s\t%s\t%s\t%s", task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	default:
		log.Fatalf("unknown command: %s", command)
	}
}

func getArgs(args []string) []string {
	if len(args) == 0 {
		return nil
	}
	return args[1:]
}
