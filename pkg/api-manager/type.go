package api_manager

/**
http 的方法
*/
type HttpMethod string

const (
	MethodGet    HttpMethod = "get"
	MethodPost   HttpMethod = "post"
	MethodPut    HttpMethod = "put"
	MethodDelete HttpMethod = "delete"
)

/**
项目中的角色
*/
type Role string

const (
	RoleAdministrator Role = "administrator"
	RoleDeveloper     Role = "developer"
)

/**
API 实例
*/
type Api struct {
	Code           string     `json:"code"`
	Name           string     `json:"name"`
	ProjectCode    string     `json:"projectCode"`
	CollectionCode string     `json:"collectionCode"`
	Method         HttpMethod `json:"method"`
	Path           string     `json:"path"`
}

/**
API 实例
*/
type ApiCollection struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	ProjectCode string `json:"projectCode"`
}

type User struct {
	UserName string
}
