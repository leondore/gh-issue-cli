package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"exercise-4.11/github"
)

func commandCreate(c *config, params []string) error {
	if len(params) < 1 {
		return errors.New("you need to pass a repo name")
	}

	reader := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a title: ")
	reader.Scan()
	title := reader.Text()
	if len(title) == 0 {
		return errors.New("you must add a title to create an issue")
	}

	body, err := openEditor(c, "[DELETE ME] Enter the content for the issue body. Save and exit to continue forward.")
	if err != nil {
		return err
	}
	if len(body) == 0 {
		return errors.New("you must add a body to create an issue")
	}

	fmt.Print("Add some labels (as a comma-separated list): ")
	reader.Scan()
	labels := reader.Text()

	labelList := []string{}
	if len(labels) > 0 {
		labelSlice := strings.Split(labels, ",")
		for _, label := range labelSlice {
			labelList = append(labelList, strings.Trim(label, " "))
		}
	}

	bodyParams := github.IssueBody{
		Title:  title,
		Body:   string(body),
		Labels: labelList,
	}

	response, err := c.client.CreateIssue(params[0], &bodyParams)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Your new issue was succesfully created!")
	fmt.Printf("You can find it here: %s\n", response.HTMLUrl)

	return nil
}
