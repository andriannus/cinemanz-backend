package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cinemanz/helper/response"
	"cinemanz/models"
	"cinemanz/user"
)

// UserHandler represent the httphandler for user
type UserHandler struct {
	UUsecase user.Usecase
}

// NewUserHandler will initialize the user/ resources endpoint
func NewUserHandler(route *chi.Mux, us user.Usecase) {
	handler := &UserHandler{
		UUsecase: us,
	}

	route.Route("/v1/user", func(route chi.Router) {
		route.Post("/login", handler.Login)
		route.Post("/register", handler.Register)
	})
}

// Login will do authentication
func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var dataLogin models.DataLogin
	json.NewDecoder(r.Body).Decode(&dataLogin)
	token, err := u.UUsecase.Login(dataLogin)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
	} else {
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"token": token,
		})
	}
}

// Register will register new user
func (u *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		ID: primitive.NewObjectID(),
	}
	json.NewDecoder(r.Body).Decode(&user)
	err := u.UUsecase.Register(user)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
	} else {
		response.JSON(w, http.StatusOK, map[string]string{
			"message": "Successfully Register New User",
		})
	}
}
