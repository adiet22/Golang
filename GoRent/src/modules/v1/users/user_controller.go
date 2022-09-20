package users

import (
	"encoding/json"
	"net/http"

	"github.com/adiet95/Golang/GoRent/src/database/orm/models"
	"github.com/adiet95/Golang/GoRent/src/helpers"
	"github.com/adiet95/Golang/GoRent/src/interfaces"
)

var respon helpers.Respons

type user_ctrl struct {
	svc interfaces.UserService
}

func NewCtrl(reps interfaces.UserService) *user_ctrl {
	return &user_ctrl{svc: reps}
}

func (re *user_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	//Header content type berguna untuk mengubah data yg dikirimkan menjadi json
	w.Header().Set("Content-type", "application/json")

	data, err := re.svc.GetAll()
	if err != nil {
		respon.Respon(w, 400, "", err, err)
	}

	respon.Respon(w, 200, "data found", data, err)
}

func (re *user_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.User
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		respon.Respon(w, 400, "", err, err)
	}

	data, err := re.svc.Add(&datas)
	if err != nil {
		respon.Respon(w, 400, "", data, err)
	}
	respon.Respon(w, 201, "data has been created", data, err)
}

func (re *user_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	val := r.URL.Query().Get("email")

	var datas models.User
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		respon.Respon(w, 400, "", err, err)
	}

	data, err := re.svc.Update(&datas, val)
	if err != nil {
		respon.Respon(w, 400, "data can't be update", data, err)
	}
	respon.Respon(w, 200, "user has been updated", data, err)
}

func (re *user_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	val := r.URL.Query().Get("email")

	data, err := re.svc.Delete(val)
	if err != nil {
		respon.Respon(w, 400, "data cant be delete", data, err)
	}
	respon.Respon(w, 200, "data has been deleted", data, err)
}
