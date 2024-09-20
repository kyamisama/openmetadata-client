package openmetadata

// CreateUser call api request struct in OpenMetadata
type CreateUserReq struct {
	Description string   `json:"description"`
	DisplayName string   `json:"displayName"`
	Email       string   `json:"email"`
	Name        string   `json:"name"`
	Password    string   `json:"password"`
	Roles       []string `json:"roles"`
	Teams       []string `json:"teams"`
}

// CreateUser call api response struct in OpenMetadata
type CreateUserRes struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	FullyQualifiedName string    `json:"fullyQualifiedName"`
	Version            float64   `json:"version"`
	UpdatedAt          int       `json:"updatedAt"`
	UpdatedBy          string    `json:"updatedBy"`
	Email              string    `json:"email"`
	IsBot              bool      `json:"isBot"`
	IsAdmin            bool      `json:"isAdmin"`
	Teams              []TeamRes `json:"teams"`
	Roles              []RoleRes `json:"roles"`
	InheritedRoles     []RoleRes `json:"inheritedRole"`
	Deleted            bool      `json:"deleted"`
}

// GetUserの引数はIDを指定するため、リクエストボディ用の構造体は不要
type GetUserRes struct {
	ID                 string        `json:"id"`
	Name               string        `json:"name"`
	FullyQualifiedName string        `json:"fullyQualifiedName"`
	Description        string        `json:"description"`
	DisplayName        string        `json:"displayName"`
	Version            float64       `json:"version"`
	UpdatedAt          int64         `json:"updatedAt"`
	UpdatedBy          string        `json:"updatedBy"`
	Email              string        `json:"email"`
	Href               string        `json:"href"`
	IsBot              bool          `json:"isBot"`
	IsAdmin            bool          `json:"isAdmin"`
	Teams              []TeamRes     `json:"teams"`
	Deleted            bool          `json:"deleted"`
	Personas           []PersonasRes `json:"personas"`
	Roles              []RoleRes     `json:"roles"`
	InheritedRoles     []RoleRes     `json:"inheritedRoles"`
}
type PersonasRes struct {
	ID                 string `json:"id"`
	Type               string `json:"type"`
	Name               string `json:"name"`
	FullyQualifiedName string `json:"fullyQualifiedName"`
	Description        string `json:"description"`
	DisplayName        string `json:"displayName"`
}
type TeamRes struct {
	ID                 string `json:"id"`
	Type               string `json:"type"`
	Name               string `json:"name"`
	FullyQualifiedName string `json:"fullyQualifiedName"`
	Description        string `json:"description"`
	DisplayName        string `json:"displayName"`
}

// InheritedRolesもふくむ
type RoleRes struct {
	ID                 string `json:"id"`
	Type               string `json:"type"`
	Name               string `json:"name"`
	FullyQualifiedName string `json:"fullyQualifiedName"`
	Description        string `json:"description"`
	DisplayName        string `json:"displayName"`
	Deleted            bool   `json:"deleted"`
	Href               string `json:"href"`
}

type GetUsersRes struct {
	Data   []TeamRes `json:"data"`
	Paging Paging    `json:"paging"`
}
type Users struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
}
type UpdateUserReq struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
}

type UpdateUserRes struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
}
type PatchUserReq struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}
type PatchUserRes struct {
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

type CreateDBReq struct {
	Name        string `json:"name"`
	ServiceType string `json:"serviceType"`
	Description string `json:"description"`
	DisplayName string `json:"displayName"`
}

type CreateDBRes struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	FullyQualifiedName string `json:"fullyQualifiedName"`
	ServiceType        string `json:"serviceType"`
	Description        string `json:"description"`
	UpdatedBy          string `json:"updatedBy"`
	Deleted            bool   `json:"deleted"`
}

type UpdateDBReq struct {
	Name        string `json:"name"`
	ServiceType string `json:"serviceType"`
	Description string `json:"description"`
	DisplayName string `json:"displayName"`
}

type UpdateDBRes struct {
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
	ID          string    `json:"id"`
	Description string    `json:"description"`
	DisplayName string    `json:"displayName"`
	TeamType    string    `json:"teamType"`
	Name        string    `json:"name"`
	Policies    []TeamRes `json:"policies"`
}

type DeteleTeamRes struct {
	ID             string    `json:"id"`
	Description    string    `json:"description"`
	DisplayName    string    `json:"displayName"`
	TeamType       string    `json:"teamType"`
	Name           string    `json:"name"`
	InheritedRoles []RoleRes `json:"inheritedRole"`
	Policies       []TeamRes `json:"policies"`
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
	Data   []TeamRes `json:"data"`
	Paging Paging    `json:"paging"`
}

type PatchTeamReq struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}
type PatchTeamRes struct {
	ID             string    `json:"id"`
	Description    string    `json:"description"`
	DisplayName    string    `json:"displayName"`
	TeamType       string    `json:"teamType"`
	Name           string    `json:"name"`
	InheritedRoles []RoleRes `json:"inheritedRole"`
	Policies       []TeamRes `json:"policies"`
}
