package services

import (
	"fmt"

	"github.com/spalqui/task-tracker-cli/repositories"
	"github.com/spalqui/task-tracker-cli/types"
	"github.com/spalqui/task-tracker-cli/validator"
)

const (
	ErrCreateTask           = "error creating task: %v"
	ErrUpdateTask           = "error updating task: %v"
	ErrMarkTaskAsDone       = "error marking task as done: %v"
	ErrMarkTaskAsInProgress = "error marking task as in-progress: %v"
	ErrDeleteTask           = "error deleting task: %v"
	ErrListTasks            = "error listing tasks: %v"
)

type TaskService interface {
	Create(description string) (*types.Task, error)
	Update(taskID int, description string) error
	MarkAsDone(taskID int) error
	MarkAsInProgress(taskID int) error
	Delete(taskID int) error
	List(status string) ([]*types.Task, error)
}

func NewTaskService(taskRepository repositories.TaskRepository) TaskService {
	return &taskService{
		taskRepository: taskRepository,
	}
}

type taskService struct {
	taskRepository repositories.TaskRepository
}

func (s *taskService) Create(description string) (*types.Task, error) {
	task := &types.Task{
		Description: description,
	}

	v := validator.New()
	v.Check(description != "", "description", "is empty")

	if !v.IsValid() {
		return nil, fmt.Errorf(ErrCreateTask, v.Errors)
	}

	err := s.taskRepository.Add(task)
	if err != nil {
		return nil, fmt.Errorf(ErrCreateTask, err)
	}

	return task, nil
}

func (s *taskService) Update(taskID int, description string) error {
	v := validator.New()
	v.Check(taskID > 0, "ID", "is zero or empty")
	v.Check(description != "", "description", "is empty")

	if !v.IsValid() {
		return fmt.Errorf(ErrUpdateTask, v.Errors)
	}

	task := types.Task{
		ID:          taskID,
		Description: description,
	}

	err := s.taskRepository.Update(&task)
	if err != nil {
		return fmt.Errorf(ErrUpdateTask, err)
	}

	return nil
}

func (s *taskService) MarkAsDone(taskID int) error {
	v := validator.New()
	v.Check(taskID > 0, "ID", "is zero or empty")

	if !v.IsValid() {
		return fmt.Errorf(ErrMarkTaskAsDone, v.Errors)
	}

	task, err := s.taskRepository.GetByID(taskID)
	if err != nil {
		return fmt.Errorf(ErrMarkTaskAsDone, err)
	}

	task.Status = types.TaskStatusDone

	err = s.taskRepository.Update(task)
	if err != nil {
		return fmt.Errorf(ErrMarkTaskAsDone, err)
	}

	return nil
}

func (s *taskService) MarkAsInProgress(taskID int) error {
	v := validator.New()
	v.Check(taskID > 0, "ID", "is zero or empty")

	if !v.IsValid() {
		return fmt.Errorf(ErrMarkTaskAsInProgress, v.Errors)
	}

	task, err := s.taskRepository.GetByID(taskID)
	if err != nil {
		return fmt.Errorf(ErrMarkTaskAsInProgress, err)
	}

	task.Status = types.TaskStatusInProgress

	err = s.taskRepository.Update(task)
	if err != nil {
		return fmt.Errorf(ErrMarkTaskAsInProgress, err)
	}

	return nil
}

func (s *taskService) Delete(taskID int) error {
	v := validator.New()
	v.Check(taskID > 0, "ID", "is zero or empty")

	if !v.IsValid() {
		return fmt.Errorf(ErrDeleteTask, v.Errors)
	}

	err := s.taskRepository.Delete(taskID)
	if err != nil {
		return fmt.Errorf(ErrDeleteTask, err)
	}

	return nil
}

func (s *taskService) List(status string) ([]*types.Task, error) {
	if status != "" {
		v := validator.New()
		if !v.In(status, types.TaskStatusTodo, types.TaskStatusInProgress, types.TaskStatusDone) {
			return nil, fmt.Errorf(ErrListTasks, status)
		}
	}

	tasks, err := s.taskRepository.List(status)
	if err != nil {
		return nil, fmt.Errorf(ErrListTasks, err)
	}

	return tasks, nil
}
