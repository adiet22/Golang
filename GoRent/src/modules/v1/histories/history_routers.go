package histories

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewHis(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/history").Subrouter()
	repo := NewRepoHis(db)
	svc := NewServiceHis(repo)
	ctrl := NewCtrlHis(svc)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/", ctrl.Add).Methods("POST")
	route.HandleFunc("/", ctrl.Update).Methods("PUT")
	route.HandleFunc("/", ctrl.Delete).Methods("DELETE")
	route.HandleFunc("/search", ctrl.Search).Methods("POST")
	route.HandleFunc("/favo", ctrl.Favorite).Methods("GET")

}
