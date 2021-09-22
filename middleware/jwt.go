/*
 * @Author: fzf404
 * @Date: 2021-09-22 19:52:23
 * @LastEditTime: 2021-09-22 19:54:20
 * @Description: jwt认证
 */
package middleware

import (
	"gin-socket/model"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("gin-socket-secret")

/**
 * @description: 生成Token
 * @param {model.User} user
 */
func ReleaseToken(user model.User) (string, error) {

	// 超时时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	// token结构生成
	claims := &model.Claims{
		// 使用ID作为	有效载荷
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "fzf404",
			Subject:   "token",
		},
	}

	// 将Claims加密存储为Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 加入key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

/**
 * @description: 验证Token
 * @param {string} tokenString
 */
func ParseToken(tokenString string) (*jwt.Token, *model.Claims, error) {
	claims := &model.Claims{}
	// 解码
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
