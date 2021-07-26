package models

import "github.com/dgrijalva/jwt-go"

// CustomClaims ...
type CustomClaims struct {
	jwt.StandardClaims `json:"-"`
	UserID             uint   `json:"-"`
	Token              string `json:"token"`
}
