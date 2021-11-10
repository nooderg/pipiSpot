package repository

import (
	"github.com/nooderg/pipiSpot/internal/models"
	"gorm.io/gorm"
)

type SpotRepository struct{}

func (s SpotRepository) CreateSpot(db *gorm.DB, spot *models.Spot) error {
	return db.Create(&spot).Error
}

func (s SpotRepository) GetSpotByID(db *gorm.DB, id uint) (*models.Spot, error) {
	var spot *models.Spot
	err := db.Model(&models.Spot{}).Where("id = ?", id).First(&spot).Error
	return spot, err
}

func (s SpotRepository) UpdateSpot(db *gorm.DB, id uint, spot *models.Spot) error {
	return db.Model(&models.Spot{}).Where("id = ?", spot.ID).Updates(&spot).Error
}

func (s SpotRepository) DeleteSpot(db *gorm.DB, id uint) error {
	return db.Delete(&models.Spot{}, id).Error
}
