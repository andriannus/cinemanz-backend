package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"cinemanz/middleware"

	_movieHttpDeliver "cinemanz/movie/delivery/http"
	_movieRepo "cinemanz/movie/repository"
	_movieUcase "cinemanz/movie/usecase"

	_theaterHttpDeliver "cinemanz/theater/delivery/http"
	_theaterRepo "cinemanz/theater/repository"
	_theaterUcase "cinemanz/theater/usecase"

	_userHttpDeliver "cinemanz/user/delivery/http"
	_userRepo "cinemanz/user/repository"
	_userUcase "cinemanz/user/usecase"
)

func init() {
	viper.SetConfigFile(`config.json`)

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	dbConn, err := dbSetup()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Berhasil terhubung ke database")

	route := chi.NewRouter()
	middL := middleware.InitMiddleware()

	route.Use(middL.CORS().Handler)

	movieRepo := _movieRepo.NewMongoMovieRepository(dbConn)
	movieUcase := _movieUcase.NewMovieUsecase(movieRepo)
	_movieHttpDeliver.NewMovieHandler(route, movieUcase)

	theaterRepo := _theaterRepo.NewMongoTheaterRepository(dbConn)
	theaterUcase := _theaterUcase.NewTheaterUsecase(theaterRepo)
	_theaterHttpDeliver.NewTheaterHandler(route, theaterUcase)

	userRepo := _userRepo.NewmongoUserRepository(dbConn)
	userUcase := _userUcase.NewUserUsecase(userRepo)
	_userHttpDeliver.NewUserHandler(route, userUcase)

	address := fmt.Sprintf(":%d", viper.GetInt(`app.port`))

	err = http.ListenAndServe(address, route)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func dbSetup() (*mongo.Database, error) {
	dbURI := viper.GetString(`database.mongo.uri`)
	dbName := viper.GetString(`database.mongo.name`)

	options := options.Client().ApplyURI(dbURI)
	client, err := mongo.NewClient(options)

	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())

	if err != nil {
		return nil, err
	}

	return client.Database(dbName), nil
}
