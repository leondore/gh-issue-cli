package github

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseUrl = "https://api.github.com/repos"

type Client struct {
	client http.Client
	token  string
	user   string
}

func NewClient(token, user string) Client {
	return Client{
		client: http.Client{
			Timeout: 5 * time.Second,
		},
		token: token,
		user:  user,
	}
}

func (c *Client) Request(method string, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("an error occurred while processing the request: %s", resp.Status)
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
