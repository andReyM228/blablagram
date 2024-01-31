package service

import (
	"blablagram/logger"
	"blablagram/models"
	"blablagram/repository"
	"blablagram/service/user"
	"context"
)

type Service struct {
	UserService
}

type UserService interface {
	RegisterUser(ctx context.Context, user *models.RegisterUser) error
}

// New constructs a new service.
func New(log logger.Logger, rep *repository.Repository, salt string) (*Service, error) {
	userService := user.New(log, rep, salt)

	return &Service{
		UserService: userService,
	}, nil
}
