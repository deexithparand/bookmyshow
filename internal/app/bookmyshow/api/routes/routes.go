package routes

import (
	"github.com/deexithparand/bookmyshow/internal/app/bookmyshow/api/handlers"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux, h *handlers.Handlers) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/create", h.UserHandler.CreateUser)
	})
}
