package commands

import (
	"errors"
	"log"

	"github.com/spalqui/task-tracker-cli/services"
)

type MarkAsDone struct {
	taskService services.TaskService
}

func NewMarkAsDoneCommand(taskService services.TaskService) *MarkAsDone {
	return &MarkAsDone{
		taskService: taskService,
	}
}

func (c *MarkAsDone) Execute(args []string) error {
	if len(args) != 1 {
		return errors.New("invalid number arguments provided (id)")
	}

	taskID, err := getTaskID(args[0])
	if err != nil {
		return err
	}

	err = c.taskService.MarkAsDone(taskID)
	if err != nil {
		return err
	}

	log.Print("task marked as done successfully")
	return nil
}
