package openmetadata

import (
	"fmt"
	"log"
	"testing"
)

// setup function to initialize the client
func setup() *Client {
	baseURL := "http://192.168.0.19:8585"                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             // OpenMetadataのAPIのベースURL
	authToken := "eyJraWQiOiJHYjM4OWEtOWY3Ni1nZGpzLWE5MmotMDI0MmJrOTQzNTYiLCJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJvcGVuLW1ldGFkYXRhLm9yZyIsInN1YiI6ImMzNzNpc20iLCJyb2xlcyI6WyJBZG1pbiJdLCJlbWFpbCI6ImMzNzNpc21AZ21haWwuY29tIiwiaXNCb3QiOnRydWUsInRva2VuVHlwZSI6IkJPVCIsImlhdCI6MTcyMzUyNDM3MiwiZXhwIjoxNzMxMzAwMzcyfQ.nXk9xbyYUTMPiKA47o1ddS_XSuFLXw13OFnPLKQCxbOgMHQflmVz5inT3hpu5IbTOHDPlPP_UOz8-voWAqHwA9epJnMQd3EXawGANvdscXXDggyIIjsaFikQVvFWEaC1qTHjOBZVdrHSUOXab5fZ_2pKkJmFNDoP3ullzSKgNhrW8hKiTJw0ArpUKxVMCupP4BqtxEqDQYqJL5buvvXERx3FmX4euhuUOiwjqrLucxWFHuhEhxe6A1c1CAHYQkkgzD9__6jcJopv_sgexCOO-MizSSU_yCkUzc9wodp3iS1VI5i8LNOM1fh8m6qUT8F5sGIf83GdDQxaBMSAhhSI4Q" // 有効な認証トークン
	client := NewClient(baseURL, authToken)
	return client
}

func TestCreateUsers(t *testing.T) {
	client := setup()

	newUser := CreateUser_req{
		Name:        "john.doe",
		Email:       "john.doe@example.com",
		DisplayName: "j.d",
		Description: "hogehoge",
		Password:    "P@ssW0rd",
		Roles:       []string{"fa8521d3-d523-4d7f-8935-f7b8379aba2d"},
		Teams:       []string{"fbbffde9-137b-4564-987d-367bcaed1ac4"},
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

	name := "john.doe" // 事前に作成したユーザーのNameを指定
	user, err := client.GetUser(name, &client.AuthToken)
	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
	}

	if user.Name != name {
		t.Errorf("GetUser returned unexpected ID: got %v, want %v", user.Name, name)
	}
}

func TestGetUsers(t *testing.T) {
	client := setup()

	_, err := client.GetUsers(&client.AuthToken)
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

	res, err := client.UpdateUser(updateData, &client.AuthToken)
	if err != nil {
		t.Fatalf("UpdateUser failed: %v", err)
	}
	log.Printf("Response Description: %s", res.Description)
	log.Printf("Update Data Description: %s", updateData.Description)
	if res.Description == updateData.Description {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}

}

func TestDeleteUser(t *testing.T) {
	client := setup()

	name := "john.doe" // 削除対象のユーザーName
	err := client.DeleteUser(name, &client.AuthToken)
	if err != nil {
		t.Fatalf("DeleteUser failed: %v", err)
	}

}
