package service

import (
	"blablagram/logger"
	"blablagram/models"
	"blablagram/repository"
	"blablagram/service/post"
	"blablagram/service/user"
	"context"
)

type Service struct {
	UserService
	PostService
}

type UserService interface {
	RegisterUser(ctx context.Context, user *models.User) error
	LoginUser(ctx context.Context, user *models.LoginUser) (*models.User, error)
}

type PostService interface {
	CreatePost(ctx context.Context, post *models.Post) error
	DeletePost(ctx context.Context, id string) error
	UpdatePost(ctx context.Context, post *models.Post) error
	GetPost(ctx context.Context, id string) (*models.Post, error)
	GetPosts(ctx context.Context) ([]models.Post, error)
}

// New constructs a new service.
func New(log logger.Logger, rep *repository.Repository, salt string) (*Service, error) {
	userService := user.New(log, rep, salt)

	postService := post.New(log, rep)

	return &Service{
		UserService: userService,
		PostService: postService,
	}, nil
}
