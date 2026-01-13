package request

type RegisterUserRequest struct {
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required"`
	Password     string `json:"password" validate:"required,min=8"`
	Telephone    string `json:"telephone" validate:"required"`
	ImageProfile string `json:"image_profile,omitempty"`
	RoleId       string `json:"role_id" validate:"required"`
}