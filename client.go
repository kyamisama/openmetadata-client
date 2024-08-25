package openmetadata

import (
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
