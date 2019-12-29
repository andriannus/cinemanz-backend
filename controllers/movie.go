package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cinemanz/constants"
	"cinemanz/models"
	"cinemanz/utils"
)

// FetchMovies return movies
func FetchMovies(w http.ResponseWriter, r *http.Request) {
	skip, limit := utils.GetSkipAndLimit(r)

	movies, total, err := models.FetchMovies(skip, limit)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
	} else {
		utils.ResponseJSON(w, http.StatusOK, map[string]interface{}{
			"data":  movies,
			"total": total,
		})
	}
}

// FetchNowPlayingMovies return movies
func FetchNowPlayingMovies(w http.ResponseWriter, r *http.Request) {
	skip, limit := utils.GetSkipAndLimit(r)

	selectedDate := r.URL.Query().Get("date")

	if selectedDate == "" {
		selectedDate = time.Now().Format(constants.DateFormat)
	}

	movies, total, err := models.FetchNowPlayingMovies(skip, limit, selectedDate)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
	} else {
		utils.ResponseJSON(w, http.StatusOK, map[string]interface{}{
			"data":  movies,
			"total": total,
		})
	}
}

// FetchUpcomingMovies return movies
func FetchUpcomingMovies(w http.ResponseWriter, r *http.Request) {
	skip, limit := utils.GetSkipAndLimit(r)

	selectedDate := r.URL.Query().Get("date")

	if selectedDate == "" {
		selectedDate = time.Now().Format(constants.DateFormat)
	}

	movies, total, err := models.FetchUpcomingMovies(skip, limit, selectedDate)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
	} else {
		utils.ResponseJSON(w, http.StatusOK, map[string]interface{}{
			"data":  movies,
			"total": total,
		})
	}
}

// FetchMovie return movie
func FetchMovie(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "movieID")
	objectID, _ := primitive.ObjectIDFromHex(movieID)

	movie, err := models.FetchMovie(objectID)

	if err != nil {
		utils.ResponseError(w, http.StatusBadGateway, err.Error())
	} else {
		utils.ResponseJSON(w, http.StatusOK, map[string]interface{}{
			"data": movie,
		})
	}
}

// AddMovie add a new movie
func AddMovie(w http.ResponseWriter, r *http.Request) {
	movie := models.Movie{
		ID: primitive.NewObjectID(),
	}

	json.NewDecoder(r.Body).Decode(&movie)

	err := models.AddMovie(movie)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
	} else {
		utils.ResponseJSON(w, http.StatusOK, map[string]string{
			"message": "Successfully Add Movie",
		})
	}
}

// UpdateMovie update a movie
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "movieID")
	objectID, _ := primitive.ObjectIDFromHex(movieID)

	movie := models.Movie{
		ID: objectID,
	}

	json.NewDecoder(r.Body).Decode(&movie)

	err := models.UpdateMovie(objectID, movie)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
	} else {
		utils.ResponseJSON(w, http.StatusOK, map[string]string{
			"message": "Successfully Update Movie",
		})
	}
}

// DeleteMovie delete a movie
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "movieID")
	objectID, _ := primitive.ObjectIDFromHex(movieID)

	err := models.DeleteMovie(objectID)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
	} else {
		utils.ResponseJSON(w, http.StatusOK, map[string]string{
			"message": "Successfully Delete Movie",
		})
	}
}
