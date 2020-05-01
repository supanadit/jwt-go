package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

// SigningHMAC is type of signing that can be used for signing and validation JWT
type SigningHMAC struct {
	Method *jwt.SigningMethodHMAC
}

// HS256 is type of SigningHMAC method
func HS256() SigningHMAC {
	return SigningHMAC{Method: jwt.SigningMethodHS256}
}

// HS384 is type of SigningHMAC method
func HS384() SigningHMAC {
	return SigningHMAC{Method: jwt.SigningMethodHS384}
}

// HS512 is type of SigningHMAC method
func HS512() SigningHMAC {
	return SigningHMAC{Method: jwt.SigningMethodHS512}
}
