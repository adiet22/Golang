package vehicles

import (
	"github.com/adiet95/Golang/GoRent/src/database/orm/models"
	"github.com/adiet95/Golang/GoRent/src/interfaces"
)

type vehi_service struct {
	vehi_repo interfaces.VehiRepo
}

func NewServiceVe(reps interfaces.VehiRepo) *vehi_service {
	return &vehi_service{reps}
}

func (r *vehi_service) GetAll() (*models.Vehicles, error) {
	data, err := r.vehi_repo.VeFindAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (re *vehi_service) Add(data *models.Vehicle) (*models.Vehicle, error) {
	data, err := re.vehi_repo.VeSave(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (re *vehi_service) Update(data *models.Vehicle, id int) (*models.Vehicle, error) {
	data, err := re.vehi_repo.VeUpdate(data, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (re *vehi_service) Delete(id int) (*models.Vehicle, error) {
	data, err := re.vehi_repo.VeDelete(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (re *vehi_service) Search(name string) (*models.Vehicles, error) {
	data, err := re.vehi_repo.FindByName(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (re *vehi_service) Popular(city string) (*models.Vehicles, error) {
	data, err := re.vehi_repo.PopularInCity(city)
	if err != nil {
		return nil, err
	}
	return data, nil
}
