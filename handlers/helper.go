package handlers

import (
	"log"
	"net/http"
)

type Logger struct {
	Info    *log.Logger
	Warning *log.Logger
	Debug   *log.Logger
	Error   *log.Logger
	Fatal   *log.Logger
}

func SetResponseStatusCode(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	return
}
