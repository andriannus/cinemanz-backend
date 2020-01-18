package usecase

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cinemanz/models"
	"cinemanz/movie"
)

type movieUsecase struct {
	movieRepo movie.Repository
}

// NewMovieUsecase will create an object that represent the movie.Usecase interface
func NewMovieUsecase(m movie.Repository) movie.Usecase {
	return &movieUsecase{
		movieRepo: m,
	}
}

func (m *movieUsecase) FetchAll(skip int64, limit int64) (movies []models.Movie, total int64, err error) {
	movies, total, err = m.movieRepo.FetchAll(skip, limit)

	return movies, total, err
}

func (m *movieUsecase) FetchNowPlaying(skip int64, limit int64, date string) (movies []models.Movie, total int64, err error) {
	movies, total, err = m.movieRepo.FetchNowPlaying(skip, limit, date)

	return movies, total, err
}

func (m *movieUsecase) FetchUpcoming(skip int64, limit int64, date string) (movies []models.Movie, total int64, err error) {
	movies, total, err = m.movieRepo.FetchUpcoming(skip, limit, date)

	return movies, total, err
}

func (m *movieUsecase) FetchByID(id string) (movie *models.Movie, err error) {
	objectID, _ := primitive.ObjectIDFromHex(id)

	movie, err = m.movieRepo.FetchByID(objectID)

	return movie, err
}

func (m *movieUsecase) Store(movie models.Movie) (err error) {
	err = m.movieRepo.Store(movie)

	return err
}

func (m *movieUsecase) Update(id string, movie *models.Movie) (err error) {
	objectID, _ := primitive.ObjectIDFromHex(id)

	movie.ID = objectID

	err = m.movieRepo.Update(movie)

	return err
}

func (m *movieUsecase) Delete(id string) (err error) {
	objectID, _ := primitive.ObjectIDFromHex(id)

	err = m.movieRepo.Delete(objectID)

	return err
}
