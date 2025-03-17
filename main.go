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
	commandArgs := args[1:]
	cmd := commands.NewCommander()

	switch command {
	case AddCommand:
		cmd.SetCommand(commands.Add, commandArgs...)
	case UpdateCommand:
		cmd.SetCommand(commands.Update, commandArgs...)
	case DeleteCommand:
		cmd.SetCommand(commands.Delete, commandArgs...)
	case ListCommand:
		cmd.SetCommand(commands.List, commandArgs...)
	default:
		log.Fatalf("unknown command: %s", command)
	}

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("command failed: %s err: %s", command, err)
	}
}
