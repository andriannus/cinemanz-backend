package user

import (
	"cinemanz/models"
)

// Usecase represent the user's usecases
type Usecase interface {
	Login(data models.DataLogin) (token string, err error)
	Register(user models.User) (err error)
}
