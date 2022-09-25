package users

import (
	"github.com/adiet95/Golang/GoRent/src/database/orm/models"
	"github.com/adiet95/Golang/GoRent/src/helpers"
	"github.com/adiet95/Golang/GoRent/src/interfaces"
)

type user_service struct {
	user_repo interfaces.UserRepo
}

func NewService(reps interfaces.UserRepo) *user_service {
	return &user_service{reps}
}

func (re *user_service) Add(data *models.User) *helpers.Response {

	hassPass, err := helpers.HashPassword(data.Password)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}

	data.Password = hassPass
	result, err := re.user_repo.Save(data)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(result, 200, false)
}

func (re *user_service) Update(data *models.User, email string) *helpers.Response {
	hassPass, err := helpers.HashPassword(data.Password)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	data.Email = email
	data.Role = "user"
	data.Password = hassPass

	result, err := re.user_repo.UpdateUser(data, email)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(result, 200, false)
}

func (re *user_service) Delete(email string) *helpers.Response {
	data, err := re.user_repo.DeleteUser(email)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(data, 200, false)
}

func (re *user_service) FindEmail(email string) *helpers.Response {
	data, err := re.user_repo.FindByEmail(email)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	if data.Role == "admin" {
		datas, err := re.user_repo.FindAll()
		if err != nil {
			return helpers.New(err.Error(), 400, true)
		}
		return helpers.New(datas, 200, false)
	}
	return helpers.New(data, 200, false)
}
