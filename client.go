package openmetadata

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client struct to hold API client configuration
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	AuthToken  string
}

// NewClient initializes and returns a new OpenMetadata API client
func NewClient(baseURL, authToken string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		AuthToken: authToken,
	}
}

// newRequest creates a new HTTP request with the appropriate headers
func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	req, err := http.NewRequest(method, c.BaseURL+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, int, error) {
	token := c.AuthToken

	if authToken != nil {
		token = *authToken
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}

	return body, res.StatusCode, nil
}
