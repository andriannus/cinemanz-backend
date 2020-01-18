package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"cinemanz/models"
	"cinemanz/theater"
)

type mongoTheaterRepo struct {
	DB *mongo.Database
}

// NewMongoTheaterRepository will create an object that represent the theater.Repository interface
func NewMongoTheaterRepository(db *mongo.Database) theater.Repository {
	return &mongoTheaterRepo{
		DB: db,
	}
}

func (m *mongoTheaterRepo) FetchAll(skip int64, limit int64) (theaters []models.Theater, total int64, err error) {
	options := options.Find().SetSkip(skip).SetLimit(limit)

	cursor, err := m.DB.Collection("theater").Find(
		context.Background(),
		bson.M{},
		options,
	)

	if err != nil {
		return nil, 0, err
	}

	total, err = m.DB.Collection("theater").CountDocuments(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		return nil, 0, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var theater models.Theater

		err := cursor.Decode(&theater)

		if err != nil {
			return nil, 0, err
		}

		theaters = append(theaters, theater)
	}

	return theaters, total, nil
}

func (m *mongoTheaterRepo) FetchByID(id primitive.ObjectID) (theater *models.Theater, err error) {
	err = m.DB.Collection("theater").FindOne(
		context.Background(),
		bson.M{
			"_id": id,
		}).Decode(&theater)

	if err != nil {
		return nil, err
	}

	return theater, nil
}

func (m *mongoTheaterRepo) Store(t models.Theater) (err error) {

	_, err = m.DB.Collection("theater").InsertOne(context.Background(), t)

	return err
}

func (m *mongoTheaterRepo) Update(t *models.Theater) (err error) {
	_, err = m.DB.Collection("theater").UpdateOne(
		context.Background(),
		bson.M{"_id": t.ID},
		bson.M{"$set": bson.M{
			"address":   t.Address,
			"name":      t.Name,
			"telephone": t.Telephone,
		}})

	return err
}

func (m *mongoTheaterRepo) Delete(id primitive.ObjectID) (err error) {
	_, err = m.DB.Collection("theater").DeleteOne(
		context.Background(),
		bson.M{
			"_id": id,
		})

	return err
}
