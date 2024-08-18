package main

import (
	"errors"
	"fmt"
)

func commandClose(c *config, params []string) error {
	if len(params) < 2 {
		return errors.New("you need to pass a repo name and issue id")
	}

	type body struct {
		State string `json:"state"`
	}

	response, err := c.client.UpdateIssue(params[0], params[1], &body{State: "closed"})
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("You've closed this issue as completed.")
	fmt.Printf("You can still access the issue here: %s\n", response.HTMLUrl)

	return nil
}
