package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (c *Client) ReadIssue(repo, issueId string) (Issue, error) {
	if len(repo) == 0 || len(issueId) == 0 {
		return Issue{}, errors.New("you need to pass a repo name and issue id")
	}

	url := fmt.Sprintf("%s/%s/%s/issues/%s", baseUrl, c.user, repo, issueId)

	data, err := c.Request(http.MethodGet, url, nil)
	if err != nil {
		return Issue{}, err
	}

	issue := Issue{}
	err = json.Unmarshal(data, &issue)
	if err != nil {
		return Issue{}, err
	}

	return issue, nil
}
