package time

import (
	"net/http"
	"time"
)

type timeHandler struct {
	UserServic UserService
}

func NewTimeHandler() *timeHandler {
	return &timeHandler{}
}

func (h *timeHandler) GetTime(w http.ResponseWriter, r *http.Request) {

	UserServic()

	w.Write([]byte(time.Now().String()))
}
