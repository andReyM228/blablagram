package handlers

import (
	"blablagram/logger"
	"blablagram/models"
	"blablagram/service"
	"encoding/json"
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

// RegisterUser is a handler for user registration, it makes email and password validation.
func (h *Handlers) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.RegisterUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		newErrorResponse(w, h.log, http.StatusBadRequest, "invalid input body", err)
		return
	}

	if err := h.service.UserService.RegisterUser(r.Context(), &user); err != nil {
		newErrorResponse(w, h.log, http.StatusInternalServerError, "failed to register user", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
