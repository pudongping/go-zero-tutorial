// Code generated by goctl. DO NOT EDIT.
package types

type Pagination struct {
	Total    int64 `json:"total"`    // 数据总条数
	Page     int64 `json:"page"`     // 当前请求页码
	PerPage  int64 `json:"perPage"`  // 每页显示数
	LastPage int64 `json:"lastPage"` // 最大页码（最后一页）
}

type UserInfo struct {
	UserId    int    `json:"userId"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Gender    int8   `json:"gender"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}

type RegisterRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshTime  int64  `json:"refreshTime"`
}

type LoginRequest struct {
	Account  string `form:"account"`
	Password string `form:"password"`
}

type LoginResponse struct {
	UserInfo     UserInfo `json:"userInfo"`
	AccessToken  string   `json:"accessToken"`
	AccessExpire int64    `json:"accessExpire"`
	RefreshTime  int64    `json:"refreshTime"`
}

type AuthMeRequest struct {
}

type AuthMeResponse struct {
	UserInfo UserInfo `json:"userInfo"`
}

type UserListRequest struct {
	Keyword string `form:"keyword,optional"`
	Page    int64  `form:"page,optional,default=1"`
	PerPage int64  `form:"perPage,optional,default=10,range=[1:200]"`
	OrderBy string `form:"orderBy,optional"`
}

type UserListResponse struct {
	Result     []UserInfo `json:"results"`
	Pagination Pagination `json:"pagination"`
}
