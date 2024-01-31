package repository

import (
	"blablagram/repository/mongo"
	"context"
	"go.mongodb.org/mongo-driver/bson"

	mongoDriver "go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Mongo
}

type Mongo interface {
	Create(ctx context.Context, collectionName string, object interface{}) error
	Get(ctx context.Context, collectionName string, filter bson.M) (*mongoDriver.SingleResult, error)
	Update(ctx context.Context, collectionName string, filter bson.M, update bson.M) (*mongoDriver.UpdateResult, error)
	Delete(ctx context.Context, collectionName string, filter bson.M) (*mongoDriver.DeleteResult, error)
}

// New constructs a new repository.
func New(ctx context.Context, mongoUrl string) (*Repository, error) {
	dbMongo := mongo.New(ctx, mongoUrl)

	return &Repository{
		Mongo: dbMongo,
	}, nil
}
