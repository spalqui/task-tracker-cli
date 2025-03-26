package commands

import (
	"errors"
	"log"

	"github.com/spalqui/task-tracker-cli/services"
)

type MarkAsInProgress struct {
	taskService services.TaskService
}

func NewMarkAsInProgressCommand(taskService services.TaskService) *MarkAsInProgress {
	return &MarkAsInProgress{
		taskService: taskService,
	}
}

func (c *MarkAsInProgress) Execute(args []string) error {
	if len(args) != 1 {
		return errors.New("invalid number arguments provided (id)")
	}

	taskID, err := getTaskID(args[0])
	if err != nil {
		return err
	}

	err = c.taskService.MarkAsInProgress(taskID)
	if err != nil {
		return err
	}

	log.Print("task marked in-progress successfully")
	return nil
}
