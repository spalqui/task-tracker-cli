package services

import (
	"fmt"

	"github.com/spalqui/task-tracker-cli/repositories"
	"github.com/spalqui/task-tracker-cli/types"
	"github.com/spalqui/task-tracker-cli/validator"
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
		return nil, fmt.Errorf("failed to create task due to validation errors: %v", v.Errors)
	}

	err := s.taskRepository.Add(task)
	if err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	return task, nil
}

func (s *taskService) Update(taskID int, description string) error {
	v := validator.New()

	v.Check(taskID > 0, "ID", "is zero or empty")
	v.Check(description != "", "description", "is empty")

	if !v.IsValid() {
		return fmt.Errorf("failed to update task due to validation errors: %v", v.Errors)
	}

	task := types.Task{
		ID:          taskID,
		Description: description,
	}

	err := s.taskRepository.Update(&task)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	return nil
}

func (s *taskService) MarkAsDone(taskID int) error {
	//TODO implement me
	panic("implement me")
}

func (s *taskService) MarkAsInProgress(taskID int) error {
	// TODO implement me
	panic("implement me")
}

func (s *taskService) Delete(taskID int) error {
	v := validator.New()

	v.Check(taskID > 0, "ID", "is zero or empty")

	if !v.IsValid() {
		return fmt.Errorf("failed to delete task due to validation errors: %v", v.Errors)
	}

	err := s.taskRepository.Delete(taskID)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	return nil
}

func (s *taskService) List(status string) ([]*types.Task, error) {
	if status != "" {
		v := validator.New()

		if !v.In(status, types.TaskStatusTodo, types.TaskStatusInProgress, types.TaskStatusDone) {
			return nil, fmt.Errorf("failed to list tasks invalid task status: %s", status)
		}
	}

	tasks, err := s.taskRepository.List(status)
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}

	return tasks, nil
}

func (s *taskService) ListAllDone() ([]*types.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (s *taskService) ListAllTodo() ([]*types.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (s *taskService) ListAllInProgress() ([]*types.Task, error) {
	//TODO implement me
	panic("implement me")
}
