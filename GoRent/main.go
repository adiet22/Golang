package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adiet95/Golang/GoRent/src/helpers"
	"github.com/adiet95/Golang/GoRent/src/routers"
)

var godotenv helpers.Godot

func main() {
	mainRoute, err := routers.New()
	if err != nil {
		log.Fatal(err)
	}

	PORT := godotenv.GoDotEnv("PORT")
	fmt.Println("Service run on port", PORT)
	http.ListenAndServe(PORT, mainRoute)

}
