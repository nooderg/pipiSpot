package command

import (
	"time"

	"github.com/nooderg/pipiSpot/internal/config"
	"github.com/nooderg/pipiSpot/internal/core/domain"
	"github.com/nooderg/pipiSpot/internal/core/ports"
	"github.com/nooderg/pipiSpot/internal/infrastructure/persistence"
)

type RegisterCommandHandler struct {
	userRepository ports.UserRepository
}

func (r RegisterCommandHandler) New() *RegisterCommandHandler {
	return &RegisterCommandHandler{
		userRepository: *persistence.NewUserRepository(config.GetDBClient()),
	}
}

func (r *RegisterCommandHandler) Handle(command RegisterCommand) (domain.User, error) {
	user := domain.User{
		FirstName: command.FirstName,
		LastName:  command.LastName,
		Email:     command.Email,
		Password:  command.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return user, r.userRepository.CreateUser(&user)
}
