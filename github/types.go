package github

import "strings"

type Labels []struct {
	Name string `json:"name"`
}

type Issue struct {
	Id      int    `json:"id"`
	State   string `json:"state"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	HTMLUrl string `json:"html_url"`
	User    struct {
		Login string `json:"login"`
	} `json:"user"`
	Labels Labels `json:"labels"`
}

type IssueBody struct {
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Labels []string `json:"labels"`
}

func LabelsRespToBody(labels Labels) []string {
	labelList := []string{}

	for _, label := range labels {
		labelList = append(labelList, label.Name)
	}

	return labelList
}

func LabelsBufferToBody(labels string) []string {
	labelList := []string{}

	if len(labels) > 0 {
		for _, label := range strings.Split(labels, ",") {
			labelList = append(labelList, strings.Trim(label, " "))
		}
	}

	return labelList
}
