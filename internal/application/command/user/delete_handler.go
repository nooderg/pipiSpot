package user_command

import (
	"github.com/nooderg/pipiSpot/internal/config"
	"github.com/nooderg/pipiSpot/internal/core/ports"
	"github.com/nooderg/pipiSpot/internal/infrastructure/persistence"
)

type DeleteCommandHandler struct {
	userRepository ports.UserRepository
}

func (c DeleteCommandHandler) New() *DeleteCommandHandler {
	return &DeleteCommandHandler{
		userRepository: persistence.NewUserRepository(config.GetDBClient()),
	}
}

func (c *DeleteCommandHandler) Handle(command DeleteCommand) error {
	return c.userRepository.DeleteUser(command.ID)
}
