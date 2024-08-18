package github

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (c *Client) UpdateIssue(repo, issueId string, body interface{}) (Issue, error) {
	if len(repo) == 0 || len(issueId) == 0 {
		return Issue{}, errors.New("you need to pass a repo name and issue id")
	}

	url := fmt.Sprintf("%s/%s/%s/issues/%s", baseUrl, c.user, repo, issueId)

	params, err := json.Marshal(&body)
	if err != nil {
		return Issue{}, err
	}

	data, err := c.Request(http.MethodPatch, url, bytes.NewReader(params))
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
