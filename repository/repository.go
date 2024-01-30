package repository

import (
	"blablagram/repository/mongo"
	"context"
)

type Repository struct {
	Mongo
}

type Mongo interface {
}

// New constructs a new repository.
func New(ctx context.Context, mongoUrl string) (*Repository, error) {
	dbMongo := mongo.New(ctx, mongoUrl)

	return &Repository{
		Mongo: dbMongo,
	}, nil
}
