package repository

import (
	"ChoTot/entity"
	"log"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetById(id string) (*entity.User, error)
	Update(user *entity.User) error
	Del(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) GetById(id string) (*entity.User, error) {
	user := &entity.User{}
	if err := ur.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) Update(info *entity.User) error {
	user := &entity.User{}
	if err := ur.db.Where("id = ?", info.Id).First(&user).Error; err != nil {
		log.Fatal("record not found")
		return err
	}

	//Update
	if err := ur.db.Updates(info).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) Del(id string) error {
	var user entity.User
	if err := ur.db.Where("id = ?", id).First(&user).Error; err != nil {
		log.Fatal("record not found")
		return err
	}

	//Delete
	if err := ur.db.Delete(user).Error; err != nil {
		log.Fatal("Fail to delete !")
		return err
	}
	return nil
}
