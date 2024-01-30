package service

import (
	"blablagram/logger"
	"blablagram/repository"
)

type Service struct {
	log logger.Logger
	rep *repository.Repository
}

// New is a constructor for handlers.
func New(log logger.Logger, rep *repository.Repository) *Service {
	return &Service{
		log: log,
		rep: rep,
	}
}
