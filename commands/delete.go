package commands

import "fmt"

type deleteCommand struct{}

func NewDeleteCommand() Command {
	return &deleteCommand{}
}

func (c *deleteCommand) Execute() error {
	fmt.Println("delete command")
	return nil
}
