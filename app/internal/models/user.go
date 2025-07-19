package models

type User struct {
	Email    string `json : "email"`
	Password string `json : "password"`
	Role     string `json :"role"`
}

type RegistrationUserReq struct {
	Email    string `json : "email"`
	Password string `json : "password"`
}
