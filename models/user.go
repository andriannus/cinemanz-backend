package models

import (
	"context"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"cinemanz/databases"
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

// Register to create new user
func Register(u User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	_, err = databases.Mongo.Collection("user").InsertOne(
		context.Background(),
		bson.M{
			"username":  u.Username,
			"password":  string(hashedPassword),
			"privilege": "admin",
		},
	)

	return err
}

// Login return token
func Login(dataLogin DataLogin) (user *User, err error) {
	err = databases.Mongo.Collection("user").FindOne(
		context.Background(),
		bson.M{"username": dataLogin.Username},
	).Decode(&user)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dataLogin.Password))

	if err != nil {
		return nil, err
	}

	return user, nil
}
