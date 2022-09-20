package interfaces

import (
	"github.com/adiet95/Golang/GoRent/src/database/orm/models"
)

type HistoryRepo interface {
	HisFindAll() (*models.Histories, error)
	HisSave(data *models.History) (*models.History, error)
	HisUpdate(data *models.History, id int) (*models.History, error)
	HisDelete(id int) (*models.History, error)
	FindByName(name string) (*models.Histories, error)
	GetFavo() (*models.Histories, error)
}

type HistoryService interface {
	GetAll() (*models.Histories, error)
	Add(data *models.History) (*models.History, error)
	Update(data *models.History, id int) (*models.History, error)
	Delete(id int) (*models.History, error)
	Search(name string) (*models.Histories, error)
	Favorite() (*models.Histories, error)
}
