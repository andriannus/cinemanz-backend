package usecase

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cinemanz/models"
	"cinemanz/theater"
)

type theaterUsecase struct {
	theaterRepo theater.Repository
}

// NewTheaterUsecase will create an object that represent the theater.Usecase interface
func NewTheaterUsecase(t theater.Repository) theater.Usecase {
	return &theaterUsecase{
		theaterRepo: t,
	}
}

func (t *theaterUsecase) FetchAll(skip int64, limit int64) (theaters []models.Theater, total int64, err error) {
	theaters, total, err = t.theaterRepo.FetchAll(skip, limit)

	return
}

func (t *theaterUsecase) FetchByID(id string) (theater *models.Theater, err error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	theater, err = t.theaterRepo.FetchByID(objectID)

	return
}

func (t *theaterUsecase) Store(theater models.Theater) error {
	err := t.theaterRepo.Store(theater)

	return err
}

func (t *theaterUsecase) Update(id string, theater *models.Theater) error {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	theater.ID = objectID
	err = t.theaterRepo.Update(theater)

	return err
}

func (t *theaterUsecase) Delete(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	err = t.theaterRepo.Delete(objectID)

	return err
}
