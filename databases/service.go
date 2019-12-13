package databases

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"cinemanz/constants"
)

// Mongo variable for MongoDB
var Mongo *mongo.Database

// MongoSetup setup MongoDB
func MongoSetup() (*mongo.Database, error) {
	options := options.Client().ApplyURI(constants.DatabaseURI)
	client, err := mongo.NewClient(options)

	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())

	if err != nil {
		return nil, err
	}

	return client.Database(constants.DatabaseName), nil
}
