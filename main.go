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

	// Setup repositories
	taskRepository, err := repositories.NewTaskRepository()
	if err != nil {
		log.Fatalf("failed to create task repository: %s", err)
	}

	// Setup services
	taskService := services.NewTaskService(taskRepository)

	command := args[0]
	commandArgs := args[1:]

	if command == AddCommand {
		description := commandArgs[0]

		task, err := taskService.Create(description)
		if err != nil {
			log.Fatalf("failed to create task: %s", err)
		}

		log.Printf("Task added successfully (ID: %d)", task.ID)
	}

	//
	//switch command {
	//case AddCommand:
	//
	//
	//case UpdateCommand:
	//	cmd.SetCommand(commands.Update, commandArgs...)
	//case DeleteCommand:
	//	cmd.SetCommand(commands.Delete, commandArgs...)
	//case ListCommand:
	//	cmd.SetCommand(commands.List, commandArgs...)
	//default:
	//	log.Fatalf("unknown command: %s", command)
	//}
	//
	//err := cmd.Execute()
	//if err != nil {
	//	log.Fatalf("command failed: %s err: %s", command, err)
	//}
}
