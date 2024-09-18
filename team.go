package openmetadata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) CreateTeam(team CreateTeamReq, authToken *string) (*CreateTeamRes, error) {
	postJSON, err := json.Marshal(team)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal team data: %w", err)
	}
	req, err := c.newRequest("POST", c.BaseURL+"/api/v1/teams", authToken, strings.NewReader(string(postJSON)))
	if err != nil {
		return nil, fmt.Errorf("failed to marshal team data: %w", err)
	}

	var createTeamRes CreateTeamRes
	statusCode, err := c.doRequest(req, &createTeamRes)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create team, status code: %d", statusCode)
	}
	return &createTeamRes, nil
}

func (c *Client) DeleteTeam(id string, authToken *string) (*DeteleTeamRes, error) {
	req, err := c.newRequest("DELETE", c.BaseURL+"/api/v1/teams/"+id, authToken, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("hardDelete", "true")
	req.URL.RawQuery = q.Encode()
	var deleteTeam DeteleTeamRes
	statusCode, err := c.doRequest(req, &deleteTeam)
	if err != nil {
		return nil, err
	}
	// log.Println(statusCode)
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to deleted user, status code: %d", statusCode)
	}

	// If status code is 204, return success without decoding
	return &deleteTeam, nil
}

// GetTeam retrieves a user by ID from OpenMetadata
func (c *Client) GetTeam(id string, authToken *string) (*GetTeamRes, error) {
	req, err := c.newRequest("GET", c.BaseURL+"/api/v1/teams/"+id, authToken, nil)
	if err != nil {
		return nil, err
	}
	var getTeam GetTeamRes
	statusCode, err := c.doRequest(req, &getTeam)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get team, status code: %d", statusCode)
	}

	return &getTeam, nil
}

func (c *Client) GetTeams(authToken *string) (*GetTeamsRes, error) {
	req, err := c.newRequest("GET", c.BaseURL+"/api/v1/teams/", authToken, nil)
	if err != nil {
		return nil, err
	}
	var getTeamsRes GetTeamsRes
	statusCode, err := c.doRequest(req, &getTeamsRes)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get teams, status code: %d", statusCode)
	}

	return &getTeamsRes, nil
}

func (c *Client) PatchTeam(patchdata []PatchTeamReq, id string, authToken *string) (*PatchTeamRes, error) {
	postJSON, err := json.Marshal(patchdata)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user data: %w", err)
	}
	req, err := c.newRequest("PATCH", c.BaseURL+"/api/v1/teams/"+id, authToken, strings.NewReader(string(postJSON)))
	if err != nil {
		return nil, fmt.Errorf("failed to patch request: %w", err)
	}

	var patchTeamRes PatchTeamRes
	statusCode, err := c.doRequest(req, &patchTeamRes)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to patch user, status code: %d", statusCode)
	}

	return &patchTeamRes, nil
}
