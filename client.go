package openmetadata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HostURL - Default Hashicups URL
const HostURL string = "http://192.168.0.19:8585"

// Client struct to hold API client configuration
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	AuthToken  string
}

// NewClient initializes and returns a new OpenMetadata API client
func NewClient(baseURL, authToken *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		BaseURL:    HostURL,
	}
	if baseURL != nil {
		c.BaseURL = *baseURL
	}
	if authToken != nil {
		c.AuthToken = *authToken
	}

	return &c, nil
}

// newRequest creates a new HTTP request with the appropriate headers
func (c *Client) newRequest(method, url string, authToken *string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	switch method {
	case "POST", "PUT":
		req.Header.Set("Content-Type", "application/json")
	case "PATCH":
		req.Header.Set("Content-Type", "application/merge-patch+json")
	default:
		req.Header.Set("Content-Type", "application/json") // デフォルトでJSONとする
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AuthToken))

	return req, nil
}

func (c *Client) doRequest(req *http.Request, out interface{}) (int, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if err := decodeBody(res, out); err != nil {
		return res.StatusCode, err
	}

	return res.StatusCode, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

// func (c *Client) doRequest(req *http.Request) ([]byte, int, error) {
// 	res, err := c.HTTPClient.Do(req)
// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, res.StatusCode, err
// 	}

// 	return body, res.StatusCode, nil
// }
