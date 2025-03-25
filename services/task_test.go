package services

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/spalqui/task-tracker-cli/mocks/repositories"
	"github.com/spalqui/task-tracker-cli/types"
)

func TestTaskService_Create(t *testing.T) {
	tests := []struct {
		name        string
		description string
		addFunc     func(task *types.Task) error
		wantErr     bool
		errMsg      string
	}{
		{
			name:        "successful creation",
			description: "Test task",
			addFunc:     func(task *types.Task) error { return nil },
			wantErr:     false,
		},
		{
			name:        "validation failure - description is empty",
			description: "",
			addFunc:     nil,
			wantErr:     true,
			errMsg:      "failed to create task due to validation errors: map[description:is empty]",
		},
		{
			name:        "repository failure",
			description: "Test task",
			addFunc:     func(task *types.Task) error { return errors.New("repository error") },
			wantErr:     true,
			errMsg:      "failed to create task",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &repositories.MockTaskRepository{
				AddFunc: tt.addFunc,
			}
			service := NewTaskService(mockRepo)

			createdTask, err := service.Create(tt.description)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				} else if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("expected error message to contain %q but got %q", tt.errMsg, err.Error())
				}
				if createdTask != nil {
					t.Errorf("expected createdTask to be nil but got %v", createdTask)
				}
			} else {
				if err != nil {
					t.Errorf("expected no error but got %v", err)
				}
				if createdTask == nil {
					t.Errorf("expected createdTask to be non-nil but got nil")
				} else if createdTask.Description != tt.description {
					t.Errorf("expected description %q but got %q", tt.description, createdTask.Description)
				}
			}
		})
	}
}

func TestTaskService_Update(t *testing.T) {
	tests := []struct {
		name        string
		taskID      int
		description string
		updateFunc  func(task *types.Task) error
		wantErr     bool
		errMsg      string
	}{
		{
			name:        "successful update",
			taskID:      1,
			description: "Updated task",
			updateFunc:  func(task *types.Task) error { return nil },
			wantErr:     false,
		},
		{
			name:        "validation failure - taskID is empty or zero",
			description: "buy milk",
			updateFunc:  nil,
			wantErr:     true,
			errMsg:      "failed to update task due to validation errors: map[ID:is zero or empty]",
		},
		{
			name:        "validation failure - description is empty",
			taskID:      1,
			description: "",
			updateFunc:  nil,
			wantErr:     true,
			errMsg:      "failed to update task due to validation errors: map[description:is empty]",
		},
		{
			name:        "repository failure",
			taskID:      1,
			description: "Updated task",
			updateFunc:  func(task *types.Task) error { return errors.New("repository error") },
			wantErr:     true,
			errMsg:      "failed to update task",
		},
		{
			name:        "task not found",
			taskID:      999,
			description: "Updated task",
			updateFunc:  func(task *types.Task) error { return errors.New("task with ID 999 not found") },
			wantErr:     true,
			errMsg:      "task with ID 999 not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &repositories.MockTaskRepository{
				UpdateFunc: tt.updateFunc,
			}
			service := NewTaskService(mockRepo)

			err := service.Update(tt.taskID, tt.description)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				} else if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("expected error message to contain %q but got %q", tt.errMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error but got %v", err)
				}
			}
		})
	}
}

func TestTaskService_Delete(t *testing.T) {
	tests := []struct {
		name       string
		taskID     int
		deleteFunc func(taskID int) error
		wantErr    bool
		errMsg     string
	}{
		{
			name:       "successful deletion",
			taskID:     1,
			deleteFunc: func(taskID int) error { return nil },
			wantErr:    false,
		},
		{
			name:       "validation failure - taskID is empty or zero",
			taskID:     0,
			deleteFunc: nil,
			wantErr:    true,
			errMsg:     "failed to delete task due to validation errors: map[ID:is zero or empty]",
		},
		{
			name:       "repository failure",
			taskID:     1,
			deleteFunc: func(taskID int) error { return errors.New("repository error") },
			wantErr:    true,
			errMsg:     "failed to delete task",
		},
		{
			name:       "task not found",
			taskID:     999,
			deleteFunc: func(taskID int) error { return errors.New("task with ID 999 not found") },
			wantErr:    true,
			errMsg:     "task with ID 999 not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &repositories.MockTaskRepository{
				DeleteFunc: tt.deleteFunc,
			}
			service := NewTaskService(mockRepo)

			err := service.Delete(tt.taskID)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				} else if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("expected error message to contain %q but got %q", tt.errMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error but got %v", err)
				}
			}
		})
	}
}

func TestTaskService_List(t *testing.T) {
	tests := []struct {
		name     string
		status   string
		listFunc func(status string) ([]*types.Task, error)
		wantErr  bool
		errMsg   string
		tasks    []*types.Task
	}{
		{
			name:   "successful list with empty status",
			status: "",
			listFunc: func(status string) ([]*types.Task, error) {
				return []*types.Task{
					{ID: 1, Description: "Task 1"},
					{ID: 2, Description: "Task 2"},
				}, nil
			},
			wantErr: false,
			tasks: []*types.Task{
				{ID: 1, Description: "Task 1"},
				{ID: 2, Description: "Task 2"},
			},
		},
		{
			name:   "repository failure",
			status: "",
			listFunc: func(status string) ([]*types.Task, error) {
				return nil, errors.New("repository error")
			},
			wantErr: true,
			errMsg:  "failed to list tasks",
		},
		{
			name:   "unsupported status",
			status: "unsupported",
			listFunc: func(status string) ([]*types.Task, error) {
				return nil, fmt.Errorf("failed to list tasks invalid task status: %s", status)
			},
			wantErr: true,
			errMsg:  "failed to list tasks invalid task status: unsupported",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &repositories.MockTaskRepository{
				ListFunc: tt.listFunc,
			}
			service := NewTaskService(mockRepo)

			tasks, err := service.List(tt.status)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				} else if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("expected error message to contain %q but got %q", tt.errMsg, err.Error())
				}
				if tasks != nil {
					t.Errorf("expected tasks to be nil but got %v", tasks)
				}
			} else {
				if err != nil {
					t.Errorf("expected no error but got %v", err)
				}
				if tasks == nil {
					t.Errorf("expected tasks to be non-nil but got nil")
				} else if len(tasks) != len(tt.tasks) {
					t.Errorf("expected %d tasks but got %d", len(tt.tasks), len(tasks))
				} else {
					for i, task := range tasks {
						if task.ID != tt.tasks[i].ID || task.Description != tt.tasks[i].Description || task.Status != tt.tasks[i].Status {
							t.Errorf("expected task %v but got %v", tt.tasks[i], task)
						}
					}
				}
			}
		})
	}
}
