package jwt

import (
	"errors"
	"strings"
)

// GetJWTFromHeader get JWT from header which provided by any of web framework, but with rules JWT "token"
func GetJWTFromHeader(header string) (token string, err error) {
	if header != "" {
		splitAuthorization := strings.Split(header, " ")
		if len(splitAuthorization) != 0 && len(splitAuthorization) == 2 {
			if splitAuthorization[0] != "JWT" {
				err = errors.New("unknown authorization type")
			} else {
				token = splitAuthorization[1]
			}
		} else {
			err = errors.New("invalid authorization header")
		}
	} else {
		err = errors.New("no authorization provided")
	}
	return token, err
}

// VerifyAndBinding is to verify and binding any given token
func VerifyAndBinding(model interface{}, t string) (bool, error) {
	isValid, err := VerifyAndBindingJWT(&model, t)
	if err != nil {
		if !IsUseAuthorization() {
			err = nil
			isValid = !IsUseAuthorization()
		}
	}
	return isValid, err
}
