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
	createdDBService, err := client.CreateDBService(createdDB, client.AuthToken)
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
	updatedDBService, err := client.UpdateDBService(UpdateDB, client.AuthToken)
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
	_, err := client.DeleteDBService(DBDeleteName, client.AuthToken)
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

	createdUser, err := client.CreateUser(newUser, client.AuthToken)
	if err != nil {
		t.Fatalf("CreatedUsers failed: %v", err)
	}
	if createdUser.Name != newUser.Name || createdUser.Email != newUser.Email {
		t.Errorf("CreateUser returned unexpected user data: got %v, want %v", createdUser, newUser)
	}
}

func TestGetUser(t *testing.T) {
	client := setup()

	id := "35a43337-9b25-41e0-a7d0-0177fd2a8214" // 事前に作成したユーザーのNameを指定
	user, err := client.GetUser(id, client.AuthToken)
	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
	}

	if user.ID != id {
		t.Errorf("GetUser returned unexpected ID: got %v, want %v", user.Name, id)
	}
}

func TestGetUsers(t *testing.T) {
	client := setup()

	_, err := client.GetUsers(client.AuthToken)
	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
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

	res, err := client.UpdateUser(updateData, client.AuthToken)
	if err != nil {
		t.Fatalf("UpdateUser failed: %v", err)
	}

	if res.Description == updateData.Description {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}

}

func TestDeleteUser(t *testing.T) {
	client := setup()

	name := "john.doe" // 削除対象のユーザーName
	_, err := client.DeleteUser(name, client.AuthToken)
	if err != nil {
		t.Fatalf("DeleteUser failed: %v", err)
	}

}
