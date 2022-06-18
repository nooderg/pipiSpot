package spot_command

import (
	"github.com/nooderg/pipiSpot/internal/config"
	"github.com/nooderg/pipiSpot/internal/core/ports"
	"github.com/nooderg/pipiSpot/internal/infrastructure/persistence"
)

type DeleteCommandHandler struct {
	spotRepository ports.SpotRepository
}

func (c DeleteCommandHandler) New() *DeleteCommandHandler {
	return &DeleteCommandHandler{
		spotRepository: persistence.NewSpotRepository(config.GetDBClient()),
	}
}

func (c *DeleteCommandHandler) Handle(command DeleteCommand) error {
	return c.spotRepository.DeleteSpot(command.ID)
}
