package persistence

import (
	"github.com/nooderg/pipiSpot/internal/core/domain"
	"github.com/nooderg/pipiSpot/internal/core/ports"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

var _ ports.UserRepository = &UserRepo{}

func (ur UserRepo) CreateUser(user *domain.User) error {
	return ur.db.Create(&user).Error
}

func (ur UserRepo) GetUserByID(id uint) (*domain.User, error) {
	var user *domain.User
	err := ur.db.Model(&domain.User{}).Where("id = ?", id).First(&user).Error
	return user, err
}

func (ur UserRepo) UpdateUser(user *domain.User) error {
	return ur.db.Model(&domain.User{}).Where("id = ?", user.ID).Updates(&user).Error
}

func (ur UserRepo) DeleteUser(id uint) error {
	return ur.db.Where("id = ?", id).Delete(&domain.User{}, id).Error
}
