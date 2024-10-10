package handlers

import (
	"mods/internal/interfaces/http/status"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HandlerInterfaces interface {
	GetStatus(w http.ResponseWriter, r *http.Request)
	PostName(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	Router chi.Router
	Status *status.Handler
	HandlerInterfaces
}

func NewHandler() *Handler {
	h := &Handler{
		Router: chi.NewRouter(),
		Status: &status.Handler{},
	}
	h.routes()
	return h
}

func (h *Handler) routes() {
	h.Router.Get("/api/ping", h.Status.GetStatus)
	h.Router.Post("/api/ping/{name}", h.Status.PostName)
}
