package openmetadata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateUser creates a new user in OpenMetadata
func (c *Client) CreateUser(user *CreateUser) (*CreateUser, error) {
	// Initialize user data to send
	postData := &CreateUser{
		Name:        user.Name,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		Description: user.Description,
		Password:    user.Password,
	}

	postJSON, err := json.Marshal(postData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user data: %w", err)
	}
	fmt.Printf("[+] %s\n", string(postJSON))

	req, err := http.NewRequest("POST", c.BaseURL+"/", bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AuthToken))
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	//fmt.Println(string(respBody))

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create user, status code: %d, response: %s", resp.StatusCode, string(respBody))
	}

	var createUser CreateUser
	err = json.NewDecoder(bytes.NewReader(respBody)).Decode(&createUser)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &createUser, nil
}

// GetUser retrieves a user by ID from OpenMetadata
func (c *Client) GetUser(name string) (*GetUser, error) {
	req, err := c.newRequest("GET", "/name/"+name, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user, status code: %d", resp.StatusCode)
	}

	var user GetUser
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUser retrieves a user by ID from OpenMetadata
func (c *Client) GetUsers() (*GetUser, error) {
	req, err := c.newRequest("GET", "/", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user, status code: %d", resp.StatusCode)
	}

	var user GetUser
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates an existing user in OpenMetadata
func (c *Client) UpdateUser(user *UpdateUser) (*UpdateUser, error) {
	// Initialize user data to send
	postData := &UpdateUser{
		Name:        user.Name,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		Description: user.Description,
		Password:    user.Password,
	}
	postJSON, err := json.Marshal(postData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user data: %w", err)
	}
	fmt.Printf("[+] %s\n", string(postJSON))

	req, err := http.NewRequest("PUT", c.BaseURL+"/", bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to update request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AuthToken))
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	//respBody, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to read response body: %w", err)
	//}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		var apiError struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&apiError); err != nil {
			return nil, fmt.Errorf("failed to decode error response: %w", err)
		}
		return nil, fmt.Errorf("failed to update user, status code: %d, error: %s", resp.StatusCode, apiError.Message)
	}

	var updatedUser UpdateUser
	err = json.NewDecoder(resp.Body).Decode(&updatedUser)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}
	log.Println(&updatedUser)
	return &updatedUser, nil
}

// DeleteUser deletes a user by ID from OpenMetadata
func (c *Client) DeleteUser(name string) (*DeleteUser, error) {
	req, err := c.newRequest("DELETE", "/name/"+name+"?hardDelete=true", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		var apiError struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&apiError); err != nil {
			return nil, fmt.Errorf("failed to decode error response: %w", err)
		}
		return nil, fmt.Errorf("failed to delete user, status code: %d, error: %s", resp.StatusCode, apiError.Message)
	}

	// If status code is 204, return success without decoding
	return nil, nil // Return nil because there's no user to return on successful deletion
}
