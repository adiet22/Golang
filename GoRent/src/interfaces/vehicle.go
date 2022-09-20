package interfaces

import (
	"github.com/adiet95/Golang/GoRent/src/database/orm/models"
)

type VehiRepo interface {
	VeFindAll() (*models.Vehicles, error)
	VeSave(data *models.Vehicle) (*models.Vehicle, error)
	VeUpdate(data *models.Vehicle, id int) (*models.Vehicle, error)
	VeDelete(id int) (*models.Vehicle, error)
	FindByName(name string) (*models.Vehicles, error)
	PopularInCity(city string) (*models.Vehicles, error)
}

type VehiService interface {
	GetAll() (*models.Vehicles, error)
	Add(data *models.Vehicle) (*models.Vehicle, error)
	Update(data *models.Vehicle, id int) (*models.Vehicle, error)
	Delete(id int) (*models.Vehicle, error)
	Search(name string) (*models.Vehicles, error)
	Popular(city string) (*models.Vehicles, error)
}
