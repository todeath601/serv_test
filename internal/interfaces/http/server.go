package http

import (
	"fmt"
	"mods/internal/interfaces/config"
	"net/http"
	"time"
)

type server struct {
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func StartServer(config config.Config) error {
	s := &server{
		Host:         config.Server.Host,
		Port:         config.Server.Port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	r := GetRouter()

	fmt.Printf("Startin server on port %d...\n", s.Port)
	if err := s.configServer(r); err != nil {
		fmt.Println("Server failed to start")
		return err
	}

	return nil
}

func (s *server) configServer(handler http.Handler) error {
	address := fmt.Sprintf("%s:%d", s.Host, s.Port)

	server := &http.Server{
		Addr:         address,
		Handler:      handler,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
	}

	return server.ListenAndServe()
}
