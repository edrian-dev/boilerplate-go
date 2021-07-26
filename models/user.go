package models

// User ...
type User struct {
	Base
	Email    string `json:"email"`
	Password string `json:"-"`
}

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
