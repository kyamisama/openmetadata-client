package openmetadata

import (
	"fmt"
	"testing"
)

// // setup function to initialize the client
func setup() *Client {
	baseURL := "http://192.168.0.19:8585"                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             // OpenMetadataのAPIのベースURL
	authToken := "eyJraWQiOiJHYjM4OWEtOWY3Ni1nZGpzLWE5MmotMDI0MmJrOTQzNTYiLCJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJvcGVuLW1ldGFkYXRhLm9yZyIsInN1YiI6ImMzNzNpc20iLCJyb2xlcyI6WyJBZG1pbiJdLCJlbWFpbCI6ImMzNzNpc21AZ21haWwuY29tIiwiaXNCb3QiOnRydWUsInRva2VuVHlwZSI6IkJPVCIsImlhdCI6MTcyMzUyNDM3MiwiZXhwIjoxNzMxMzAwMzcyfQ.nXk9xbyYUTMPiKA47o1ddS_XSuFLXw13OFnPLKQCxbOgMHQflmVz5inT3hpu5IbTOHDPlPP_UOz8-voWAqHwA9epJnMQd3EXawGANvdscXXDggyIIjsaFikQVvFWEaC1qTHjOBZVdrHSUOXab5fZ_2pKkJmFNDoP3ullzSKgNhrW8hKiTJw0ArpUKxVMCupP4BqtxEqDQYqJL5buvvXERx3FmX4euhuUOiwjqrLucxWFHuhEhxe6A1c1CAHYQkkgzD9__6jcJopv_sgexCOO-MizSSU_yCkUzc9wodp3iS1VI5i8LNOM1fh8m6qUT8F5sGIf83GdDQxaBMSAhhSI4Q" // 有効な認証トークン
	client, err := NewClient(&baseURL, &authToken)
	if err != nil {
		fmt.Printf("setup failed clent: %v", err)
	}
	return client
}

func TestCreateDBService(t *testing.T) {
	client := setup()

	createdDB := CreateDB_req{
		Name:        "Snowflake_DB",
		ServiceType: "Snowflake",
		// Description: "snowflake",
		DisplayName: "Snowflake_DB",
	}
	createdDBService, err := client.CreateDBService(createdDB, &client.AuthToken)
	if err != nil {
		t.Fatalf("CreatedDBService failed: %v", err)
	}

	if createdDBService.Name != createdDB.Name || createdDBService.ServiceType != createdDB.ServiceType {
		t.Errorf("CreateDBService returned unexpected DB data: got %v, want %v", createdDBService, createdDB)
	}
}

func TestUpdateDBService(t *testing.T) {
	client := setup()

	UpdateDB := UpdateDB_req{
		Name:        "Snowflake_DB",
		ServiceType: "Snowflake",
		Description: "snowflake2024",
		DisplayName: "Snowflake_DB2024",
	}
	updatedDBService, err := client.UpdateDBService(UpdateDB, &client.AuthToken)
	if err != nil {
		return
	}

	if updatedDBService.Description != UpdateDB.Description {
		t.Errorf("UpdateDBService returned unexpected DB data: got %v, want %v", updatedDBService, UpdateDB)
	}
}

func TestDeleteDBService(t *testing.T) {
	client := setup()

	DBDeleteName := "Snowflake_DB"
	_, err := client.DeleteDBService(DBDeleteName, &client.AuthToken)
	if err != nil {
		t.Fatalf("DeleteDBService failed: %v", err)
	}
}

func TestCreateUsers(t *testing.T) {
	client := setup()

	newUser := CreateUser_req{
		Name:        "john.doe",
		Email:       "john.doe@example.com",
		DisplayName: "j.d",
		//Description: "hogehoge",
		Password: "P@ssW0rd",
		Roles:    []string{"fa8521d3-d523-4d7f-8935-f7b8379aba2d"},
		Teams:    []string{"3f9ddb39-84b2-40cd-a5a2-d2e50ca1f478"},
	}

	createdUser, err := client.CreateUser(newUser, &client.AuthToken)
	if err != nil {
		t.Fatalf("CreatedUsers failed: %v", err)
	}
	if createdUser.Name != newUser.Name || createdUser.Email != newUser.Email {
		t.Errorf("CreateUser returned unexpected user data: got %v, want %v", createdUser, newUser)
	}
}

func TestGetUser(t *testing.T) {
	client := setup()

	id := "f40054a6-dcde-4b87-a318-03bcab048cf0" // 事前に作成したユーザーのNameを指定
	user, err := client.GetUser(id, &client.AuthToken)
	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
	}

	if user.ID != id {
		t.Errorf("GetUser returned unexpected ID: got %v, want %v", user.Name, id)
	}
}

