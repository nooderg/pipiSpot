package ports

import "github.com/nooderg/pipiSpot/internal/core/domain"


type SpotRepository interface {
	CreateSpot(*domain.Spot) error
	GetSpotByID(uint) (*domain.Spot, error)
	UpdateSpot(*domain.Spot) error
	DeleteSpot(uint) error
}