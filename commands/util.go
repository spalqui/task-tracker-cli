package commands

import (
	"fmt"
	"strconv"
)

func getTaskID(v string) (int, error) {
	ID, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("error parsing ID(%s): %w", v, err)
	}

	return ID, nil
}
