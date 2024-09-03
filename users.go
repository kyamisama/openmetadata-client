package openmetadata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateUser creates a new user in OpenMetadata
func (c *Client) CreateUser(user CreateUser_req, authToken *string) (*CreateUser_res, error) {
	// Initialize user data to send

	postJSON, err := json.Marshal(user)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user data: %w", err)
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/api/v1/users/", strings.NewReader(string(postJSON)))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	body, statusCode, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}
	// POST成功のレスポンスコードは201なので、StatusCreatedてエラーハンドリグする
	if statusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create user, status code: %d, response: %s", statusCode, string(body))
	}
	createdUser := CreateUser_res{}
	err = json.Unmarshal(body, &createdUser)
	if err != nil {
		return nil, err
	}
	return &createdUser, nil
}

// GetUser retrieves a user by ID from OpenMetadata
func (c *Client) GetUser(id string, authToken *string) (*GetUser_res, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/api/v1/users/"+id, nil)
	if err != nil {
		return nil, err
	}

	body, statusCode, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user, status code: %d", statusCode)
	}

	getUser := GetUser_res{}
	err = json.Unmarshal(body, &getUser)
	if err != nil {
		return nil, err
	}

	return &getUser, nil
}

// GetUser retrieves a user by ID from OpenMetadata
func (c *Client) GetUsers(authToken *string) (*GetUsers_res, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/api/v1/users", nil)
	if err != nil {
		return nil, err
	}

	body, statusCode, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user, status code: %d", statusCode)
	}

	getUsers := GetUsers_res{}
	err = json.Unmarshal(body, &getUsers)
	if err != nil {
		return nil, err
	}

	return &getUsers, nil
}

// UpdateUser updates an existing user in OpenMetadata
func (c *Client) UpdateUser(user UpdateUser_req, authToken *string) (*UpdateUser_res, error) {
	// Initialize user data to send
	postJSON, err := json.Marshal(user)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user data: %w", err)
	}

	req, err := http.NewRequest("PUT", c.BaseURL+"/api/v1/users", strings.NewReader(string(postJSON)))
	if err != nil {
		return nil, fmt.Errorf("failed to update request: %w", err)
	}

	body, statusCode, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to create user, status code: %d, response: %s", statusCode, string(body))
	}
	updatedUser := UpdateUser_res{}
	err = json.Unmarshal(body, &updatedUser)
	if err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

// DeleteUser deletes a user by ID from OpenMetadata
func (c *Client) DeleteUser(name string, authToken *string) error {
	req, err := http.NewRequest("DELETE", c.BaseURL+"/api/v1/users/name/"+name, nil)

	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("hardDelete", "true")
	req.URL.RawQuery = q.Encode()

	body, statusCode, err := c.doRequest(req, authToken)
	if err != nil {
		return err
	}
	// log.Println(statusCode)
	if statusCode != http.StatusOK {
		return fmt.Errorf("failed to deleted user, status code: %d, response: %s", statusCode, string(body))
	}

	// If status code is 204, return success without decoding
	return nil
}
