package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"exercise-4.11/github"
)

func commandUpdate(c *config, params []string) error {
	if len(params) < 3 {
		return errors.New("you need to pass a repo name, issue id and field to edit")
	}

	issue, err := c.client.ReadIssue(params[0], params[1])
	if err != nil {
		return err
	}

	updateParams := github.IssueBody{
		Title:  issue.Title,
		Body:   issue.Body,
		Labels: github.LabelsRespToBody(issue.Labels),
	}
	field := params[2]
	reader := bufio.NewScanner(os.Stdin)

	switch field {
	case "title":
		fmt.Printf("Current title: %s\n", issue.Title)
		fmt.Print("Enter a new title: ")
		reader.Scan()
		input := reader.Text()
		if len(input) == 0 {
			return errors.New("an issue requires a title")
		}
		updateParams.Title = input
	case "body":
		fmt.Println("Press ENTER to edit the body in your preferred text editor...")
		reader.Scan()
		body, err := openEditor(c, issue.Body)
		if err != nil {
			return err
		}
		updateParams.Body = string(body)
	case "labels":
		fmt.Printf("Current labels: %s\n", github.LabelsRespToBody(issue.Labels))
		fmt.Print("Replace current labels, as a comma-separated list (leave empty to remove all labels): ")
		reader.Scan()
		input := reader.Text()
		updateParams.Labels = github.LabelsBufferToBody(input)
	default:
		return errors.New("you can only update the title, body and labels fields")
	}

	response, err := c.client.UpdateIssue(params[0], params[1], &updateParams)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Issue has been succesfully updated!")
	fmt.Printf("You can view the changes here: %s\n", response.HTMLUrl)

	return nil
}
