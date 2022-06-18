package user_command

import (
	"time"

	"github.com/nooderg/pipiSpot/internal/config"
	"github.com/nooderg/pipiSpot/internal/core/domain"
	"github.com/nooderg/pipiSpot/internal/core/ports"
	"github.com/nooderg/pipiSpot/internal/infrastructure/persistence"
)

type UpdateCommandHandler struct {
	userRepository ports.UserRepository
}

func (c UpdateCommandHandler) New() *UpdateCommandHandler {
	return &UpdateCommandHandler{
		userRepository: persistence.NewUserRepository(config.GetDBClient()),
	}
}

func (c *UpdateCommandHandler) Handle(command UpdateCommand) (*domain.User, error) {
	user := domain.User{
		FirstName: command.FirstName,
		LastName:  command.LastName,
		Email:     command.Email,
		Password:  command.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &user, c.userRepository.UpdateUser(&user)
}
