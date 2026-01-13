package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/rifftyo/presence-be/pkg/delivery/http/request"
	"github.com/rifftyo/presence-be/pkg/entity"
	"github.com/rifftyo/presence-be/pkg/repository"
	"github.com/rifftyo/presence-be/utils"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{userRepo}
}

func (u *userUseCase) Register(req *request.RegisterUserRequest) (*entity.User, string, error) {
	exist, _ := u.userRepo.FindByEmail(req.Email)
	if exist != nil {
		return nil, "", errors.New("email alredy in use")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := &entity.User{
		ID:           uuid.New().String(),
		Name:         req.Name,
		Email:        req.Email,
		Password:     string(hashed),
		Telephone:    req.Telephone,
		ImageProfile: req.ImageProfile,
		RoleId:       req.RoleId,
	}

	err := u.userRepo.Create(user)
	if err != nil {
		return nil, "", err
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (u *userUseCase) Login(req *request.LoginUserRequest) (*entity.User, string, error) {
	user, err := u.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, "", errors.New("email not registered")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, "", errors.New("wrong password")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (u *userUseCase) GetUserByID(id string) (*entity.User, error) {
	return u.userRepo.GetById(id)
}