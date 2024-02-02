package user

import (
	"blablagram/logger"
	"blablagram/models"
	"blablagram/repository"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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
func (s *Service) RegisterUser(ctx context.Context, user *models.User) error {
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

	user.IsLogged = true

	if err := s.rep.Mongo.Create(ctx, "users", user); err != nil {
		return wrap("failed to create user", err)
	}

	return nil
}

// LoginUser checks if user exists and if password is correct.
func (s *Service) LoginUser(ctx context.Context, user *models.LoginUser) (*models.User, error) {
	filter := map[string]interface{}{
		"email": user.Email,
	}

	result, err := s.rep.Mongo.Get(ctx, "users", filter)
	if err != nil {
		return nil, wrap("failed to get user", err)
	}

	var userFromDB models.User
	if err := result.Decode(&userFromDB); err != nil {
		return nil, wrap("failed to decode user", err)
	}

	hashPassword := hashPasswordWithSalt(user.Password, s.salt)
	if userFromDB.Password != hashPassword {
		return nil, errors.New("invalid password")
	}

	userFromDB.IsLogged = true

	update := bson.M{
		"isLogged": userFromDB.IsLogged,
	}

	if _, err := s.rep.Mongo.Update(ctx, "users", filter, update); err != nil {
		return nil, wrap("failed to update user", err)
	}

	return &userFromDB, nil
}
