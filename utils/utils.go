package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// generate random data
func Token() string {

	salt := make([]byte, 32)
	rand.Read(salt)

	output := hex.EncodeToString(salt)
	output = strings.TrimSpace(output)
	return output
}

func HashRegisterPassword(password string) (string, string) {

	salt := Token()

	newPassword := password + salt

	hashedPassword := sha256.Sum256([]byte(newPassword))
	passwordToString := hex.EncodeToString(hashedPassword[:])
	password = strings.TrimSpace(passwordToString)

	return salt, password
}

func HashLoginPassword(password, salt string) string {
	salt = strings.TrimSpace(salt)

	output := password + salt
	output = strings.TrimSpace(output)

	hashedPassword := sha256.Sum256([]byte(output))
	passwordToString := hex.EncodeToString(hashedPassword[:])
	password = strings.TrimSpace(passwordToString)

	return password
}
