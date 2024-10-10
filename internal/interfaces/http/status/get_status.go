package status

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type statusHandler struct {
}

func NewStatusHandler() *statusHandler {
	return &statusHandler{}
}

func (h *statusHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (h *statusHandler) PostName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	fmt.Fprintf(w, "OK %s", name)
}
