package validators

import (
	"unicode"

	"github.com/ftsog/auth/customerrors"
)

func ValidatePassword(password string) (*string, error) {
	var (
		hasMin     = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	if len(password) >= 8 {
		hasMin = true
	}

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsLower(char) {
			hasLower = true
		} else if unicode.IsNumber(char) {
			hasNumber = true
		} else if unicode.IsPunct(char) {
			hasSpecial = true
		}
	}

	if !hasMin || !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return nil, customerrors.InvalidPassword
	}

	return &password, nil
}
