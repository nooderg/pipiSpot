package persistence

import (
	"github.com/nooderg/pipiSpot/internal/core/domain"
	"github.com/nooderg/pipiSpot/internal/core/ports"
	"gorm.io/gorm"
)

type RatingsRepo struct {
	db *gorm.DB
}

func NewRatingsRepository(db *gorm.DB) *RatingsRepo {
	return &RatingsRepo{db: db}
}

var _ ports.RatingsRepository = &RatingsRepo{}

func (rr RatingsRepo) CreateRatings(ratings *domain.Ratings) error {
	return rr.db.Create(&ratings).Error
}

func (rr RatingsRepo) GetRatingsByID(id uint) (*domain.Ratings, error) {
	var ratings *domain.Ratings
	err := rr.db.Model(&domain.Ratings{}).Where("id = ?", id).First(&ratings).Error
	return ratings, err
}

func (rr RatingsRepo) UpdateRatings(ratings *domain.Ratings) error {
	return rr.db.Model(&domain.Ratings{}).Where("id = ?", ratings.ID).Updates(&ratings).Error
}

func (rr RatingsRepo) DeleteRatings(id uint) error {
	return rr.db.Delete(&domain.Ratings{}, id).Error
}
