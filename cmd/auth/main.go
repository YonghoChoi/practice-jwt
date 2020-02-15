package main

import (
	"fmt"
	"github.com/YonghoChoi/practice-jwt/pkg/token"
	"os"
)

func main() {
	jwt, err := token.MakeJWT("yongho")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(jwt)
}
