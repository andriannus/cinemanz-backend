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

	return
}

func (m *movieUsecase) FetchNowPlaying(skip int64, limit int64, date string) (movies []models.Movie, total int64, err error) {
	movies, total, err = m.movieRepo.FetchNowPlaying(skip, limit, date)

	return
}

func (m *movieUsecase) FetchUpcoming(skip int64, limit int64, date string) (movies []models.Movie, total int64, err error) {
	movies, total, err = m.movieRepo.FetchUpcoming(skip, limit, date)

	return
}

func (m *movieUsecase) FetchByID(id string) (movie *models.Movie, err error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	movie, err = m.movieRepo.FetchByID(objectID)

	return
}

func (m *movieUsecase) Store(movie models.Movie) error {
	err := m.movieRepo.Store(movie)

	return err
}

func (m *movieUsecase) Update(id string, movie *models.Movie) error {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	movie.ID = objectID
	err = m.movieRepo.Update(movie)

	return err
}

func (m *movieUsecase) Delete(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	err = m.movieRepo.Delete(objectID)

	return err
}
