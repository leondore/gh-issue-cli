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

	updateParams := github.IssueBody{}
	field := params[2]
	reader := bufio.NewScanner(os.Stdin)

	switch field {
	case "title":
		fmt.Printf("Current title: %s\n", issue.Title)
		fmt.Print("Enter a new title: ")
		reader.Scan()
		title := reader.Text()
		if len(title) == 0 {
			return errors.New("an issue requires a title")
		}
		updateParams.Title = title
	case "body":
		fmt.Println("Press ENTER to edit the body in your preferred text editor...")
		reader.Scan()
		body, err := openEditor(c, issue.Body)
		if err != nil {
			return err
		}
		updateParams.Body = string(body)
	}

	// fmt.Print("Add some labels (as a comma-separated list): ")
	// reader.Scan()
	// labels := reader.Text()

	// labelList := []string{}
	// if len(labels) > 0 {
	// 	labelSlice := strings.Split(labels, ",")
	// 	for _, label := range labelSlice {
	// 		labelList = append(labelList, strings.Trim(label, " "))
	// 	}
	// }

	// bodyParams := github.IssueBody{
	// 	Title:  title,
	// 	Body:   string(body),
	// 	Labels: labelList,
	// }

	// response, err := c.client.CreateIssue(params[0], &bodyParams)
	// if err != nil {
	// 	return err
	// }

	fmt.Printf("Test: %v\n", updateParams)

	return nil
}
