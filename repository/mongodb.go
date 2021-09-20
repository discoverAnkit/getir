package repository

import (
	"context"
	"github.com/discoverAnkit/getir/contract"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type MongoClient interface {
	GetRecordsByCreationTime(ctx context.Context, startTime, endTime time.Time) ([]contract.KVRecord,error)
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

func (m *MongoStore) GetRecordsByCreationTime(ctx context.Context, startTime, endTime time.Time) ([]contract.KVRecord,error) {

	var keyValueRecords []contract.KVRecord

	//I didn't find(https://docs.mongodb.com/v5.0/tutorial/query-arrays/) any operator to do sum on array elements directly while querying
	//Not sure if there are other ways - aggregation? pipelines - I need more time to understand them and decide if they can be used or not in this use case
	//So filtering is based on just dates - which will not scale if there are too many keyValueRecords in that date range

	filter := bson.D{
		{"createdAt", bson.D{{"$gt", primitive.NewDateTimeFromTime(startTime)}}},
		{"createdAt", bson.D{{"$lt", primitive.NewDateTimeFromTime(endTime)}}},
	}

	findCursor, findErr := m.Collection.Find(ctx, filter)
	if findErr != nil {
		log.Println("Something went wrong in querying mongoDB",findErr)
		return keyValueRecords,findErr
	}

	if findErr = findCursor.All(context.TODO(), &keyValueRecords); findErr != nil {
		log.Println("Something went wrong in querying mongoDB",findErr)
		return keyValueRecords,findErr
	}

	return keyValueRecords, nil
}