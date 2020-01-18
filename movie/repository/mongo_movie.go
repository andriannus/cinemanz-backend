package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"cinemanz/models"
	"cinemanz/movie"
)

type mongoMovieRepo struct {
	DB *mongo.Database
}

// NewMongoMovieRepository will create an object that represent the movie.Repository interface
func NewMongoMovieRepository(db *mongo.Database) movie.Repository {
	return &mongoMovieRepo{
		DB: db,
	}
}

func (m *mongoMovieRepo) FetchAll(skip int64, limit int64) (movies []models.Movie, total int64, err error) {
	filter := bson.M{}
	options := options.Find().SetSkip(skip).SetLimit(limit)

	movies, total, err = m.fetchMovies(filter, options)

	return movies, total, err
}

func (m *mongoMovieRepo) FetchNowPlaying(skip int64, limit int64, date string) (movies []models.Movie, total int64, err error) {
	filter := bson.M{
		"$and": []interface{}{
			bson.M{
				"start": bson.M{
					"$lte": date,
				},
			},
			bson.M{
				"end": bson.M{
					"$gte": date,
				},
			},
		},
	}
	options := options.Find().SetSkip(skip).SetLimit(limit)

	movies, total, err = m.fetchMovies(filter, options)

	return movies, total, err
}

func (m *mongoMovieRepo) FetchUpcoming(skip int64, limit int64, date string) (movies []models.Movie, total int64, err error) {
	filter := bson.M{
		"start": bson.M{
			"$gte": date,
		},
	}
	options := options.Find().SetSkip(skip).SetLimit(limit)

	movies, total, err = m.fetchMovies(filter, options)

	return movies, total, err
}

func (m *mongoMovieRepo) FetchByID(id primitive.ObjectID) (movie *models.Movie, err error) {
	err = m.DB.Collection("movie").FindOne(
		context.Background(),
		bson.M{
			"_id": id,
		}).Decode(&movie)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (m *mongoMovieRepo) Store(mv models.Movie) (err error) {
	_, err = m.DB.Collection("movie").InsertOne(context.Background(), mv)

	return err
}

func (m *mongoMovieRepo) Update(movie *models.Movie) (err error) {
	_, err = m.DB.Collection("movie").UpdateOne(
		context.Background(),
		bson.M{"_id": movie.ID},
		bson.M{"$set": bson.M{
			"title":      movie.Title,
			"genre":      movie.Genre,
			"minutes":    movie.Minutes,
			"synopsis":   movie.Synopsis,
			"producer":   movie.Producer,
			"production": movie.Production,
			"director":   movie.Director,
			"writer":     movie.Writer,
			"cast":       movie.Cast,
			"start":      movie.Start,
			"end":        movie.End,
			"rate":       movie.Rate,
			"theater":    movie.Theater,
			"poster":     movie.Poster,
		}})

	return err
}

func (m *mongoMovieRepo) Delete(id primitive.ObjectID) (err error) {
	_, err = m.DB.Collection("movie").DeleteOne(
		context.Background(),
		bson.M{
			"_id": id,
		})

	return err
}

func (m *mongoMovieRepo) fetchMovies(filter primitive.M, options *options.FindOptions) (movies []models.Movie, total int64, err error) {
	cursor, err := m.DB.Collection("movie").Find(
		context.Background(),
		filter,
		options,
	)

	if err != nil {
		return nil, 0, err
	}

	total, err = m.DB.Collection("movie").CountDocuments(
		context.Background(),
		filter,
	)

	if err != nil {
		return nil, 0, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var movie models.Movie

		err := cursor.Decode(&movie)

		if err != nil {
			return nil, 0, err
		}

		movies = append(movies, movie)
	}

	return movies, total, nil
}
