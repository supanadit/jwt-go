# JWT Go

[![Build Status](https://travis-ci.org/supanadit/jwt-go.svg?branch=master)](https://travis-ci.org/supanadit/jwt-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/supanadit/jwt-go)](https://goreportcard.com/report/github.com/supanadit/jwt-go)

A JWT library for Golang and easy to use, but this is not original library meaning this library depends on other library such as [jwt-go](https://github.com/dgrijalva/jwt-go) which is the original JWT library for golang, also [mapstructure](github.com/mitchellh/mapstructure), and [bcrypt](golang.org/x/crypto/bcrypt).

## Installation

`go get github.com/supanadit/jwt-go`

## Quick Start

```go
package main

import (
	"fmt"
	"github.com/supanadit/jwt-go"
	"log"
)

func main() {
	// Set Your JWT Secret Code, its optional but important, because default secret code is very insecure
	jwt.SetJWTSecretCode("Your Secret Code")

	// Create default authorization
	auth := jwt.Authorization{
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

<details><summary><b>Custom Authorization</b></summary>

<p>

```go
package main

import (
	"fmt"
	"github.com/supanadit/jwt-go"
	"log"
)

type Login struct {
	Email    string
	Password string
	Name     string
}

func main() {
	// Set Your JWT Secret Code, its optional but important, because default secret code is very insecure
	jwt.SetJWTSecretCode("Your Secret Code")

	// Create default authorization
	auth := Login{
		Email:    "asd@asd.com",
		Password: "asd",
		Name:     "asd",
	}

	// Generate JWT Token from default authorization model
	token, err := jwt.GenerateJWT(auth)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("JWT Token : " + token)

	// Variable for decoded JWT token
	var dataAuth Login
	// Verify the token
	valid, err := jwt.VerifyAndBindingJWT(&dataAuth, token)
	if err != nil {
		fmt.Println(err)
	}

	// or simply you can do this, if you don't need to decode the JWT
	// valid, err := jwt.VerifyJWT(token)
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

</p>
</details>

<details><summary><b>Encrypt & Verify Password</b></summary>

<p>

```go
package main

import (
	"fmt"
	"github.com/supanadit/jwt-go"
	"log"
)

type Login struct {
	Email    string
	Password string
}

func main() {
	// Set Your JWT Secret Code, its optional but important, because default secret code is very insecure
	jwt.SetJWTSecretCode("Your Secret Code")

	// Create authorization from your own struct
	auth := Login{
		Email:    "example@email.com",
		Password: "123",
	}

	// Encrypt password, which you can save to database
	ep, err := jwt.EncryptPassword(auth.Password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encrypted Password " + string(ep))

	// Verify Encrypted Password
	valid, err := jwt.VerifyPassword(string(ep), auth.Password)
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

</p>
</details>

<details><summary><b>Decrypt Password</b></summary>
<p>

No you can't, as the thread at [Stack Exchange](https://security.stackexchange.com/questions/193943/is-it-possible-to-decrypt-bcrypt-encryption)

> bcrypt is not an encryption function, it's a password hashing function, relying on Blowfish's key scheduling, not its encryption. Hashing are mathematical one-way functions, meaning there is no way to reverse the output string to get the input string.
  <br/> of course only Siths deal in absolutes and there are a few attacks against hashes. But none of them are "reversing" the hashing, AFAIK.

so that enough to secure the password

</p>
</details>

<details><summary><b>Set Expired Time</b></summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/supanadit/jwt-go"
	"log"
)

func main() {
	// Set Your JWT Secret Code, its optional but important, because default secret code is very insecure
	jwt.SetJWTSecretCode("Your Secret Code")
 
    // You can simply do this, jwt.setExpiredTime(Hour,Minute,Second)
	jwt.SetExpiredTime(0, 0, 1)
}
```

</p>
</details>

<details><summary><b>Support Gin Web Framework</b></summary>
<p>

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/supanadit/jwt-go"
	"net/http"
)

func main() {
	// Set Your JWT Secret Code, its optional but important, because default secret code is very insecure
	jwt.SetJWTSecretCode("Your Secret Code")

	// Create authorization
	auth := jwt.Authorization{
		Username: "admin",
		Password: "123",
	}

	router := gin.Default()

	// Login / Authorization for create JWT Token
	router.POST("/auth", func(c *gin.Context) {
		var a jwt.Authorization
		err := c.Bind(&a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Invalid body request",
				"token":  nil,
			})
		} else {
			valid, err := auth.VerifyPassword(a.Password)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": "Wrong username or password",
					"token":  nil,
				})
			} else {
				if valid {
					token, err := a.GenerateJWT()
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{
							"status": "Can't generate JWT token",
							"token":  nil,
						})
					} else {
						c.JSON(http.StatusOK, gin.H{
							"status": "Success",
							"token":  token,
						})
					}
				} else {
					c.JSON(http.StatusBadRequest, gin.H{
						"status": "Wrong username or password",
						"token":  nil,
					})
				}
			}
		}
	})

	// Test Authorization
	router.GET("/test", func(c *gin.Context) {
		// Variable for binding if you need decoded JWT
		var dataAuth jwt.Authorization
		// Verify and binding JWT
		token, valid, err := jwt.VerifyAndBindingGinHeader(&dataAuth, c)

		// in case if you don't want to decode the JWT, simply use this code
		// token, valid, err := jwt.VerifyGinHeader(c)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": err.Error(),
			})
		} else {
			if valid {
				c.JSON(http.StatusOK, gin.H{
					"status": token + " is valid",
				})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": "Invalid",
				})
			}
		}
	})

	_ = router.Run(":8080")
}
```

</p>
</details>

<details><summary><b>Support Echo Web Framework</b></summary>
<p>

```go
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/supanadit/jwt-go"
	"net/http"
)

