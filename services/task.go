package services

import (
	"fmt"

	"github.com/spalqui/task-tracker-cli/repositories"
	"github.com/spalqui/task-tracker-cli/types"
)

type TaskService interface {
	Create(description string) (*types.Task, error)
	Update(taskID int, description string) error
	MarkAsDone(taskID int) error
	MarkAsInProgress(taskID int) error
	Delete(taskID int) error
	List() ([]*types.Task, error)
	ListAllDone() ([]*types.Task, error)
	ListAllTodo() ([]*types.Task, error)
	ListAllInProgress() ([]*types.Task, error)
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

	err := s.taskRepository.Add(task)
	if err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	return task, nil
}

func (s *taskService) Update(taskID int, description string) error {
	//TODO implement me
	panic("implement me")
}

func (s *taskService) MarkAsDone(taskID int) error {
	//TODO implement me
	panic("implement me")
}

func (s *taskService) MarkAsInProgress(taskID int) error {
	//TODO implement me
	panic("implement me")
}

func (s *taskService) Delete(taskID int) error {
	//TODO implement me
	panic("implement me")
}

func (s *taskService) List() ([]*types.Task, error) {
	//TODO implement me
	panic("implement me")
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
