package theater

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cinemanz/models"
)

// Repository represent the theater's repositories contract
type Repository interface {
	FetchAll(skip int64, limit int64) (theaters []models.Theater, total int64, err error)
	FetchByID(id primitive.ObjectID) (theater *models.Theater, err error)
	Store(t models.Theater) (err error)
	Update(theater *models.Theater) (err error)
	Delete(id primitive.ObjectID) (err error)
}
