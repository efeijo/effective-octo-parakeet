package router

import (
	"encoding/json"
	"net/http"
	"sender/internal/model"
	"sender/internal/sender"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	*chi.Mux
	sender *sender.Sender
}

func NewRouter(sender *sender.Sender) http.Handler {
	chiRouter := chi.NewRouter()
	router := &Router{
		sender: sender,
	}
	router.Mux = chiRouter

	chiRouter.Use(middleware.Logger)
	chiRouter.Route("/send", func(r chi.Router) {
		r.Post("/", router.Send)
		r.Get("/", router.GetAllSended)
	})

	return router
}

func (r *Router) Send(w http.ResponseWriter, req *http.Request) {
	var message model.Message
	err := json.NewDecoder(req.Body).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

}
func (r *Router) GetAllSended(w http.ResponseWriter, req *http.Request) {}
