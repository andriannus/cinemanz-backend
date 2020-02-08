package mongo

import (
	"context"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Setup will create setup for Mongo
func Setup() (*mongo.Database, error) {
	dbURI := viper.GetString(`database.mongo.uri`)
	dbName := viper.GetString(`database.mongo.name`)

	options := options.Client().ApplyURI(dbURI)
	client, err := mongo.NewClient(options)

	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())

	if err != nil {
		return nil, err
	}

	return client.Database(dbName), nil
}
