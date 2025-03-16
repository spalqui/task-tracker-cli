package commands

import "fmt"

type addCommand struct{}

func NewAddCommand() Command {
	return &addCommand{}
}

func (c *addCommand) Execute() error {
	fmt.Println("add command")
	return nil
}
