package main

import (
	"log"
	"os"

	"github.com/spalqui/task-tracker-cli/commands"
	"github.com/spalqui/task-tracker-cli/repositories"
	"github.com/spalqui/task-tracker-cli/services"
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
	repository, err := repositories.NewTaskRepository()
	if err != nil {
		log.Fatalf("failed to create task repository: %s", err)
	}

	// Setup services
	service := services.NewTaskService(repository)

	// Map commands
	commandMap := map[string]commands.Command{
		AddCommand:     commands.NewAddCommand(service),
		UpdateCommand:  commands.NewUpdateCommand(service),
		DeleteCommand:  commands.NewDeleteCommand(service),
		MarkDone:       commands.NewMarkAsDoneCommand(service),
		MarkInProgress: commands.NewMarkAsInProgressCommand(service),
		ListCommand:    commands.NewListCommand(service),
	}

	commandStr := args[0]
	commandArgs := make([]string, 0)
	if len(args) > 1 {
		commandArgs = args[1:]
	}

	command, exist := commandMap[commandStr]
	if !exist {
		log.Fatalf("unknown command: %s", commandStr)
	}

	err = command.Execute(commandArgs)
	if err != nil {
		log.Fatal(err)
	}
}
