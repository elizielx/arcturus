package utils

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
