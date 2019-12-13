package main

import (
	"fmt"
	"net/http"

	"cinemanz/constants"
	"cinemanz/databases"
	"cinemanz/routes"
)

func main() {
	var err error

	databases.Mongo, err = databases.MongoSetup()

	address := fmt.Sprintf(":%d", constants.Port)

	err = http.ListenAndServe(address, routes.Routes())

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Server started at %s\n", address)
}
