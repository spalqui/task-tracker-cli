package main

import (
	"log"
	"os"

	"github.com/spalqui/task-tracker-cli/commands"
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
		return
	}

	command := args[0]
	//commandArgs := args[1:]
	var cmd commands.Command

	switch command {
	case AddCommand:
		cmd = commands.NewAddCommand()
	case UpdateCommand:
		cmd = commands.NewUpdateCommand()
	case DeleteCommand:
		cmd = commands.NewDeleteCommand()
	case ListCommand:
		cmd = commands.NewListCommand()
	default:
		log.Fatalf("unknown command: %s", command)
	}

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("command failed: %s err: %s", command, err)
	}
}
