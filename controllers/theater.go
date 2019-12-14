package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"cinemanz/constants"
	"cinemanz/models"
	"cinemanz/utils"
)

// FetchTheaters return theaters
func FetchTheaters(w http.ResponseWriter, r *http.Request) {
	skip, err := strconv.ParseInt(r.URL.Query().Get("skip"), 10, 64)

	if err != nil {
		skip = constants.SkipPerPage
	}

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)

	if err != nil {
		limit = constants.LimitPerPage
	}

	theaters, total, err := models.FetchTheaters(skip, limit)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
	} else {
		utils.ResponseJSON(w, http.StatusOK, map[string]interface{}{
			"data":  theaters,
			"total": total,
		})
	}
}

// FetchTheater return theaters
func FetchTheater(w http.ResponseWriter, r *http.Request) {
	theaterID := chi.URLParam(r, "theaterID")
	objectID, _ := primitive.ObjectIDFromHex(theaterID)

	theater, err := models.FetchTheater(objectID)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
	} else {
		utils.ResponseJSON(w, http.StatusOK, map[string]interface{}{
			"data": theater,
		})
	}
}

// AddTheater add a new theaters
func AddTheater(w http.ResponseWriter, r *http.Request) {
	theater := models.Theater{
		ID: primitive.NewObjectID(),
	}

	json.NewDecoder(r.Body).Decode(&theater)

	err := models.AddTheater(theater)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
	} else {
		utils.ResponseJSON(w, http.StatusOK, map[string]string{
			"message": "Successfully Add Theater",
		})
	}
}

// UpdateTheater update a theaters
func UpdateTheater(w http.ResponseWriter, r *http.Request) {
	theaterID := chi.URLParam(r, "theaterID")
	objectID, _ := primitive.ObjectIDFromHex(theaterID)

	theater := models.Theater{
		ID: objectID,
	}

	json.NewDecoder(r.Body).Decode(&theater)

	err := models.UpdateTheater(objectID, theater)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
	} else {
		utils.ResponseJSON(w, http.StatusOK, map[string]string{
			"message": "Successfully Update Theater",
		})
	}
}

// DeleteTheater delete a theaters
func DeleteTheater(w http.ResponseWriter, r *http.Request) {
	theaterID := chi.URLParam(r, "theaterID")
	objectID, _ := primitive.ObjectIDFromHex(theaterID)

	err := models.DeleteTheater(objectID)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
	} else {
		utils.ResponseJSON(w, http.StatusOK, map[string]string{
			"message": "Successfully Delete Theater",
		})
	}
}
