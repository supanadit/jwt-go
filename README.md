# Easy JWT Setup for Golang

A JWT library for Golang, this is not original library meaning this library depends on other library such as [jwt-go](https://github.com/dgrijalva/jwt-go) which is the original JWT library for golang, [mapstructure](github.com/mitchellh/mapstructure), and [bcrypt](golang.org/x/crypto/bcrypt).

## Installation

`go get -u github.com/supanadit/easy-jwt-go`

## Quick Start

```go
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
```

## Create JWT Authorization from your own struct ?
```go
package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/supanadit/easy-jwt-go"
	"log"
)

type Login struct {
	Email    string
	Password string
	jwt.Claims
}

func main() {
	// Set Your JWT Secret Code, its optional but important, because default secret code is very insecure
	ej.SetJWTSecretCode("Your Secret Code")

	// Create authorization from your own struct
	auth := Login{
		Email:    "hello@email.com",
		Password: "123",
	}

	// Generate JWT Token
	token, err := ej.GenerateJWT(auth)
	if err != nil {
		log.Fatal(err)
	}

	// Verify token
	valid, err := ej.VerifyJWT(auth, token)
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
```