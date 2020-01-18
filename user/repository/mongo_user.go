package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"cinemanz/models"
	"cinemanz/user"
)

type mongoUserRepo struct {
	DB *mongo.Database
}

// NewmongoUserRepository will create an object that represent the user.Repository interface
func NewmongoUserRepository(db *mongo.Database) user.Repository {
	return &mongoUserRepo{
		DB: db,
	}
}

func (m *mongoUserRepo) Login(dataLogin models.DataLogin) (user *models.User, err error) {
	err = m.DB.Collection("user").FindOne(
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

// Register to create new user
func (m *mongoUserRepo) Register(u models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	_, err = m.DB.Collection("user").InsertOne(
		context.Background(),
		bson.M{
			"username":  u.Username,
			"password":  string(hashedPassword),
			"privilege": "admin",
		},
	)

	return err
}
