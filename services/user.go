package services

import (
	"errors"
	"go-backend/models"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (usv *UserService) Create(user models.User) (err error) {
	result := usv.DB.Create(user)

	if result.Error != nil {
		return errors.New("Fail to create user")
	}

	return nil
}

func (usv *UserService) Read() (err error) {
	return nil
}

func (usv *UserService) Update() (err error) {
	return nil
}

func (usv *UserService) Delete() (err error) {
	return nil
}
