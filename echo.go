package jwt

import (
	"errors"
	"github.com/labstack/echo/v4"
	"strings"
)

func GetJWTFromEchoHeader(c echo.Context) (token string, err error) {
	header := c.Request().Header.Get("Authorization")
	if header != "" {
		splitAuthorization := strings.Split(header, " ")
		if len(splitAuthorization) != 0 && len(splitAuthorization) == 2 {
			token = splitAuthorization[1]
		} else {
			err = errors.New("invalid authorization header")
		}
	} else {
		err = errors.New("no authorization provided")
	}
	return token, err
}

func VerifyEchoHeader(c echo.Context) (string, bool, error) {
	return VerifyAndBindingEchoHeader(nil, c)
}

func VerifyAndBindingEchoHeader(model interface{}, c echo.Context) (token string, isValid bool, err error) {
	token, err = GetJWTFromEchoHeader(c)
	if err == nil {
		isValid, err = VerifyAndBindingJWT(&model, token)
	}
	if err != nil {
		if !IsUseAuthorization() {
			err = nil
			isValid = !IsUseAuthorization()
		}
	}
	return token, isValid, err
}
