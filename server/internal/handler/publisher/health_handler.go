package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type HealthHandler struct{}

func NewHealthHandler(r chi.Router) {
	handler := &HealthHandler{}

	r.Get("/health/alive", handler.Alive)
}

func (h *HealthHandler) Alive(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{"message": "ok"})
}
