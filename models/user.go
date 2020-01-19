package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct for user
type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Privilege string             `json:"privilege" bson:"privilege"`
}

// DataLogin struct for data login
type DataLogin struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// MyClaims struct for JWT Claim
type MyClaims struct {
	jwt.StandardClaims
	Username  string `json:"username" bson:"username"`
	Privilege string `json:"privilege" bson:"privilege"`
}
