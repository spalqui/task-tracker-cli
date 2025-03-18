package repositories

import "github.com/spalqui/task-tracker-cli/types"

type TaskRepository interface {
	Add(task types.Task) error
	Update(task types.Task) error
	Delete(taskID int) error
	List() ([]types.Task, error)
}

type taskRepository struct {
}

func NewTaskRepository() TaskRepository {
	return &taskRepository{}
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
