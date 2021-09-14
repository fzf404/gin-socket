package model

import "github.com/golang-jwt/jwt"

// Claims jwt
type Claims struct {
	UserID uint
	jwt.StandardClaims
}
