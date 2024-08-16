package main

import (
	"fmt"
)

func commandHelp(c *config, params []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Github issue manager!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, props := range genCommands() {
		fmt.Printf("%s: %s\n", props.command, props.description)
	}
	fmt.Println()
	return nil
}
