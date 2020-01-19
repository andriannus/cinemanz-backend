package user

import (
	"cinemanz/models"
)

// Repository represent the user's repositories contract
type Repository interface {
	Login(dataLogin models.DataLogin) (user *models.User, err error)
	Register(user models.User) (err error)
}
