package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cinemanz/helper/response"
	"cinemanz/middleware"
	"cinemanz/models"
	"cinemanz/movie"
)

const (
	dateFormat   = "2006-01-02"
	limitPerPage = 25
	skipPerPage  = 0
)

// MovieHandler represent the httphandler for movie
type MovieHandler struct {
	MUsecase movie.Usecase
}

// NewMovieHandler will initialize the movie/ resources endpoint
func NewMovieHandler(route *chi.Mux, us movie.Usecase) {
	handler := &MovieHandler{
		MUsecase: us,
	}

	middL := middleware.InitMiddleware()

	route.Route("/v1/movies", func(route chi.Router) {
		route.Get("/", handler.FetchAll)
		route.Get("/now-playing", handler.FetchNowPlaying)
		route.Get("/upcoming", handler.FetchUpcoming)
		route.Get("/{movieID}", handler.FetchByID)

		route.With(middL.IsAuthenticated).Group(func(route chi.Router) {
			route.Post("/", handler.Store)
			route.Put("/{movieID}", handler.Update)
			route.Delete("/{movieID}", handler.Delete)
		})
	})
}

// FetchAll will fetch the all articles
func (m *MovieHandler) FetchAll(w http.ResponseWriter, r *http.Request) {
	skip, limit := getSkipAndLimit(r)

	movies, total, err := m.MUsecase.FetchAll(skip, limit)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
	} else {
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"data":  movies,
			"total": total,
		})
	}
}

func (m *MovieHandler) FetchNowPlaying(w http.ResponseWriter, r *http.Request) {
	skip, limit := getSkipAndLimit(r)

	selectedDate := r.URL.Query().Get("date")

	if selectedDate == "" {
		selectedDate = time.Now().Format(dateFormat)
	}

	movies, total, err := m.MUsecase.FetchNowPlaying(skip, limit, selectedDate)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
	} else {
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"data":  movies,
			"total": total,
		})
	}
}

func (m *MovieHandler) FetchUpcoming(w http.ResponseWriter, r *http.Request) {
	skip, limit := getSkipAndLimit(r)
	selectedDate := r.URL.Query().Get("date")

	if selectedDate == "" {
		selectedDate = time.Now().Format(dateFormat)
	}

	movies, total, err := m.MUsecase.FetchUpcoming(skip, limit, selectedDate)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
	} else {
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"data":  movies,
			"total": total,
		})
	}
}

// FetchByID will fetch movie by given id
func (m *MovieHandler) FetchByID(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "movieID")
	movie, err := m.MUsecase.FetchByID(movieID)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
	} else {
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"data": movie,
		})
	}
}

// Store will store the movie by given request body
func (m *MovieHandler) Store(w http.ResponseWriter, r *http.Request) {
	movie := models.Movie{
		ID: primitive.NewObjectID(),
	}

	json.NewDecoder(r.Body).Decode(&movie)

	err := m.MUsecase.Store(movie)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
	} else {
		response.JSON(w, http.StatusOK, map[string]string{
			"message": "Successfully Add Movie",
		})
	}
}

// Update will update article by given id and request body
func (m *MovieHandler) Update(w http.ResponseWriter, r *http.Request) {
	var movie *models.Movie

	movieID := chi.URLParam(r, "movieID")

	json.NewDecoder(r.Body).Decode(&movie)

	err := m.MUsecase.Update(movieID, movie)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
	} else {
		response.JSON(w, http.StatusOK, map[string]string{
			"message": "Successfully Update Movie",
		})
	}
}

// Delete will delete movie by given id
func (m *MovieHandler) Delete(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "movieID")

	err := m.MUsecase.Delete(movieID)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
	} else {
		response.JSON(w, http.StatusOK, map[string]string{
			"message": "Successfully Delete Movie",
		})
	}
}

func getSkipAndLimit(r *http.Request) (skip int64, limit int64) {
	skip, err := strconv.ParseInt(r.URL.Query().Get("skip"), 10, 64)

	if err != nil {
		skip = skipPerPage
	}

	limit, err = strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)

	if err != nil {
		limit = limitPerPage
	}

	return skip, limit
}
