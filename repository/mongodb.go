package repository

import (
	"context"
	"github.com/discoverAnkit/getir/contract"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type MongoClient interface {
	GetRecords(ctx context.Context, filters contract.GetCountsRequest) error
}

type MongoStore struct {
	MongoClient *mongo.Client
	Collection  *mongo.Collection
}

func NewMongoStore (mongoURI, database, collection string) *MongoStore {

	// Create a new mongo client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to MongoDB and pinged.")

	return &MongoStore{
		MongoClient: client,
		Collection: client.Database(database).Collection(collection),
	}
}

func (m *MongoStore) GetRecords(ctx context.Context, filters contract.GetCountsRequest) error {

	return nil
}