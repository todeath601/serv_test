package http

import (
	"fmt"
	"net/http"
	"strings"
)

func GetReq(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func PostName(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	name := strings.TrimPrefix(r.URL.Path, "/api/ping/")

	fmt.Fprintf(w, "OK %s\n", name)
}
