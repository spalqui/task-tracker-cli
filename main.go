package main

import (
	"log"
	"os"

	"github.com/spalqui/task-tracker-cli/repositories"
	"github.com/spalqui/task-tracker-cli/services"
)

const (
	AddCommand    = "add"
	UpdateCommand = "update"
	DeleteCommand = "delete"
	ListCommand   = "list"
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
		if len(commandArgs) == 0 {
			log.Fatalf("no arguments provided")
		}

		description := commandArgs[0]

		task, err := taskService.Create(description)
		if err != nil {
			log.Fatalf("failed to create task: %s", err)
		}

		log.Printf("task added successfully (ID: %d)", task.ID)
	case UpdateCommand:
		log.Fatalf("not yet implemented")
	case DeleteCommand:
		log.Fatalf("not yet implemented")
	case ListCommand:
		log.Fatalf("not yet implemented")
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
