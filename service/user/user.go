package user

import (
	"blablagram/logger"
	"blablagram/models"
	"blablagram/repository"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

// wrap is a wrapper for mongodb errors.
var wrap = func(msg string, err error) error {
	return fmt.Errorf("user-service error: %s: %w", msg, err)
}

type Service struct {
	log  logger.Logger
	rep  *repository.Repository
	salt string
}

// New is a constructor for handlers.
func New(log logger.Logger, rep *repository.Repository, salt string) *Service {
	return &Service{
		log:  log,
		rep:  rep,
		salt: salt,
	}
}

// RegisterUser makes email and password validation, hashes password and saves user to database.
func (s *Service) RegisterUser(ctx context.Context, user *models.RegisterUser) error {
	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	id := uuid.New().String()
	id = strings.ReplaceAll(id, "-", "")
	user.ID = id

	if !IsEmailValid(user.Email) {
		return errors.New("invalid email format")
	}

	if err := PasswordValidation(user.Password); err != nil {
		return wrap("invalid password", err)
	}

	hashedPassword := hashPasswordWithSalt(user.Password, s.salt)

	user.Password = hashedPassword

	if err := s.rep.Mongo.Create(ctx, "users", user); err != nil {
		return wrap("failed to create user", err)
	}

	return nil
}
