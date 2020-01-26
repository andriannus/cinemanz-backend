package theater

import (
	"cinemanz/models"
)

// Usecase represent the theater's usecases
type Usecase interface {
	FetchAll(skip int64, limit int64) (theaters []models.Theater, total int64, err error)
	FetchByID(id string) (theater *models.Theater, err error)
	Store(theater models.Theater) error
	Update(id string, theater *models.Theater) error
	Delete(id string) error
}
