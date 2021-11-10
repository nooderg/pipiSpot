package repository

import (
	"github.com/nooderg/pipiSpot/internal/models"
	"gorm.io/gorm"
)

type RatingsRepository struct{}

func (r RatingsRepository) CreateRatings(db *gorm.DB, ratings *models.Ratings) error {
	return db.Create(&ratings).Error
}

func (r RatingsRepository) GetRatingsByID(db *gorm.DB, spotID, id uint) (*models.Ratings, error) {
	var ratings *models.Ratings
	err := db.Model(&models.Ratings{}).Where("id = ?", id).First(&ratings).Error
	return ratings, err
}

func (r RatingsRepository) UpdateRatings(db *gorm.DB, spotID, id uint, ratings *models.Ratings) error {
	return db.Model(&models.Ratings{}).Where("id = ?", ratings.ID).Updates(&ratings).Error
}

func (r RatingsRepository) DeleteRatings(db *gorm.DB, spotID, id uint) error {
	return db.Delete(&models.Ratings{}, id).Error
}
