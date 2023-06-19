package validators

import (
	"github.com/ftsog/auth/customerrors"
)

const (
	userNameMinLength = 6
	userNameMaxLength = 20
)

func ValidateUsername(username string) (*string, error) {
	if len(username) < userNameMinLength || len(username) > userNameMaxLength {
		return nil, customerrors.InvalidUsername
	}

	return &username, nil
}
