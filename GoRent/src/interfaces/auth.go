package interfaces

import (
	"github.com/adiet95/Golang/GoRent/src/database/orm/models"
	"github.com/adiet95/Golang/GoRent/src/helpers"
)

type AuthService interface {
	Login(body models.User) *helpers.Response
	Register(body *models.User) *helpers.Response
}
