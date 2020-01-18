package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cinemanz/config/middleware/auth"
	"cinemanz/helper/response"
	"cinemanz/models"
	"cinemanz/theater"
)

const (
	skipPerPage  = 0
	limitPerPage = 25
)

// TheaterHandler represent the httphandler for theater
type TheaterHandler struct {
	TUsecase theater.Usecase
}

// NewTheaterHandler will initialize the theater/ resources endpoint
func NewTheaterHandler(route *chi.Mux, us theater.Usecase) {
	handler := &TheaterHandler{
		TUsecase: us,
	}

	route.Route("/v1/theaters", func(route chi.Router) {
		route.Get("/", handler.FetchAll)
		route.Get("/{theaterID}", handler.FetchByID)

		route.With(auth.IsAuthenticated).Group(func(route chi.Router) {
			route.Post("/", handler.Store)
			route.Put("/{theaterID}", handler.Update)
			route.Delete("/{theaterID}", handler.Delete)
		})
	})
}

// FetchAll will fetch the all articles
func (t *TheaterHandler) FetchAll(w http.ResponseWriter, r *http.Request) {
	skip, err := strconv.ParseInt(r.URL.Query().Get("skip"), 10, 64)

	if err != nil {
		skip = skipPerPage
	}

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)

	if err != nil {
		limit = limitPerPage
	}

	theaters, total, err := t.TUsecase.FetchAll(skip, limit)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
	} else {
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"data":  theaters,
			"total": total,
		})
	}
}

// FetchByID will fetch theater by given id
func (t *TheaterHandler) FetchByID(w http.ResponseWriter, r *http.Request) {
	theaterID := chi.URLParam(r, "theaterID")

	theater, err := t.TUsecase.FetchByID(theaterID)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
	} else {
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"data": theater,
		})
	}
}

// Store will store the theater by given request body
func (t *TheaterHandler) Store(w http.ResponseWriter, r *http.Request) {
	theater := models.Theater{
		ID: primitive.NewObjectID(),
	}

	json.NewDecoder(r.Body).Decode(&theater)

	err := t.TUsecase.Store(theater)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
	} else {
		response.JSON(w, http.StatusOK, map[string]string{
			"message": "Successfully Add Theater",
		})
	}
}

// Update will update article by given id and request body
func (t *TheaterHandler) Update(w http.ResponseWriter, r *http.Request) {
	var theater *models.Theater

	theaterID := chi.URLParam(r, "theaterID")

	json.NewDecoder(r.Body).Decode(&theater)

	err := t.TUsecase.Update(theaterID, theater)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
	} else {
		response.JSON(w, http.StatusOK, map[string]string{
			"message": "Successfully Update Theater",
		})
	}
}

// Delete will delete theater by given id
func (t *TheaterHandler) Delete(w http.ResponseWriter, r *http.Request) {
	theaterID := chi.URLParam(r, "theaterID")

	err := t.TUsecase.Delete(theaterID)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
	} else {
		response.JSON(w, http.StatusOK, map[string]string{
			"message": "Successfully Delete Theater",
		})
	}
}
