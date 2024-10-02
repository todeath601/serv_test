package http

import (
	"fmt"
	"net/http"
)

func Router() {
	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8080", nil)
}
