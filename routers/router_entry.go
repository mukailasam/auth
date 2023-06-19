package routers

import (
	"github.com/ftsog/auth/handlers"
	"github.com/go-chi/chi"
)

type Router struct {
	Mux     *chi.Mux
	Handler *handlers.Handler
}
