package validators

import (
	"strings"
)

func IsEmpty(username, email, password string) bool {
	username = strings.TrimSpace(username)
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(email)
	if username == "" || email == "" || password == "" {
		return true
	}

	return false
}
