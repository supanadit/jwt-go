package main

import (
	"fmt"
	"github.com/supanadit/easy-jwt-go"
	"log"
)

func main() {
	// Set Your JWT Secret Code, its optional but important, because default secret code is very insecure
	ej.SetJWTSecretCode("Your Secret Code")

	// Create default authorization
	auth := ej.Authorization{
		Username: "admin",
		Password: "admin",
	}

	// Generate JWT Token from default authorization model
	token, err := auth.GenerateJWT()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("JWT Token : " + token)

	// Verify the token
	valid, err := auth.VerifyJWT(token)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Status : ")

	if valid {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}
