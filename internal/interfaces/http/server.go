package http

import (
	"fmt"
	"time"
)

func StartServer() {
	s := &Server{
		Host:         "localhost",
		Port:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	r := NewRouter()

	fmt.Println("Startin server on port :8080...")
	if err := s.ConfigServer(r); err != nil {
		fmt.Println("Server failed to start")
	}
}
