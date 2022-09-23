package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hassPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hassPass), nil
}

func CheckPass(hassPassword, passwordDb string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hassPassword), []byte(passwordDb))
	if err != nil {
		return false
	}
	return true
}
