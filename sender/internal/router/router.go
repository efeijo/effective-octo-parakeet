package router

import (
	"encoding/json"
	"sender/internal/model"
	"sender/internal/sender"

	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Router struct {
	*chi.Mux
	sender sender.Sender
	logger *zap.SugaredLogger
}

func NewRouter(sender sender.Sender, logger *zap.SugaredLogger) http.Handler {
	chiRouter := chi.NewRouter()
	router := &Router{
		sender: sender,
		Mux:    chiRouter,
		logger: logger,
	}

	// setupRoutes
	routes(router)

	return router
}

func (r *Router) Send(w http.ResponseWriter, req *http.Request) {
	logger := r.logger
	var message *model.Message
	err := json.NewDecoder(req.Body).Decode(&message)
	if err != nil {
		logger.Errorln("error unmarshalling body", err)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.Infoln("message to send", message)

	err = r.sender.Publish(req.Context(), message, r.logger)
	if err != nil {
		logger.Errorln("error publishing message", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func (r *Router) GetAllMessagesSended(w http.ResponseWriter, req *http.Request) {
	logger := r.logger
	messages, err := r.sender.GetAllPublished(req.Context(), r.logger)
	if err != nil {
		logger.Errorln("error getting all messages", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	marshMessages, err := json.Marshal(messages)
	if err != nil {
		logger.Errorln("error unmarshalling messages", req)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(marshMessages)
}
