package kroger

import (
	"fmt"
	"net/http"
)

type Client struct {
	Token  string
	Client *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		Token:  token,
		Client: &http.Client{},
	}
}

func (c *Client) Get(endpoint string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
