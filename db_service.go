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
	req, err := http.NewRequest("POST", c.BaseURL+"/api/v1/services/databaseServices", strings.NewReader(string(postJSON)))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	body, statusCode, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create db_service, status code: %d, response: %s", statusCode, string(body))

	}
	create_db := CreateDB_res{}
	err = json.Unmarshal(body, &create_db)
	if err != nil {
		return nil, err
	}
	return &create_db, nil
}

func (c *Client) UpdateDBService(db UpdateDB_req, authToken *string) (*UpdateDB_res, error) {
	postJSON, err := json.Marshal(db)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal db data: %w", err)
	}
	req, err := http.NewRequest("PUT", c.BaseURL+"/api/v1/services/databaseServices", strings.NewReader(string(postJSON)))
	if err != nil {
		return nil, fmt.Errorf("failed to db request: %w", err)
	}
	body, statusCode, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to Update DBService, status code: %d, response: %s", statusCode, string(body))
	}
	updatedDBService := UpdateDB_res{}
	err = json.Unmarshal(body, &updatedDBService)
	if err != nil {
		return nil, err
	}
	return &updatedDBService, nil
}

func (c *Client) DeleteDBService(name string, authToken *string) error {
	req, err := http.NewRequest("DELETE", c.BaseURL+"/api/v1/services/databaseServices/name/"+name, nil)
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

	if statusCode != http.StatusOK {
		return fmt.Errorf("failed to deleted DBService, status code: %d, response: %s", statusCode, string(body))
	}
	return nil
}
