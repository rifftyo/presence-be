package response

import "time"

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Telephone    string    `json:"telephone"`
	ImageProfile string    `json:"image_profile"`
	RoleId       string    `json:"role_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	User 	User `json:"user"`
}