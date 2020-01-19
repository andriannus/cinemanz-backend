package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Theater struct for theater
type Theater struct {
	ID        primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Name      string             `json:"name" bson:"name"`
	Address   string             `json:"address" bson:"address"`
	Telephone string             `json:"telephone" bson:"telephone"`
}
