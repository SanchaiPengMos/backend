package models

import 	jwt "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}
