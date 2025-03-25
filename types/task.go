package types

import "time"

type Task struct {
	ID          int
	Description string
	Status      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

const (
	TaskStatusTODO = iota
	TaskStatusInProgress
	TaskStatusDone
)
