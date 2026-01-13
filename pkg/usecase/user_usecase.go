package usecase

import (
	"github.com/rifftyo/presence-be/pkg/delivery/http/request"
	"github.com/rifftyo/presence-be/pkg/entity"
)

type UserUseCase interface {
	Register(req *request.RegisterUserRequest) (*entity.User, string, error)
	Login(req *request.LoginUserRequest) (*entity.User, string, error)
	GetUserByID(id string) (*entity.User, error)
}