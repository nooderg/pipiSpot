package ports

import "github.com/nooderg/pipiSpot/internal/core/domain"


type RatingsRepository interface {
	CreateRatings(*domain.Ratings) error
	GetRatingsByID(uint) (*domain.Ratings, error)
	UpdateRatings(*domain.Ratings) error
	DeleteRatings(uint) error
}