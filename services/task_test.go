package services

import (
	"errors"
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
