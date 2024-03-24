package jwtM

import "github.com/golang-jwt/jwt/v4"

type Privileges struct {
	Admin bool `json:"admin"`
	User  bool `json:"user"`
	Coach bool `json:"coach"`
}

type CustomClaims struct {
	jwt.RegisteredClaims
	Username   string     `json:"username"`
	Privileges Privileges `json:"privileges"`
}
