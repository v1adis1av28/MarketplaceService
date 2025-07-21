package models

type User struct {
	Email    string `json : "email"`
	Password string `json : "password"`
	Role     string `json :"role"`
}

type AuthUserReq struct {
	Email    string `json : "email" example:"test1@example.com"`
	Password string `json : "password" example:"1234"`
}
