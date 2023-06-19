package utils

import (
	"github.com/gofrs/uuid"
)

func GenerateUUID() string {
	uuID := uuid.Must(uuid.NewV7()).String()

	return uuID

}
