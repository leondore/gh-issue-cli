package github

type Issue struct {
	Id      int    `json:"id"`
	State   string `json:"state"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	HTMLUrl string `json:"html_url"`
	User    struct {
		Login string `json:"login"`
	} `json:"user"`
	Labels []struct {
		Name string `json:"name"`
	} `json:"labels"`
}

type IssueBody struct {
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Labels []string `json:"labels"`
}
