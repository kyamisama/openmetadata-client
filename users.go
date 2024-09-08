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

	// req, err := http.NewRequest("POST", c.BaseURL+"/api/v1/users/", strings.NewReader(string(postJSON)))
	req, err := c.newRequest("POST", c.BaseURL+"/api/v1/users/", authToken, strings.NewReader(string(postJSON)))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	var createUser_res CreateUser_res
	statusCode, err := c.doRequest(req, &createUser_res)
	if err != nil {
		return nil, err
	}
	// POST成功のレスポンスコードは201なので、StatusCreatedてエラーハンドリグする
	if statusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create user, status code: %d", statusCode)
	}
	return &createUser_res, nil
}

// GetUser retrieves a user by ID from OpenMetadata
func (c *Client) GetUser(id string, authToken *string) (*GetUser_res, error) {
	req, err := c.newRequest("GET", c.BaseURL+"/api/v1/users/"+id, authToken, nil)
	if err != nil {
		return nil, err
	}
	var getUser GetUser_res
	statusCode, err := c.doRequest(req, &getUser)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user, status code: %d", statusCode)
	}

	return &getUser, nil
}

// GetUser retrieves a user by ID from OpenMetadata
func (c *Client) GetUsers(authToken *string) (*GetUsers_res, error) {
	req, err := c.newRequest("GET", c.BaseURL+"/api/v1/users/", authToken, nil)
	if err != nil {
		return nil, err
	}
	var getUsers GetUsers_res
	statusCode, err := c.doRequest(req, &getUsers)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user, status code: %d", statusCode)
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

	req, err := c.newRequest("PUT", c.BaseURL+"/api/v1/users", authToken, strings.NewReader(string(postJSON)))
	if err != nil {
		return nil, fmt.Errorf("failed to update request: %w", err)
	}

	var updateUser_res UpdateUser_res
	statusCode, err := c.doRequest(req, &updateUser_res)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to create user, status code: %d", statusCode)
	}

	return &updateUser_res, nil
}

func (c *Client) PatchUser(patchdata []PatchUser_req, id string, authToken *string) (*PatchUser_res, error) {
	postJSON, err := json.Marshal(patchdata)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user data: %w", err)
	}
	req, err := c.newRequest("PATCH", c.BaseURL+"/api/v1/users/"+id, authToken, strings.NewReader(string(postJSON)))
	if err != nil {
		return nil, fmt.Errorf("failed to patch request: %w", err)
	}

	var patchUser_res PatchUser_res
	statusCode, err := c.doRequest(req, &patchUser_res)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to patch user, status code: %d", statusCode)
	}

	return &patchUser_res, nil
}

// DeleteUser deletes a user by ID from OpenMetadata
func (c *Client) DeleteUser(name string, authToken *string) (*DeleteUser, error) {
	req, err := c.newRequest("DELETE", c.BaseURL+"/api/v1/users/name/"+name, authToken, nil)

	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("hardDelete", "true")
	req.URL.RawQuery = q.Encode()
	var deleteUser DeleteUser
	statusCode, err := c.doRequest(req, &deleteUser)
	if err != nil {
		return nil, err
	}
	// log.Println(statusCode)
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to deleted user, status code: %d", statusCode)
	}

	// If status code is 204, return success without decoding
	return &deleteUser, nil
}
