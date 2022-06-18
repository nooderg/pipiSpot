package persistence

import (
	"github.com/nooderg/pipiSpot/internal/core/ports"
	"github.com/nooderg/pipiSpot/internal/core/domain"
	"gorm.io/gorm"
)

type SpotRepo struct {
	db *gorm.DB
}

func NewSpotRepository(db *gorm.DB) *SpotRepo {
	return &SpotRepo{db: db}
}

var _ ports.SpotRepository = &SpotRepo{}

func (sr SpotRepo) CreateSpot(spot *domain.Spot) error {
	return sr.db.Create(&spot).Error
}

func (sr SpotRepo) GetSpotByID(id uint) (*domain.Spot, error) {
	var spot *domain.Spot
	err := sr.db.Model(&domain.Spot{}).Where("id = ?", id).First(&spot).Error
	return spot, err
}

func (sr SpotRepo) UpdateSpot(spot *domain.Spot) error {
	return sr.db.Model(&domain.Spot{}).Where("id = ?", spot.ID).Updates(&spot).Error
}

func (sr SpotRepo) DeleteSpot(id uint) error {
	return sr.db.Delete(&domain.Spot{}, id).Error
}
