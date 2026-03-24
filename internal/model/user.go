package model

type UserRegister struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserInfo struct {
	Email string `json:"email"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

var (
	UserRoleAdmin = "admin"
	UserRoleUser  = "user"
)
