package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"cinemanz/databases"
)

// Theater struct for theater
type Theater struct {
	ID        primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Name      string             `json:"name" bson:"name"`
	Address   string             `json:"address" bson:"address"`
	Telephone string             `json:"telephone" bson:"telephone"`
}

// FetchTheaters return all theater
func FetchTheaters(skip int64, limit int64) (theaters []Theater, total int64, err error) {
	options := options.Find().SetSkip(skip).SetLimit(limit)

	cursor, err := databases.Mongo.Collection("theater").Find(
		context.Background(),
		bson.M{},
		options,
	)

	if err != nil {
		return nil, 0, err
	}

	total, err = databases.Mongo.Collection("theater").CountDocuments(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		return nil, 0, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var theater Theater

		err := cursor.Decode(&theater)

		if err != nil {
			return nil, 0, err
		}

		theaters = append(theaters, theater)
	}

	return theaters, total, nil
}

// FetchTheater return a theater
func FetchTheater(objectID primitive.ObjectID) (theater *Theater, err error) {
	err = databases.Mongo.Collection("theater").FindOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		}).Decode(&theater)

	if err != nil {
		return nil, err
	}

	return theater, nil
}

// AddTheater create a theater
func AddTheater(t Theater) error {
	var err error

	_, err = databases.Mongo.Collection("theater").InsertOne(context.Background(), t)

	return err
}

// UpdateTheater update a theater
func UpdateTheater(objectID primitive.ObjectID, t Theater) error {
	_, err := databases.Mongo.Collection("theater").UpdateOne(
		context.Background(),
		bson.M{"_id": objectID},
		bson.M{"$set": bson.M{
			"address":   t.Address,
			"name":      t.Name,
			"telephone": t.Telephone,
		}})

	return err
}

// DeleteTheater delete a theater
func DeleteTheater(objectID primitive.ObjectID) error {
	_, err := databases.Mongo.Collection("theater").DeleteOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		})

	return err
}
