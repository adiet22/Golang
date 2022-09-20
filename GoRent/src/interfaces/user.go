package interfaces

import (
	"github.com/adiet95/Golang/GoRent/src/database/orm/models"
)

type UserRepo interface {
	FindAll() (*models.Users, error)
	Save(data *models.User) (*models.User, error)
	UpdateUser(data *models.User, email string) (*models.User, error)
	DeleteUser(email string) (*models.User, error)
}

type UserService interface {
	GetAll() (*models.Users, error)
	Add(data *models.User) (*models.User, error)
	Update(data *models.User, email string) (*models.User, error)
	Delete(email string) (*models.User, error)
}
