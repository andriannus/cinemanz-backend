package movie

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cinemanz/models"
)

// Repository represent the movie's repositories contract
type Repository interface {
	FetchAll(skip int64, limit int64) (movies []models.Movie, total int64, err error)
	FetchNowPlaying(skip int64, limit int64, date string) (movies []models.Movie, total int64, err error)
	FetchUpcoming(skip int64, limit int64, date string) (movies []models.Movie, total int64, err error)
	FetchByID(id primitive.ObjectID) (movie *models.Movie, err error)
	Store(m models.Movie) error
	Update(movie *models.Movie) error
	Delete(id primitive.ObjectID) error
}
