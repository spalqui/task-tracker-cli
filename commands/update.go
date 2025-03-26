package commands

import (
	"errors"
	"log"

	"github.com/spalqui/task-tracker-cli/services"
)

type Update struct {
	taskService services.TaskService
}

func NewUpdateCommand(taskService services.TaskService) *Update {
	return &Update{
		taskService: taskService,
	}
}

func (c *Update) Execute(args []string) error {
	if len(args) != 2 {
		return errors.New("invalid number arguments provided (id, description)")
	}

	taskID, err := getTaskID(args[0])
	if err != nil {
		return err
	}

	description := args[1]

	err = c.taskService.Update(taskID, description)
	if err != nil {
		return err
	}

	log.Printf("task updated successfully (ID: %d)", taskID)
	return nil
}
