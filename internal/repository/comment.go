package repository

import (
	"github.com/nooderg/pipiSpot/internal/models"
	"gorm.io/gorm"
)

type CommentRepository struct{}

func (c CommentRepository) CreateComment(db *gorm.DB, comment *models.Comment) error {
	return db.Create(&comment).Error
}

func (c CommentRepository) GetCommentByID(db *gorm.DB, spotID, id uint) (*models.Comment, error) {
	var comment *models.Comment
	err := db.Model(&models.Comment{}).Where("id = ?", id).First(&comment).Error
	return comment, err
}

func (c CommentRepository) UpdateComment(db *gorm.DB, spotID, id uint, comment *models.Comment) error {
	return db.Model(&models.Comment{}).Where("id = ?", comment.ID).Updates(&comment).Error
}

func (c CommentRepository) DeleteComment(db *gorm.DB, spotID, id uint) error {
	return db.Delete(&models.Comment{}, id).Error
}
