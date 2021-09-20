package contract

import "go.mongodb.org/mongo-driver/bson/primitive"

type KVRecord struct {
	ID        primitive.ObjectID `bson:"_id"`
	Counts    []int              `bson:"counts"`
	CreatedAt primitive.DateTime `bson:"created_at"`
	Key       string             `bson:"key"`
	Value     string             `bson:"value"`
}