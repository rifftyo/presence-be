package repository

import "github.com/rifftyo/presence-be/pkg/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	GetById(id string) (*entity.User, error)
}