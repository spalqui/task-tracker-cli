package commands

import (
	"errors"
	"log"

	"github.com/spalqui/task-tracker-cli/services"
)

type Add struct {
	taskService services.TaskService
}

func NewAddCommand(taskService services.TaskService) *Add {
	return &Add{
		taskService: taskService,
	}
}

func (c *Add) Execute(args []string) error {
	if len(args) != 1 {
		return errors.New("invalid number arguments provided (description)")
	}

	description := args[0]

	task, err := c.taskService.Create(description)
	if err != nil {
		return err
	}

	log.Printf("task added successfully (ID: %d)", task.ID)
	return nil
}
