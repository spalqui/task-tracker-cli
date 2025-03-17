package commands

import "fmt"

func List(params ...string) error {
	fmt.Println("list command")
	return nil
}
