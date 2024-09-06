package openmetadata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) CreateDBService(db CreateDB_req, authToken *string) (*CreateDB_res, error) {
	postJSON, err := json.Marshal(db)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal db data: %w", err)
	}
	req, err := c.newRequest("POST", c.BaseURL+"/api/v1/services/databaseServices", authToken, strings.NewReader(string(postJSON)))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	var createDB_res CreateDB_res
	statusCode, err := c.doRequest(req, &createDB_res)
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create db_service, status code: %d", statusCode)

	}

	return &createDB_res, nil
}

func (c *Client) UpdateDBService(db UpdateDB_req, authToken *string) (*UpdateDB_res, error) {
	postJSON, err := json.Marshal(db)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal db data: %w", err)
	}
	req, err := c.newRequest("PUT", c.BaseURL+"/api/v1/services/databaseServices", authToken, strings.NewReader(string(postJSON)))
	if err != nil {
		return nil, fmt.Errorf("failed to db request: %w", err)
	}
	var updateDB_res UpdateDB_res
	statusCode, err := c.doRequest(req, &updateDB_res)
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to Update DBService, status code: %d", statusCode)
	}

	return &updateDB_res, nil
}

func (c *Client) DeleteDBService(name string, authToken *string) (*DeleteDB, error) {
	req, err := c.newRequest("DELETE", c.BaseURL+"/api/v1/services/databaseServices/name/"+name, authToken, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("hardDelete", "true")
	req.URL.RawQuery = q.Encode()
	var deleteDB DeleteDB
	statusCode, err := c.doRequest(req, &deleteDB)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to deleted DBService, status code: %d", statusCode)
	}
	return &deleteDB, nil
}
