package repositories

import (
	"github.com/spalqui/task-tracker-cli/types"
)

type MockTaskRepository struct {
	GetByIDFunc func(taskID int) (*types.Task, error)
	AddFunc     func(t *types.Task) error
	UpdateFunc  func(task *types.Task) error
	DeleteFunc  func(taskID int) error
	ListFunc    func(status string) ([]*types.Task, error)
}

func (m *MockTaskRepository) GetByID(taskID int) (*types.Task, error) {
	return m.GetByIDFunc(taskID)
}

func (m *MockTaskRepository) Add(task *types.Task) error {
	return m.AddFunc(task)
}

func (m *MockTaskRepository) Update(task *types.Task) error {
	return m.UpdateFunc(task)
}

func (m *MockTaskRepository) Delete(taskID int) error {
	return m.DeleteFunc(taskID)
}

func (m *MockTaskRepository) List(status string) ([]*types.Task, error) {
	return m.ListFunc(status)
}
