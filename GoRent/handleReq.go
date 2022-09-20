package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func reqs() {
	mainRoute := mux.NewRouter()

	fmt.Println("app running on port 8080")
	err := http.ListenAndServe(":8080", mainRoute)

	if err != nil {
		panic(err)
	}
}
