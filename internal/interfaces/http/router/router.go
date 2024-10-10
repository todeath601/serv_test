package router

import (
	"fmt"
	handlers "mods/internal/interfaces"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (s *Server) ConfigServer(handler http.Handler) error {
	address := fmt.Sprintf("%s:%d \n", s.Host, s.Port)
	server := &http.Server{
		Addr:         address,
		Handler:      handler,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
	}

	return server.ListenAndServe()
}

func BuildRouter() chi.Router {
	h := handlers.NewHandler()
	return h.Router
}
