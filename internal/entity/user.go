package entity

import "time"

type User struct {
	ID           string `gorm:"primaryKey" json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password	 string `json:"password"`
	Telephone    string `json:"telephone"`
	ImageProfile string `json:"image_profile"`
	RoleId       string `json:"role_id"`
	Role 		 Role `gorm:"foreignKey:RoleId;references:ID"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt	 time.Time `json:"updated_at"`
}