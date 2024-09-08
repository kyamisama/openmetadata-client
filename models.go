package openmetadata

// CreateUser call api request struct in OpenMetadata
type CreateUser_req struct {
	Description string   `json:"description"`
	DisplayName string   `json:"displayName"`
	Email       string   `json:"email"`
	Name        string   `json:"name"`
	Password    string   `json:"password"`
	Roles       []string `json:"roles"`
	Teams       []string `json:"teams"`
}

// CreateUser call api response struct in OpenMetadata
type CreateUser_res struct {
	ID                 string          `json:"id"`
	Name               string          `json:"name"`
	FullyQualifiedName string          `json:"fullyQualifiedName"`
	Version            float64         `json:"version"`
	UpdatedAt          int             `json:"updatedAt"`
	UpdatedBy          string          `json:"updatedBy"`
	Email              string          `json:"email"`
	IsBot              bool            `json:"isBot"`
	IsAdmin            bool            `json:"isAdmin"`
	Teams              []team_role     `json:"teams"`
	Roles              []team_role     `json:"roles"`
	InheritedRoles     []inheritedRole `json:"inheritedRole"`
	Deleted            bool            `json:"deleted"`
}

type team_role struct {
	ID                 string `json:"id"`
	Type               string `json:"type"`
	Name               string `json:"name"`
	FullyQualifiedName string `json:"fullyQualifiedName"`
	DisplayName        string `json:"displayName"`
	Deleted            bool   `json:"deleted"`
	Href               string `json:"href"`
}

type inheritedRole struct {
	ID                 string `json:"id"`
	Type               string `json:"type"`
	Name               string `json:"name"`
	FullyQualifiedName string `json:"fullyQualifiedName"`
	Description        string `json:"description"`
	DisplayName        string `json:"displayName"`
	Deleted            bool   `json:"deleted"`
	Href               string `json:"href"`
}

type GetUser_req struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
}

type GetUser_res struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
}

type GetUsers_res struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
}
type UpdateUser_req struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	Password    string `json:"password"`
}

type UpdateUser_res struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	Password    string `json:"password"`
}
type PatchUser_req struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}
type PatchUser_res struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	IsAdmin     bool   `json:"isAdmin"`
}
type DeleteUser struct {
	Name string `json:"name"`
}

type CreateDB_req struct {
	Name        string `json:"name"`
	ServiceType string `json:"serviceType"`
	Description string `json:"description"`
	DisplayName string `json:"displayName"`
}

type CreateDB_res struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	FullyQualifiedName string `json:"fullyQualifiedName"`
	ServiceType        string `json:"serviceType"`
	Description        string `json:"description"`
	UpdatedBy          string `json:"updatedBy"`
	Deleted            bool   `json:"deleted"`
}

type UpdateDB_req struct {
	Name        string `json:"name"`
	ServiceType string `json:"serviceType"`
	Description string `json:"description"`
	DisplayName string `json:"displayName"`
}

type UpdateDB_res struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	FullyQualifiedName string `json:"fullyQualifiedName"`
	ServiceType        string `json:"serviceType"`
	Description        string `json:"description"`
	UpdatedBy          string `json:"updatedBy"`
	Deleted            bool   `json:"deleted"`
}

type DeleteDB struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	FullyQualifiedName string `json:"fullyQualifiedName"`
	ServiceType        string `json:"serviceType"`
	Description        string `json:"description"`
	UpdatedBy          string `json:"updatedBy"`
	Deleted            bool   `json:"deleted"`
}
