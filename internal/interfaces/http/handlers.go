package http

import (
	"mods/internal/interfaces/http/status"
	"mods/internal/interfaces/http/time"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type StatusHandler interface {
	GetStatus(w http.ResponseWriter, r *http.Request)
	PostName(w http.ResponseWriter, r *http.Request)
}

type TimeHandler interface {
	GetTime(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	Router chi.Router
	status StatusHandler
	time   TimeHandler
}

func GetRouter() chi.Router {
	return newHandler().Router
}

func newHandler() *Handler {
	h := &Handler{
		Router: chi.NewRouter(),
		status: status.NewStatusHandler(),
		time:   time.NewTimeHandler(),
	}
	h.routes()
	return h
}

func (h *Handler) routes() {
	// status
	h.Router.Get("/api/ping", h.status.GetStatus)
	h.Router.Post("/api/ping/{name}", h.status.PostName)

	// time
	h.Router.Get("/api/time", h.time.GetTime)

}