func main() {
	// Set Your JWT Secret Code, its optional but important, because default secret code is very insecure
	jwt.SetJWTSecretCode("Your Secret Code")

	// Create authorization
	auth := jwt.Authorization{
		Username: "admin",
		Password: "123",
	}

	e := echo.New()

	// Login / Authorization for create JWT Token
	e.POST("/auth", func(c echo.Context) error {
		a := new(jwt.Authorization)
		// Create struct for response, or you can create globally by your self
		var r struct {
			Status string
			Token  string
		}
		err := c.Bind(a)
		if err != nil {
			r.Status = "Invalid body request"
			return c.JSON(http.StatusBadRequest, &r)
		} else {
			valid, err := auth.VerifyPassword(a.Password)
			if err != nil {
				r.Status = "Wrong username or password"
				return c.JSON(http.StatusBadRequest, &r)
			} else {
				if valid {
					token, err := a.GenerateJWT()
					if err != nil {
						r.Status = "Can't generate JWT Token"
						return c.JSON(http.StatusInternalServerError, &r)
					} else {
						r.Status = "Success"
						r.Token = token
						return c.JSON(http.StatusOK, &r)
					}
				} else {
					r.Status = "Wrong username or password"
					return c.JSON(http.StatusBadRequest, &r)
				}
			}
		}
	})

	// Test Authorization
	e.GET("/test", func(c echo.Context) error {
		// Create struct for response
		var r struct {
			Status string
		}
		// Variable for binding if you need decoded JWT
		dataAuth := new(jwt.Authorization)
		// Verify and binding JWT
		token, valid, err := jwt.VerifyAndBindingEchoHeader(&dataAuth, c)

		// in case if you don't want to decode the JWT, simply use this code
		// Token, valid, err := jwt.VerifyEchoHeader(c)

		if err != nil {
			r.Status = err.Error()
			return c.JSON(http.StatusBadRequest, &r)
		} else {
			if valid {
				r.Status = token + " is valid"
				return c.JSON(http.StatusOK, &r)
			} else {
				r.Status = "Invalid"
				return c.JSON(http.StatusBadRequest, &r)
			}
		}
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
```

</p>
</details>

<details><summary><b>Bonus</b></summary>
<p>

You can simply `Enable` and `Disable` authorization by code bellow

```go
package main

import (
	"github.com/supanadit/jwt-go"
)

func main() {
	// Set Your JWT Secret Code, its optional but important, because default secret code is very insecure
	jwt.SetJWTSecretCode("Your Secret Code")

	jwt.DisableAuthorization() // Disable authorization, meaning when verify jwt token it will return true even if the token was expired or invalid

	// or

	jwt.EnableAuthorization() // Enable authorization
}
```

</p>
</details>
