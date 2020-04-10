# Easy JWT Setup for Golang

A JWT library for Golang, this is not original library meaning this library depends on other library such as [jwt-go](https://github.com/dgrijalva/jwt-go) which is the original JWT library for golang, [mapstructure](github.com/mitchellh/mapstructure), and [bcrypt](golang.org/x/crypto/bcrypt).

## Installation

`go get github.com/supanadit/easy-jwt-go`

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

## Custom Authorization

```go
package main

import (
	"fmt"
	"github.com/supanadit/easy-jwt-go"
	"log"
)

type Login struct {
	Email    string
	Password string
	Name     string
}

func main() {
	// Set Your JWT Secret Code, its optional but important, because default secret code is very insecure
	ej.SetJWTSecretCode("Your Secret Code")

	// Create default authorization
	auth := Login{
		Email:    "asd@asd.com",
		Password: "asd",
		Name:     "asd",
	}

	// Generate JWT Token from default authorization model
	token, err := ej.GenerateJWT(auth)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("JWT Token : " + token)

	// Variable for decoded JWT token
	var dataAuth Login
	// Verify the token
	valid, err := ej.VerifyAndBindingJWT(&dataAuth, token)
	if err != nil {
		fmt.Println(err)
	}

	// or simply you can do this, if you don't need to decode the JWT
	// valid, err := ej.VerifyJWT(token)
	// if err != nil {
	//	 fmt.Println(err)
	// }

	fmt.Print("Status : ")

	if valid {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}
```

## Encrypt & Verify Password

```go
package main

import (
	"fmt"
	"github.com/supanadit/easy-jwt-go"
	"log"
)

type Login struct {
	Email    string
	Password string
}

func main() {
	// Set Your JWT Secret Code, its optional but important, because default secret code is very insecure
	ej.SetJWTSecretCode("Your Secret Code")

	// Create authorization from your own struct
	auth := Login{
		Email:    "hello@email.com",
		Password: "123",
	}

	// Encrypt password, which you can save to database
	ep, err := ej.EncryptPassword(auth.Password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encrypted Password " + string(ep))

	// Verify Encrypted Password
	valid, err := ej.VerifyPassword(string(ep), auth.Password)
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

## Decrypt Password

No you can't, as the topic at [Stack Exchange](https://security.stackexchange.com/questions/193943/is-it-possible-to-decrypt-bcrypt-encryption)

> bcrypt is not an encryption function, it's a password hashing function, relying on Blowfish's key scheduling, not its encryption. Hashing are mathematical one-way functions, meaning there is no way to reverse the output string to get the input string.
  <br/> of course only Siths deal in absolutes and there are a few attacks against hashes. But none of them are "reversing" the hashing, AFAIK.

so that enough to secure the password

## Set Expired Time

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
    // You can simply do this, ej.setExpiredTime(Hour,Minute,Second)
	ej.SetExpiredTime(0, 0, 1)
}
```