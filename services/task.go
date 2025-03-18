package services

import (
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

func (t taskService) Create(description string) (*types.Task, error) {
	return &types.Task{
		ID: 1,
	}, nil
}

func (t taskService) Update(taskID int, description string) error {
	//TODO implement me
	panic("implement me")
}

func (t taskService) MarkAsDone(taskID int) error {
	//TODO implement me
	panic("implement me")
}

func (t taskService) MarkAsInProgress(taskID int) error {
	//TODO implement me
	panic("implement me")
}

func (t taskService) Delete(taskID int) error {
	//TODO implement me
	panic("implement me")
}

func (t taskService) List() ([]*types.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t taskService) ListAllDone() ([]*types.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t taskService) ListAllTodo() ([]*types.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t taskService) ListAllInProgress() ([]*types.Task, error) {
	//TODO implement me
	panic("implement me")
}
