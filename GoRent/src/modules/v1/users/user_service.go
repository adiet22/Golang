package users

import (
	"github.com/adiet95/Golang/GoRent/src/database/orm/models"
	"github.com/adiet95/Golang/GoRent/src/interfaces"
)

type user_service struct {
	user_repo interfaces.UserRepo
}

func NewService(reps interfaces.UserRepo) *user_service {
	return &user_service{reps}
}

func (r *user_service) GetAll() (*models.Users, error) {
	data, err := r.user_repo.FindAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (re *user_service) Add(data *models.User) (*models.User, error) {
	data, err := re.user_repo.Save(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (re *user_service) Update(data *models.User, email string) (*models.User, error) {
	data, err := re.user_repo.UpdateUser(data, email)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (re *user_service) Delete(email string) (*models.User, error) {
	data, err := re.user_repo.DeleteUser(email)
	if err != nil {
		return nil, err
	}
	return data, nil
}
