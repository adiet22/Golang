package auth

import (
	"encoding/json"
	"net/http"

	"github.com/adiet95/Golang/GoRent/src/database/orm/models"
	"github.com/adiet95/Golang/GoRent/src/helpers"
	"github.com/adiet95/Golang/GoRent/src/interfaces"
)

type auth_ctrl struct {
	user_repo interfaces.AuthService
}

func NewCtrl(reps interfaces.AuthService) *auth_ctrl {
	return &auth_ctrl{reps}
}

func (a auth_ctrl) SignIn(w http.ResponseWriter, r *http.Request) {
	var data models.User

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helpers.New(err.Error(), 401, true)
		return
	}
	a.user_repo.Login(data).Send(w)
}