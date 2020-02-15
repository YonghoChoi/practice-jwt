package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func main() {
	claims := &Claims{
		Name: "yongho",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString("secret key")
	if err != nil {

	}

}
