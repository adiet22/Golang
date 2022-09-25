package auth

import (
	"github.com/adiet95/Golang/GoRent/src/modules/v1/users"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewAu(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/auth").Subrouter()
	repo := users.NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.SignIn).Methods("POST")
	route.HandleFunc("/register", ctrl.Register).Methods("POST")

}
