package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Router chi.Router
}

func NewHandler() *Handler {
	h := &Handler{
		Router: chi.NewRouter(),
	}
	h.routes()
	return h
}

func (h *Handler) routes() {
	h.Router.Get("/api/ping", h.GetReq)
	h.Router.Post("/api/ping/{name}", h.PostName)
}

func (h *Handler) GetReq(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (h *Handler) PostName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	fmt.Fprintf(w, "OK %s", name)
}
