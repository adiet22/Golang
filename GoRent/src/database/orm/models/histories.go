package models

import "time"

type History struct {
	HistoryId    uint      `gorm:"type:uint;primaryKey;" json:"id,omitempty"`
	VehiclesId   uint      `json:"vehicles_id"`
	Vehicle      Vehicle   `gorm:"foreignKey:VehiclesId;" json:"vehicle_data"`
	HistoryName  string    `json:"history_name,omitempty"`
	Favorite     bool      `json:"favorite"`
	RentFrom     string    `json:"rent_from,omitempty"`
	RentTo       string    `json:"rent_to,omitempty"`
	StatusReturn string    `json:"status_return,omitempty"`
	Prepay       string    `json:"prepay,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Histories []History
