package config

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func GetTokenAuth() *jwtauth.JWTAuth {
	//TODO: Get secret and public key from ssh
	tokenAuth = jwtauth.New("HS256", []byte("secretkey"), "verifyKey")
	return tokenAuth
}
