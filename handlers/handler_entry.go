package handlers

import (
	"github.com/ftsog/auth/models"
)

type Handler struct {
	Model *models.Model
	Log   *Logger
}
