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
