package aut


import "github.com/dgrijalva/jwt-go"

type Login struct {
	Id       uint   `sql:"primary_key" json:"id" form:"id" query:"id"`
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

type JwtClaims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}
