package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"cinemanz/databases"
)

// Movie struct for movie
type Movie struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Title      string             `json:"title" bson:"title"`
	Genre      []string           `json:"genre" bson:"genre"`
	Minutes    int                `json:"minutes" bson:"minutes"`
	Synopsis   string             `json:"synopsis" bson:"synopsis"`
	Producer   []string           `json:"producer" bson:"producer"`
	Production string             `json:"production" bson:"production"`
	Director   string             `json:"director" bson:"director"`
	Writer     string             `json:"writer" bson:"writer"`
	Cast       []string           `json:"cast" bson:"cast"`
	Start      string             `json:"start" bson:"start"`
	End        string             `json:"end" bson:"end"`
	Rate       float32            `json:"rate" bson:"rate"`
	Theater    []string           `json:"theater" bson:"theater"`
	Poster     string             `json:"poster" bson:"poster"`
}

// FetchMovies return all movie
func FetchMovies(skip int64, limit int64) (movies []Movie, total int64, err error) {
	options := options.Find().SetSkip(skip).SetLimit(limit)

	cursor, err := databases.Mongo.Collection("movie").Find(context.Background(), bson.M{}, options)

	if err != nil {
		return nil, 0, err
	}

	total, err = databases.Mongo.Collection("movie").CountDocuments(context.Background(), bson.M{})

	if err != nil {
		return nil, 0, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var movie Movie

		err := cursor.Decode(&movie)

		if err != nil {
			return nil, 0, err
		}

		movies = append(movies, movie)
	}

	return movies, total, nil
}

// FetchMovie return a movie
func FetchMovie(objectID primitive.ObjectID) (movie *Movie, err error) {
	err = databases.Mongo.Collection("movie").FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&movie)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

// AddMovie create a movie
func AddMovie(m Movie) error {
	_, err := databases.Mongo.Collection("movie").InsertOne(context.Background(), m)

	return err
}

// UpdateMovie update a movie
func UpdateMovie(objectID primitive.ObjectID, m Movie) error {
	_, err := databases.Mongo.Collection("movie").UpdateOne(
		context.Background(),
		bson.M{"_id": objectID},
		bson.M{"$set": bson.M{
			"title":      m.Title,
			"genre":      m.Genre,
			"minutes":    m.Minutes,
			"synopsis":   m.Synopsis,
			"producer":   m.Producer,
			"production": m.Production,
			"director":   m.Director,
			"writer":     m.Writer,
			"cast":       m.Cast,
			"start":      m.Start,
			"end":        m.End,
			"rate":       m.Rate,
			"theater":    m.Theater,
			"poster":     m.Poster,
		}})

	return err
}

// DeleteMovie delete a movie
func DeleteMovie(objectID primitive.ObjectID) error {
	_, err := databases.Mongo.Collection("movie").DeleteOne(context.Background(), bson.M{"_id": objectID})

	return err
}
