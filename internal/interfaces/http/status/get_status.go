package status

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
}

func (h *Handler) GetStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (h *Handler) PostName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	fmt.Fprintf(w, "OK %s", name)
}
