package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandRead(c *config, params []string) error {
	if len(params) < 2 {
		return errors.New("you need to pass a repo name and issue id")
	}

	issue, err := c.client.ReadIssue(params[0], params[1])
	if err != nil {
		return err
	}

	var labels strings.Builder
	for i, label := range issue.Labels {
		if i != 0 {
			labels.WriteString(", ")
		}
		labels.WriteString(label.Name)
	}

	fmt.Printf("Issue: %s | Status: %s\n", issue.Title, issue.State)
	fmt.Printf("Creator: %s\n", issue.User.Login)
	fmt.Printf("\n%s\n\n", issue.Body)
	fmt.Printf("Labels: %s\n", labels.String())

	return nil
}
