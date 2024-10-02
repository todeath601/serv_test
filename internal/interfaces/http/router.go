package http

import (
	"fmt"

	"net/http"
)

func Router() {

	http.HandleFunc("/api/ping", GetReq)
	http.HandleFunc("/api/ping/{name}", PostName)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)

}
