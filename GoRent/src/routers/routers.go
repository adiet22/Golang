package routers

import (
	"errors"

	"github.com/adiet95/Golang/GoRent/src/database/orm"
	"github.com/adiet95/Golang/GoRent/src/modules/v1/histories"
	"github.com/adiet95/Golang/GoRent/src/modules/v1/users"
	vehicles "github.com/adiet95/Golang/GoRent/src/modules/v1/vehicles_1"
	"github.com/gorilla/mux"
)

func New() (*mux.Router, error) {

	mainRoute := mux.NewRouter()

	db, err := orm.New()
	if err != nil {
		return nil, errors.New("gagal init database")
	}

	users.New(mainRoute, db)
	vehicles.NewVe(mainRoute, db)
	histories.NewHis(mainRoute, db)

	return mainRoute, nil
}
