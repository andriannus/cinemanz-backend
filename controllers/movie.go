package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cinemanz/constants"
	"cinemanz/models"
	"cinemanz/utils"
)

// FetchMovies return movies
func FetchMovies(w http.ResponseWriter, r *http.Request) {
	skip, err := strconv.ParseInt(r.URL.Query().Get("skip"), 10, 64)

	if err != nil {
		skip = constants.SkipPerPage
	}

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)

	if err != nil {
		limit = constants.LimitPerPage
	}

	movies, total, err := models.FetchMovies(skip, limit)

	if err != nil {
		fmt.Print(err)
		utils.ResponseError(w, 500, "Something wrong")
		return
	}

	utils.ResponseJSON(w, 200, map[string]interface{}{
		"data":  movies,
		"total": total,
	})
}

// FetchMovie return movie
func FetchMovie(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "movieID")
	objectID, _ := primitive.ObjectIDFromHex(movieID)

	movie, err := models.FetchMovie(objectID)

	if err != nil {
		utils.ResponseError(w, 500, "Something wrong")
	} else {
		utils.ResponseJSON(w, 200, map[string]interface{}{
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
		utils.ResponseError(w, 500, "Something wrong")
	} else {
		utils.ResponseJSON(w, 200, map[string]string{"message": "Successfully Add Movie"})
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
		utils.ResponseError(w, 500, "Something wrong")
	} else {
		utils.ResponseJSON(w, 200, map[string]string{"message": "Successfully Update Movie"})
	}
}

// DeleteMovie delete a movie
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "movieID")
	objectID, _ := primitive.ObjectIDFromHex(movieID)

	err := models.DeleteMovie(objectID)

	if err != nil {
		utils.ResponseError(w, 500, "Something wrong")
	} else {
		utils.ResponseJSON(w, 200, map[string]string{"message": "Successfully Delete Movie"})
	}
}
