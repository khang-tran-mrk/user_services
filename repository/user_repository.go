package repository

import (
	"ChoTot/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	UserProfile(id int) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) UserProfile(id int) (*entity.User, error) {
	user := &entity.User{}
	if err := ur.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
