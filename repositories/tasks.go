package repositories

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/spalqui/task-tracker-cli/types"
)

const fileName = "tasks.json"

type TaskRepository interface {
	Add(task *types.Task) error
	Update(task *types.Task) error
	Delete(taskID int) error
	List() ([]*types.Task, error)
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

	_, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.WriteFile(filePath, []byte("[]"), 0600)
			if err != nil {
				return nil, fmt.Errorf("failed to create %s file: %w", fileName, err)
			}
		} else {
			return nil, fmt.Errorf("failed to read %s file: %w", fileName, err)
		}
	}

	return &taskRepository{
		filePath: filePath,
	}, nil
}

func (r *taskRepository) Add(task *types.Task) error {
	tasks, err := r.readTasks()
	if err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}

	lastID := len(tasks)
	timeNow := time.Now()

	task.ID = lastID + 1
	task.Status = "TODO"
	task.CreatedAt = timeNow
	task.UpdatedAt = timeNow

	tasks = append(tasks, task)

	err = r.writeTasks(tasks)
	if err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}

	return nil
}

func (r *taskRepository) Update(task *types.Task) error {
	tasks, err := r.readTasks()
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == task.ID {
			tasks[i].UpdatedAt = time.Now()
			tasks[i].Description = task.Description
			found = true

			err = r.writeTasks(tasks)
		}
	}
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	if !found {
		return fmt.Errorf("failed to find task with ID %v", task.ID)
	}

	return nil
}

func (r *taskRepository) Delete(taskID int) error {
	//TODO implement me
	panic("implement me")
}

func (r *taskRepository) List() ([]*types.Task, error) {
	tasks, err := r.readTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}

	return tasks, nil
}

func (r *taskRepository) readTasks() ([]*types.Task, error) {
	file, err := r.readFile()
	if err != nil {
		return nil, fmt.Errorf("failed to read tasks: %w", err)
	}

	var tasks []*types.Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return nil, fmt.Errorf("failed to read tasks: %w", err)
	}

	return tasks, nil
}

func (r *taskRepository) writeTasks(tasks []*types.Task) error {
	fileBytes, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("failed to write tasks: %w", err)
	}

	err = r.writeFile(fileBytes)
	if err != nil {
		return fmt.Errorf("failed to write tasks: %w", err)
	}

	return nil
}

func (r *taskRepository) readFile() ([]byte, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	fileBytes, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s file: %w", fileName, err)
	}

	return fileBytes, nil
}

func (r *taskRepository) writeFile(data []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	err := os.WriteFile(r.filePath, data, 0600)
	if err != nil {
		return fmt.Errorf("failed to write %s file: %w", fileName, err)
	}

	return nil
}
