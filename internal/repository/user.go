package repository

import (
	"github.com/nooderg/pipiSpot/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (u UserRepository) CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(&user).Error
}

func (u UserRepository) GetUserByID(db *gorm.DB, id uint) (*models.User, error) {
	var user *models.User
	err := db.Model(&models.User{}).Where("id = ?", id).First(&user).Error
	return user, err
}

func (u UserRepository) UpdateUser(db *gorm.DB, id uint, user *models.User) error {
	return db.Model(&models.User{}).Where("id = ?", user.ID).Updates(&user).Error
}

func (u UserRepository) DeleteUser(db *gorm.DB, id uint) error {
	return db.Where("id = ?", id).Delete(&models.User{}, id).Error
}
