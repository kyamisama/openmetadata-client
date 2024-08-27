package openmetadata

import (
	"testing"
)

// setup function to initialize the client
func setup() *Client {
	baseURL := "http://192.168.0.19:8585/api/v1/users"                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                // OpenMetadataのAPIのベースURL
	authToken := "eyJraWQiOiJHYjM4OWEtOWY3Ni1nZGpzLWE5MmotMDI0MmJrOTQzNTYiLCJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJvcGVuLW1ldGFkYXRhLm9yZyIsInN1YiI6ImMzNzNpc20iLCJyb2xlcyI6WyJBZG1pbiJdLCJlbWFpbCI6ImMzNzNpc21AZ21haWwuY29tIiwiaXNCb3QiOnRydWUsInRva2VuVHlwZSI6IkJPVCIsImlhdCI6MTcyMzUyNDM3MiwiZXhwIjoxNzMxMzAwMzcyfQ.nXk9xbyYUTMPiKA47o1ddS_XSuFLXw13OFnPLKQCxbOgMHQflmVz5inT3hpu5IbTOHDPlPP_UOz8-voWAqHwA9epJnMQd3EXawGANvdscXXDggyIIjsaFikQVvFWEaC1qTHjOBZVdrHSUOXab5fZ_2pKkJmFNDoP3ullzSKgNhrW8hKiTJw0ArpUKxVMCupP4BqtxEqDQYqJL5buvvXERx3FmX4euhuUOiwjqrLucxWFHuhEhxe6A1c1CAHYQkkgzD9__6jcJopv_sgexCOO-MizSSU_yCkUzc9wodp3iS1VI5i8LNOM1fh8m6qUT8F5sGIf83GdDQxaBMSAhhSI4Q" // 有効な認証トークン
	client := NewClient(baseURL, authToken)
	return client
}

func TestCreateUser(t *testing.T) {
	client := setup()

	newUser := &CreateUser{
		Name:        "john.doe",
		Email:       "john.doe@example.com",
		DisplayName: "j.d",
		Description: "hogehoge",
		Password:    "P@ssW0rd",
	}

	createdUser, err := client.CreateUser(newUser)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	if createdUser.Name != newUser.Name || createdUser.Email != newUser.Email {
		t.Errorf("CreateUser returned unexpected user data: got %v, want %v", createdUser, newUser)
	}
}

func TestGetUser(t *testing.T) {
	client := setup()

	name := "john.doe" // 事前に作成したユーザーのNameを指定
	user, err := client.GetUser(name)
	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
	}

	if user.Name != name {
		t.Errorf("GetUser returned unexpected ID: got %v, want %v", user.Name, name)
	}
}

func TestGetUsers(t *testing.T) {
	client := setup()

	users, err := client.GetUsers()
	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
	}

	// Check if users list is not empty
	if len(users) == 0 {
		t.Errorf("Expected at least one user, got 0")
	}
}

func TestUpdateUser(t *testing.T) {
	client := setup()

	updateData := &UpdateUser{
		Name:        "john.doe",
		Email:       "john.doe@example.com",
		DisplayName: "j.d2222",
		Description: "hogehoge2222",
		Password:    "P@ssW0rd",
	}

	updatedUser, err := client.UpdateUser(updateData)
	if err != nil {
		t.Fatalf("UpdateUser failed: %v", err)
	}

	if updatedUser.Name != updateData.Name || updatedUser.Email != updateData.Email {
		t.Errorf("UpdateUser returned unexpected user data: got %v, want %v", updatedUser, updateData)
	}
}

func TestDeleteUser(t *testing.T) {
	client := setup()

	name := "john.doe" // 削除対象のユーザーName
	_, err := client.DeleteUser(name)
	if err != nil {
		t.Fatalf("DeleteUser failed: %v", err)
	}

	// もしユーザーを再度取得し、その存在を確認するテストを行う場合
	_, err = client.GetUser(name)
	if err == nil {
		t.Errorf("Expected error when getting a deleted user, but got none")
	}
}
