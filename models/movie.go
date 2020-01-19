package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Movie struct for movie
type Movie struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Title      string             `json:"title" bson:"title"`
	Genre      []string           `json:"genre" bson:"genre"`
	Minutes    int                `json:"minutes" bson:"minutes"`
	Synopsis   string             `json:"synopsis" bson:"synopsis"`
	Producer   []string           `json:"producer" bson:"producer"`
	Production string             `json:"production" bson:"production"`
	Director   string             `json:"director" bson:"director"`
	Writer     string             `json:"writer" bson:"writer"`
	Cast       []string           `json:"cast" bson:"cast"`
	Start      string             `json:"start" bson:"start"`
	End        string             `json:"end" bson:"end"`
	Rate       float32            `json:"rate" bson:"rate"`
	Theater    []string           `json:"theater" bson:"theater"`
	Poster     string             `json:"poster" bson:"poster"`
}
