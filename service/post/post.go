package post

import (
	"blablagram/logger"
	"blablagram/models"
	"blablagram/repository"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"time"
)

// wrap is a wrapper for mongodb errors.
var wrap = func(msg string, err error) error {
	return fmt.Errorf("post-service error: %s: %w", msg, err)
}

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

// CreatePost creates a new post and saves it to the database.
func (s *Service) CreatePost(ctx context.Context, post *models.Post) error {
	post.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	id := uuid.New().String()
	id = strings.ReplaceAll(id, "-", "")
	post.ID = id

	if err := s.rep.Mongo.Create(ctx, "posts", post); err != nil {
		return wrap("failed to create post", err)
	}

	return nil
}

// DeletePost deletes a post from the database by its id.
func (s *Service) DeletePost(ctx context.Context, id string) error {
	filter := bson.M{"id": id}

	if _, err := s.rep.Mongo.Delete(ctx, "posts", filter); err != nil {
		return wrap("failed to delete post", err)
	}

	return nil
}

// UpdatePost updates a post in the database by its id.
func (s *Service) UpdatePost(ctx context.Context, post *models.Post) error {
	filter := bson.M{"id": post.ID}

	update := bson.M{
		"description": post.Description,
	}

	_, err := s.rep.Mongo.Update(ctx, "posts", filter, update)
	if err != nil {
		return wrap("failed to update post", err)
	}

	return nil
}

// GetPost gets a post from the database by its id.
func (s *Service) GetPost(ctx context.Context, id string) (*models.Post, error) {
	filter := bson.M{"id": id}

	result, err := s.rep.Mongo.Get(ctx, "posts", filter)
	if err != nil {
		return nil, wrap("failed to get post", err)
	}

	var post models.Post
	if err := result.Decode(&post); err != nil {
		return nil, wrap("failed to decode post", err)
	}

	return &post, nil
}

// GetPosts gets all posts from the database.
func (s *Service) GetPosts(ctx context.Context) ([]models.Post, error) {
	filter := bson.M{}

	cursor, err := s.rep.Mongo.Get(ctx, "posts", filter)
	if err != nil {
		return nil, wrap("failed to get posts", err)
	}

	var posts []models.Post
	if err := cursor.Decode(&posts); err != nil {
		return nil, wrap("failed to decode posts", err)
	}

	return posts, nil
}
