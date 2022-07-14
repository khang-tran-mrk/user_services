package service

import (
	"ChoTot/entity"
	"ChoTot/repository"
)

type IUserService interface {
	GetById(id string) (*entity.User, error)
	Update(info *entity.User) error
	Del(id string) error
}

type userService struct {
	userRepo repository.IUserRepository
}

func NewUserService(userRepo repository.IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

func (svc *userService) GetById(id string) (*entity.User, error) {
	return svc.userRepo.GetById(id)
}

func (svc *userService) Update(info *entity.User) error {
	return svc.userRepo.Update(info)
}

func (svc *userService) Del(id string) error {
	return svc.userRepo.Del(id)
}
