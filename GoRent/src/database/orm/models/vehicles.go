package models

import "time"

//Struct ke data sesuai dengan data di database

type Vehicle struct {
	VehicleId    uint          `gorm:"type:uint;primaryKey;" json:"vehicles_id"`
	Vehicleitem  []Vehicleitem `gorm:"many2many:vehicle_vehicleitem"`
	UserId       string        `json:"user_id"`
	User         User          `gorm:"type:uuid;foreignKey:UserId;references:user_id"`
	Vehicle_Name string        `json:"vehicle_name,omitempty"`
	City         string        `json:"city,omitempty"`
	Status       string        `json:"status,omitempty"`
	Capacity     string        `json:"capacity,omitempty"`
	Type         string        `json:"type,omitempty"`
	Description  string        `json:"description,omitempty"`
	Price        string        `json:"price,omitempty"`
	Popular      int           `json:"popular,omitempty"`
	CreatedAt    time.Time     `gorm:"default:now(); not null" json:"created_at"`
	UpdatedAt    time.Time     `gorm:"default:now(); not null" json:"updated_at"`
}

type Vehicleitem struct {
	VehicleitemId uint      `gorm:"type:uint;primaryKey;" json:"vehicles_id"`
	Vehicle       []Vehicle `gorm:"many2many:vehicle_vehicleitem"`
	HistoryId     uint      `json:"history_id"`
	History       History   `gorm:"type:uint;foreignKey:HistoryId;references:history_id"`
	Amount        int       `json:"amount,omitempty"`
	CreatedAt     time.Time `gorm:"default:now(); not null" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:now(); not null" json:"updated_at"`
}

type Vehicles []Vehicle
