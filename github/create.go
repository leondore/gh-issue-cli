package github

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (c *Client) CreateIssue(repo string, issueBody *IssueBody) (Issue, error) {
	if len(repo) == 0 {
		return Issue{}, errors.New("you need to pass a repo name")
	}

	url := fmt.Sprintf("%s/%s/%s/issues", baseUrl, c.user, repo)

	params, err := json.Marshal(&issueBody)
	if err != nil {
		return Issue{}, err
	}

	data, err := c.Request(http.MethodPost, url, bytes.NewReader(params))
	if err != nil {
		return Issue{}, err
	}

	newIssue := Issue{}
	err = json.Unmarshal(data, &newIssue)
	if err != nil {
		return Issue{}, err
	}

	return newIssue, nil
}
