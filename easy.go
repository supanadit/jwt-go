package ej

// Default configuration
var secretCodeJWT string = "Default Secret Code"
var useAuthorization bool = true

// Get JWT Secret Code as a string
func GetStringJWTSecretCode() string {
	return secretCodeJWT
}

// Get JWT Secret Code as a byte
func GetJWTSecretCode() []byte {
	return []byte(secretCodeJWT)
}

// Set JWT Secret Code
func SetJWTSecretCode(secretCode string) {
	secretCodeJWT = secretCode
}

// Get status whether use authorization or not
func IsUseAuthorization() bool {
	return useAuthorization
}

// Enable authorization
func EnableAuthorization() {
	useAuthorization = true
}

// Disable the authorization
func DisableAuthorization() {
	useAuthorization = false
}
