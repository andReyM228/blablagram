package handlers

import (
	"blablagram/logger"
	"encoding/json"
	"fmt"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(w http.ResponseWriter, log logger.Logger, statusCode int, msg string, err error) {
	err = fmt.Errorf("handlers error: %w", err)
	log.Error(msg, err)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorResponse{Message: err.Error()})
}
