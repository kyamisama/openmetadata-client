package openmetadata

import (
	"testing"
)

// setup function to initialize the client
func setup() *Client {
	baseURL := "http://192.168.0.19:8585/api/v1/users" // OpenMetadataのAPIのベースURL
	authToken := ""                                    // 有効な認証トークン
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
