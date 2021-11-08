package model

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Mobile string
	Name   string
	jwt.StandardClaims
}
