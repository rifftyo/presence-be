package utils

import (
	"github.com/rifftyo/presence-be/internal/delivery/http/response"
	"github.com/rifftyo/presence-be/internal/entity"
)

func MapUserToUserResponse(u *entity.User) response.User {
	return response.User{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		Telephone:    u.Telephone,
		ImageProfile: u.ImageProfile,
		RoleId:       u.RoleId,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

func MapUserToUserDetailResponse(u *entity.User) response.UserDetailResponse {
    return response.UserDetailResponse{
        ID:         u.ID,
        Name:       u.Name,
        Photo:      u.ImageProfile,
        Role:       u.Role.Name,
        Email:      u.Email,
        Telephone:  u.Telephone,
        Department: u.Role.Department.Name,
    }
}