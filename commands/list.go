package commands

import (
	"errors"
	"fmt"
	"log"

	"github.com/spalqui/task-tracker-cli/services"
	"github.com/spalqui/task-tracker-cli/types"
)

type List struct {
	taskService services.TaskService
}

func NewListCommand(taskService services.TaskService) *List {
	return &List{
		taskService: taskService,
	}
}

func (c *List) Execute(args []string) error {
	if len(args) > 1 {
		return errors.New("invalid number arguments provided (status [todo, done, in-progress] or not argument for all)")
	}

	status := ""
	if len(args) == 1 {
		status = args[0]
	}

	tasks, err := c.taskService.List(status)
	if err != nil {
		return fmt.Errorf("failed to list tasks: %w", err)
	}

	title := "all tasks"
	switch status {
	case types.TaskStatusTodo:
		title = "all todo tasks"
	case types.TaskStatusInProgress:
		title = "all in-progress tasks"
	case types.TaskStatusDone:
		title = "all done tasks"
	}
	title = fmt.Sprintf("%s (%d)", title, len(tasks))

	log.Println(title)
	for _, task := range tasks {
		log.Printf("%d\t%s\t%s\t%s\t%s", task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	}
	return nil
}
