package http

import (
	"fmt"
	"mods/internal/interfaces/http/router"
	"time"
)

func StartServer() {
	s := &Server{
		Host:         "localhost",
		Port:         8080,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	r := router.BuildRouter()

	fmt.Printf("Startin server on port %d...\n", s.Port)
	if err := s.ConfigServer(r); err != nil {
		fmt.Println("Server failed to start")
	}
}
