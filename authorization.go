package ej

// Default authorization model
type Authorization struct {
	Username string `form:"username" json:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	Claims
}

// Generate JWT Token by authorization model
func (auth Authorization) GenerateJWT() (string, error) {
	return GenerateJWT(auth)
}

// Create password encryption using bcrypt
func (auth Authorization) GetPasswordEncryption() ([]byte, error) {
	return EncryptPassword(auth.Password)
}

// Create password encryption using bcrypt and string as the result
func (auth Authorization) GetStringPasswordEncryption() (string, error) {
	encryption, err := auth.GetPasswordEncryption()
	return string(encryption), err
}

// VerifyJWT the authorization
func (auth Authorization) VerifyJWT(token string) (bool, error) {
	return VerifyJWT(auth, token)
}
