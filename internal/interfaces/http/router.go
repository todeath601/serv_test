package http

import (
	handlers "mods/internal/interfaces"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (s *Server) ConfigServer(handler http.Handler) error {
	server := &http.Server{
		Addr:         s.Host + s.Port,
		Handler:      handler,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
	}

	return server.ListenAndServe()
}

func NewRouter() chi.Router {
	h := handlers.NewHandler()
	return h.Router
}
