// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Account  string `form:"account"`
	Password string `form:"password"`
}

type LoginUserInfo struct {
	UserId    int    `json:"userId"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Gender    int8   `json:"gender"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}

type LoginResponse struct {
	LoginUserInfo LoginUserInfo `json:"loginUserInfo"`
	AccessToken   string        `json:"accessToken"`
	AccessExpire  int64         `json:"accessExpire"`
	RefreshTime   int64         `json:"refreshTime"`
}