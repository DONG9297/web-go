package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"web-go/src/model"
)

var jwtSecret = []byte("DengHanzi") // jwt密钥

// GenerateToken 生成 token
func GenerateToken(mobile, name string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(300 * time.Second)
	issuer := "frank"
	claims := model.Claims{
		Mobile: mobile,
		Name:   name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 解析 token
func ParseToken(token string) (*model.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*model.Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
