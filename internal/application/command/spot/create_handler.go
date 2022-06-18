package spot_command

import (
	"time"

	"github.com/nooderg/pipiSpot/internal/config"
	"github.com/nooderg/pipiSpot/internal/core/domain"
	"github.com/nooderg/pipiSpot/internal/core/ports"
	"github.com/nooderg/pipiSpot/internal/infrastructure/persistence"
)

type CreateCommandHandler struct {
	userRepository ports.SpotRepository
}

func (c CreateCommandHandler) New() *CreateCommandHandler {
	return &CreateCommandHandler{
		userRepository: persistence.NewSpotRepository(config.GetDBClient()),
	}
}

func (c *CreateCommandHandler) Handle(command UpdateCommand) (*domain.Spot, error) {
	spot := domain.Spot{
		Title:     command.Title,
		Content:   command.Content,
		Long:      command.Long,
		Lat:       command.Lat,
		UpdatedAt: time.Now(),
	}
	return &spot, c.userRepository.CreateSpot(&spot)
}
