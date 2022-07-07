package service

import (
	"ChoTot/entity"
	"ChoTot/repository"
)

type IUserService interface {
	UserProfile(id int) (*entity.User, error)
}

type userService struct {
	userRepo repository.IUserRepository
}

func NewUserService(userRepo repository.IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

func (svc *userService) UserProfile(id int) (*entity.User, error) {
	return svc.userRepo.UserProfile(id)
}
