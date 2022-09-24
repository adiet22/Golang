package models

import "time"

//Struct ke data sesuai dengan data di database

type User struct {
	UserId      string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id,omitempty"`
	User_Name   string    `json:"user_name,omitempty"`
	Email       string    `json:"email,omitempty"`
	Role        string    `json:"role,omitempty"`
	Address     string    `json:"address,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	Password    string    `json:"Password,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Users []User
