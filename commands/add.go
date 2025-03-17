package commands

import "fmt"

func Add(params ...string) error {
	if len(params) != 1 {
		return fmt.Errorf("add command requires 1 parameter (description)")
	}

	description := params[0]
	if description == "" {
		return fmt.Errorf("description is required")
	}

	fmt.Println("add command")
	return nil
}
