package vehicles

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/adiet95/Golang/GoRent/src/database/orm/models"
	"github.com/adiet95/Golang/GoRent/src/helpers"
	"github.com/adiet95/Golang/GoRent/src/interfaces"
)

var respon helpers.Respons

type vehi_ctrl struct {
	svc interfaces.VehiService
}

func NewCtrlVe(reps interfaces.VehiService) *vehi_ctrl {
	return &vehi_ctrl{svc: reps}
}

func (re *vehi_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	data, err := re.svc.GetAll()
	if err != nil {
		respon.Respon(w, 400, "", err, err)
	}
	respon.Respon(w, 200, "data found", data, err)
}

func (re *vehi_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var datas models.Vehicle
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

func (re *vehi_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	val := r.URL.Query().Get("id")
	v, _ := strconv.Atoi(val)

	var datas models.Vehicle
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		respon.Respon(w, 400, "", err, err)
	}

	data, err := re.svc.Update(&datas, v)
	if err != nil {
		respon.Respon(w, 400, "data can't be update", data, err)
	}
	respon.Respon(w, 200, "user has been updated", data, err)
}

func (re *vehi_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	val := r.URL.Query().Get("id")
	v, _ := strconv.Atoi(val)

	data, err := re.svc.Delete(v)
	if err != nil {
		respon.Respon(w, 400, "data cant be delete", data, err)
	}
	respon.Respon(w, 200, "data has been deleted", data, err)
}

func (re *vehi_ctrl) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	val := r.URL.Query().Get("name")
	v := strings.ToLower(val)

	data, err := re.svc.Search(v)
	if err != nil {
		respon.Respon(w, 400, "data not found", data, err)
	}
	respon.Respon(w, 200, "data found", data, err)
}

func (re *vehi_ctrl) Popular(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	val := r.URL.Query().Get("city")
	v := strings.ToLower(val)

	data, err := re.svc.Popular(v)
	if err != nil {
		respon.Respon(w, 400, "data not found", data, err)
	}
	respon.Respon(w, 200, "data found", data, err)
}
