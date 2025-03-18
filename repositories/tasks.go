package repositories

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/spalqui/task-tracker-cli/types"
)

const fileName = "tasks.json"

type TaskRepository interface {
	Add(task types.Task) error
	Update(task types.Task) error
	Delete(taskID int) error
	List() ([]types.Task, error)
}

type taskRepository struct {
	mu       sync.Mutex
	filePath string
}

func NewTaskRepository() (TaskRepository, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current directory: %w", err)
	}

	filePath := filepath.Join(dir, fileName)
	err = os.WriteFile(filePath, []byte{}, 0600)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}

	return &taskRepository{
		filePath: filePath,
	}, nil
}

func (r *taskRepository) Add(task types.Task) error {
	//TODO implement me
	panic("implement me")
}

func (r *taskRepository) Update(task types.Task) error {
	//TODO implement me
	panic("implement me")
}

func (r *taskRepository) Delete(taskID int) error {
	//TODO implement me
	panic("implement me")
}

func (r *taskRepository) List() ([]types.Task, error) {
	//TODO implement me
	panic("implement me")
}
