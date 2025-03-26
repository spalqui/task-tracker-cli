package commands

import (
	"errors"
	"log"

	"github.com/spalqui/task-tracker-cli/services"
)

type Delete struct {
	taskService services.TaskService
}

func NewDeleteCommand(taskService services.TaskService) *Delete {
	return &Delete{
		taskService: taskService,
	}
}

func (c *Delete) Execute(args []string) error {
	if len(args) != 1 {
		return errors.New("invalid number arguments provided (id)")
	}

	taskID, err := getTaskID(args[0])
	if err != nil {
		return err
	}

	err = c.taskService.Delete(taskID)
	if err != nil {
		return err
	}

	log.Print("task deleted successfully")
	return nil
}
