package usecase

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"cinemanz/models"
	"cinemanz/user"
)

const (
	secretKey = "c0b4d1b4c4"
)

type userUsecase struct {
	userRepo user.Repository
}

// NewUserUsecase will create an object that represent the user.Usecase interface
func NewUserUsecase(u user.Repository) user.Usecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (u *userUsecase) Login(dataLogin models.DataLogin) (signedToken string, err error) {
	user, err := u.userRepo.Login(dataLogin)

	if err != nil {
		return "", err
	}

	expiredTime := time.Now().Add(time.Duration(1) * time.Hour).Unix()

	claims := models.MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime,
		},
		Username:  user.Username,
		Privilege: user.Privilege,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(secretKey))

	return signedToken, err
}

func (u *userUsecase) Register(user models.User) (err error) {
	err = u.userRepo.Register(user)

	return err
}
