package commands

import "fmt"

type listCommand struct{}

func NewListCommand() Command {
	return &listCommand{}
}

func (c *listCommand) Execute() error {
	fmt.Println("list command")
	return nil
}
