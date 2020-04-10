package ej

var secretCodeJWT string = "Default Secret Code"
var authorizationStatus bool = true

func GetJWTSecretCode() []byte {
	return []byte(secretCodeJWT)
}

func SetJWTSecretCode(secretCode string) {
	secretCodeJWT = secretCode
}

func GetAuthorizationStatus() bool {
	return authorizationStatus
}

func SetAuthorizationStatus(status bool) {
	authorizationStatus = status
}
