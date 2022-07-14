package repository

import (
	"ChoTot/entity"

	"gorm.io/gorm"
)

type ICateRepository interface {
	Create(input *entity.SubCategories) error
	GetAll() ([]entity.SubCategories, error)
	GetById(id string) (*entity.SubCategories, error)
	Update(cate *entity.SubCategories) error
	Del(id string) error
}

type cateRepository struct {
	db *gorm.DB
}

func NewCateRepository(db *gorm.DB) ICateRepository {
	return &cateRepository{db: db}
}

func (cr *cateRepository) GetAll() ([]entity.SubCategories, error) {
	var Cates []entity.SubCategories
	err := cr.db.Find(&Cates).Error
	return Cates, err
}

func (cr *cateRepository) Create(input *entity.SubCategories) error {
	
	return cr.db.Create(input).Error
}

func (cr *cateRepository) GetById(id string) (*entity.SubCategories, error) {
	var Cate entity.SubCategories
	err := cr.db.Where("id = ?", id).First(&Cate).Error
	return &Cate, err
}

func (cr *cateRepository) Update(info *entity.SubCategories) error {
	return cr.db.Updates(info).Error
}

func (cr *cateRepository) Del(id string) error {
	cate, _ := cr.GetById(id)
	return cr.db.Delete(cate).Error
}
