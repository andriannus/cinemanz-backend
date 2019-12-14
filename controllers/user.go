package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cinemanz/models"
	"cinemanz/utils"
)

// Register for create user
func Register(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		ID: primitive.NewObjectID(),
	}

	json.NewDecoder(r.Body).Decode(&user)

	err := models.Register(user)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
	} else {
		utils.ResponseJSON(w, http.StatusOK, map[string]string{
			"message": "Successfully Register New User",
		})
	}
}

// Login check auth
func Login(w http.ResponseWriter, r *http.Request) {
	var dataLogin models.DataLogin

	json.NewDecoder(r.Body).Decode(&dataLogin)

	user, err := models.Login(dataLogin)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	claims := models.MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour).Unix(),
		},
		Username:  user.Username,
		Privilege: user.Privilege,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("c0b4d1b4c4"))

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
	} else {
		utils.ResponseJSON(w, http.StatusOK, map[string]interface{}{
			"token": signedToken,
		})
	}
}
