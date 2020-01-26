package movie

import (
	"cinemanz/models"
)

// Usecase represent the movie's usecases
type Usecase interface {
	FetchAll(skip int64, limit int64) (movies []models.Movie, total int64, err error)
	FetchNowPlaying(skip int64, limit int64, date string) (movies []models.Movie, total int64, err error)
	FetchUpcoming(skip int64, limit int64, date string) (movies []models.Movie, total int64, err error)
	FetchByID(id string) (movie *models.Movie, err error)
	Store(movie models.Movie) error
	Update(id string, movie *models.Movie) error
	Delete(id string) error
}