func TestGetUsers(t *testing.T) {
	client := setup()

	clientRes, err := client.GetUsers(&client.AuthToken)
	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
	}
	for _, team := range clientRes.Data {
		if team.Name == "john.doe" {
			fmt.Println("ok")
		}
	}
}

func TestUpdateUser(t *testing.T) {
	client := setup()

	updateData := UpdateUser_req{
		Name:        "john.doe",
		Email:       "john.doe@example.com",
		DisplayName: "j.d2222",
		Description: "hogehoge2222",
	}

	res, err := client.UpdateUser(updateData, &client.AuthToken)
	if err != nil {
		t.Fatalf("UpdateUser failed: %v", err)
	}

	if res.Description == updateData.Description {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}

}

func TestPatchUser(t *testing.T) {
	client := setup()

	patchData := []PatchUser_req{
		{
			Op:    "replace",
			Path:  "/isAdmin",
			Value: true,
		},
		{
			Op:    "replace",
			Path:  "/description",
			Value: "hogehogepatch",
		},
		{
			Op:    "add",
			Path:  "/displayName",
			Value: "hogehoge",
		},
		{
			Op:    "remove",
			Path:  "/displayName",
			Value: "",
		},
	}
	// log.Println(patchData)
	id := "f40054a6-dcde-4b87-a318-03bcab048cf0"
	res, err := client.PatchUser(patchData, id, &client.AuthToken)
	if err != nil {
		t.Fatalf("PatchUser failed: %v", err)
	}

	if res.IsAdmin != true {
		t.Errorf("Expected IsAdmin to be true, got %v", res.IsAdmin)
	}
	// log.Println(res.DisplayName)
	if res.Description != "hogehogepatch" {
		t.Errorf("Expected Description to be 'hogehogepatch', got %v", res.Description)
	}
}

func TestDeleteUser(t *testing.T) {
	client := setup()

	name := "john.doe" // 削除対象のユーザーName
	_, err := client.DeleteUser(name, &client.AuthToken)
	if err != nil {
		t.Fatalf("DeleteUser failed: %v", err)
	}

}

func TestCreateTeam(t *testing.T) {
	client := setup()
	createTeam := CreateTeamReq{
		Name:        "testTeam",
		TeamType:    "Group",
		Description: "testDesc",
		DisplayName: "testDisp",
		Policies:    []string{"e481b473-7043-4561-83b4-ef7ae372f80a"},
	}
	createdTeam, err := client.CreateTeam(createTeam, &client.AuthToken)
	if err != nil {
		t.Fatalf("CreatedTeam failed: %v", err)
	}

	if createdTeam.Name != createTeam.Name || createdTeam.TeamType != createTeam.TeamType {
		t.Errorf("CreateUser returned unexpected user data: got %v, want %v", createdTeam, createTeam)
	}
}

func TestGetTeams(t *testing.T) {
	client := setup()
	clientRes, err := client.GetTeams(&client.AuthToken)
	if err != nil {
		t.Fatalf("GetTeams failed: %v", err)
	}

	for _, team := range clientRes.Data {
		if team.Name == "testTeam" {
			fmt.Println("ok")
		}
	}
}

func TestDeleteTeam(t *testing.T) {
	client := setup()
	clientRes, err := client.GetTeams(&client.AuthToken)
	if err != nil {
		t.Fatalf("GetTeams failed: %v", err)
	}

	for _, team := range clientRes.Data {
		if team.Name == "testTeam" {
			_, err := client.DeleteTeam(team.ID, &client.AuthToken)
			if err != nil {
				t.Fatalf("DeleteTeam failed: %v", err)
			}
		}
	}
}
func TestPatchTeam(t *testing.T) {
	client := setup()

	patchData := []PatchTeamReq{
		{
			Op:    "replace",
			Path:  "/description",
			Value: "hogehogepatch",
		},
		{
			Op:    "add",
			Path:  "/displayName",
			Value: "hogehoge",
		},
		{
			Op:    "remove",
			Path:  "/displayName",
			Value: "",
		},
	}
	clientRes, err := client.GetTeams(&client.AuthToken)
	if err != nil {
		t.Fatalf("GetTeams failed: %v", err)
	}
	for _, team := range clientRes.Data {
		if team.Name == "testTeam" {
			res, err := client.PatchTeam(patchData, team.ID, &client.AuthToken)
			if err != nil {
				t.Fatalf("PatchUser failed: %v", err)
			}
			if res.Description != "hogehogepatch" {
				t.Errorf("Expected Description to be 'hogehogepatch', got %v", res.Description)
			}
		}
	}
}
