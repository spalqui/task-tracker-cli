package commands

import "fmt"

type updateCommand struct{}

func NewUpdateCommand() Command {
	return &updateCommand{}
}

func (c *updateCommand) Execute() error {
	fmt.Println("update command")
	return nil
}
