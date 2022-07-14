package service

import (
	"ChoTot/entity"
	"ChoTot/repository"
)

type ICateService interface {
	Create(input *entity.SubCategories) error
	GetAll() ([]entity.SubCategories, error)
	GetById(id string) (*entity.SubCategories, error)
	Update(cate *entity.SubCategories) error
	Del(id string) error
}

type cateService struct {
	cateRepo repository.ICateRepository
}

func NewCateService(cateRepo repository.ICateRepository) ICateService {
	return &cateService{cateRepo: cateRepo}
}

func (svc *cateService) GetAll() ([]entity.SubCategories, error) {
	return svc.cateRepo.GetAll()
}

func (svc *cateService) Create(input *entity.SubCategories) error {
	return svc.cateRepo.Create(input)
}

func (svc *cateService) GetById(id string) (*entity.SubCategories, error) {
	return svc.cateRepo.GetById(id)
}

func (svc *cateService) Update(info *entity.SubCategories) error {
	return svc.cateRepo.Update(info)
}

func (svc *cateService) Del(id string) error {
	return svc.cateRepo.Del(id)
}
