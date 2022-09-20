package vehicles

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewVe(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/vehi").Subrouter()
	repo := NewRepoVe(db)
	svc := NewServiceVe(repo)
	ctrl := NewCtrlVe(svc)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/", ctrl.Add).Methods("POST")
	route.HandleFunc("/", ctrl.Update).Methods("PUT")
	route.HandleFunc("/", ctrl.Delete).Methods("DELETE")
	route.HandleFunc("/search", ctrl.Search).Methods("POST")
	route.HandleFunc("/popular", ctrl.Popular).Methods("POST")

}
