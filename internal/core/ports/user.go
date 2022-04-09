package ports

import "github.com/nooderg/pipiSpot/internal/core/domain"

type userRepo struct {
	ur UserRepository
}

var _ UserRepository = &userRepo{}

type UserRepository interface {
	CreateUser(*domain.User) error
	GetUserByID(uint) (*domain.User, error)
	UpdateUser(*domain.User) error
	DeleteUser(uint) error
}