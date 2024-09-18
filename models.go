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
	Teams              []Team_role     `json:"teams"`
	Roles              []Team_role     `json:"roles"`
	InheritedRoles     []InheritedRole `json:"inheritedRole"`
	Deleted            bool            `json:"deleted"`
}

type Team_role struct {
	ID                 string `json:"id"`
	Type               string `json:"type"`
	Name               string `json:"name"`
	FullyQualifiedName string `json:"fullyQualifiedName"`
	DisplayName        string `json:"displayName"`
	Deleted            bool   `json:"deleted"`
	Href               string `json:"href"`
}

type InheritedRole struct {
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

type GetUsersRes struct {
	Data   []Team `json:"data"`
	Paging Paging `json:"paging"`
}
type Users struct {
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
}

type UpdateUser_res struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
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

type DeleteDBRes struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	FullyQualifiedName string `json:"fullyQualifiedName"`
	ServiceType        string `json:"serviceType"`
	Description        string `json:"description"`
	UpdatedBy          string `json:"updatedBy"`
	Deleted            bool   `json:"deleted"`
}

type CreateTeamReq struct {
	Description string   `json:"description"`
	DisplayName string   `json:"displayName"`
	TeamType    string   `json:"teamType"`
	Name        string   `json:"name"`
	Policies    []string `json:"policies"`
}
type CreateTeamRes struct {
	ID          string      `json:"id"`
	Description string      `json:"description"`
	DisplayName string      `json:"displayName"`
	TeamType    string      `json:"teamType"`
	Name        string      `json:"name"`
	Policies    []Team_role `json:"policies"`
}

type DeteleTeamRes struct {
	ID             string          `json:"id"`
	Description    string          `json:"description"`
	DisplayName    string          `json:"displayName"`
	TeamType       string          `json:"teamType"`
	Name           string          `json:"name"`
	InheritedRoles []InheritedRole `json:"inheritedRole"`
	Policies       []Team_role     `json:"policies"`
}

type GetTeamRes struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	DisplayName string `json:"displayName"`
	TeamType    string `json:"teamType"`
	Name        string `json:"name"`
}

type Paging struct {
	Total int `json:"total"`
}

type GetTeamsRes struct {
	Data   []Team `json:"data"`
	Paging Paging `json:"paging"`
}
type Team struct {
	ID             string          `json:"id"`
	Description    string          `json:"description"`
	DisplayName    string          `json:"displayName"`
	TeamType       string          `json:"teamType"`
	Name           string          `json:"name"`
	InheritedRoles []InheritedRole `json:"inheritedRole"`
	Policies       []Team_role     `json:"policies"`
}
type PatchTeamReq struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}
type PatchTeamRes struct {
	ID             string          `json:"id"`
	Description    string          `json:"description"`
	DisplayName    string          `json:"displayName"`
	TeamType       string          `json:"teamType"`
	Name           string          `json:"name"`
	InheritedRoles []InheritedRole `json:"inheritedRole"`
	Policies       []Team_role     `json:"policies"`
}
