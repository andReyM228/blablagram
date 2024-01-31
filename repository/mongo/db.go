package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

// wrap is a wrapper for mongodb errors.
var wrap = func(msg string, err error) error {
	return fmt.Errorf("mongodb error: %s: %w", msg, err)
}

// MongoDB is a representation of mongo db service.
type MongoDB struct {
	mongoURL   string
	dbName     string
	client     *mongo.Client
	clientOnce sync.Once
}

// New is a constructor for mongodb.
func New(ctx context.Context, mongoURL string) *MongoDB {
	mongoDB := MongoDB{
		mongoURL: mongoURL,
		dbName:   "dev",
	}

	// Initialize the client only once
	mongoDB.clientOnce.Do(func() {
		var err error
		mongoDB.client, err = mongoDB.connect(ctx)
		if err != nil {
			panic(err)
		}
	})

	// Send a ping to confirm a successful connection
	if err := mongoDB.client.Database(mongoDB.dbName).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return &mongoDB
}

func (m *MongoDB) connect(ctx context.Context) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(m.mongoURL).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return client, err
	}

	return client, nil
}

// Create inserts a new document into the collection.
func (m *MongoDB) Create(ctx context.Context, collectionName string, object interface{}) error {
	collection := m.client.Database(m.dbName).Collection(collectionName)

	_, err := collection.InsertOne(ctx, object)
	if err != nil {
		return wrap("failed to insert document", err)
	}

	return nil
}

// Get finds a single document from the collection.
func (m *MongoDB) Get(ctx context.Context, collectionName string, filter bson.M) (*mongo.SingleResult, error) {
	collection := m.client.Database(m.dbName).Collection(collectionName)

	result := collection.FindOne(ctx, filter)

	return result, nil
}

// Update updates a single document in the collection.
func (m *MongoDB) Update(ctx context.Context, collectionName string, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	collection := m.client.Database(m.dbName).Collection(collectionName)

	result, err := collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return nil, wrap("failed to update document", err)
	}

	return result, nil
}

// Delete deletes a single document from the collection.
func (m *MongoDB) Delete(ctx context.Context, collectionName string, filter bson.M) (*mongo.DeleteResult, error) {
	collection := m.client.Database(m.dbName).Collection(collectionName)

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, wrap("failed to delete document", err)
	}

	return result, nil
}
