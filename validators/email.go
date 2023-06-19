package validators

import (
	"regexp"

	"github.com/ftsog/auth/customerrors"
)

func ValidateEmail(email string) (*string, error) {

	emailRegexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegexp.MatchString(email) {
		return nil, customerrors.InvalidEmail
	}

	return &email, nil

}
