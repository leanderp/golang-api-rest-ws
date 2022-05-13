package models

import "github.com/golang-jwt/jwt/v4"

type AppClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}
