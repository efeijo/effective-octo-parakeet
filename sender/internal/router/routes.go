package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func routes(router *Router) {
	router.Mux.Route("/send", func(r chi.Router) {
		r.Post("/", router.Send)
		r.Get("/", router.GetAllMessagesSended)
	})
	router.Mux.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
