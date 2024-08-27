package openmetadata

// User represents the structure of a user in OpenMetadata
type CreateUser struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	Password    string `json:"password"`
}

type GetUser struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
}

type UpdateUser struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	Password    string `json:"password"`
}

type DeleteUser struct {
	Name string `json:"name"`
}
