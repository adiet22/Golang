package models

import "time"

//Struct ke data sesuai dengan data di database

type Vehicle struct {
	VehiclesId   uint   `gorm:"type:uint;primaryKey;" json:"vehicles_id,omitempty"`
	Vehicle_Name string `json:"vehicle_name,omitempty"`
	City         string `json:"city,omitempty"`
	Status       string `json:"status,omitempty"`
	Capacity     string `json:"capacity,omitempty"`
	Type         string `json:"type,omitempty"`
	Description  string `json:"description,omitempty"`
	Price        string `json:"price,omitempty"`
	Amount       int    `json:"amount,omitempty"`
	Popular      int    `json:"popular,omitempty"`
	// HistoryId    uint      `json:"history_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Vehicles []Vehicle
