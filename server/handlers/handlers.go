package handlers

import (
	"blablagram/logger"
	"blablagram/service"
	"net/http"
)

type Handlers struct {
	log     logger.Logger
	service *service.Service
}

// New is a constructor for handlers.
func New(log logger.Logger, service *service.Service) *Handlers {
	return &Handlers{
		log:     log,
		service: service,
	}
}

// Status is a handler for status.
func (h *Handlers) Status(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
