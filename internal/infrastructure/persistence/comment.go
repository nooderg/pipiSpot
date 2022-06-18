package persistence

import (
	"github.com/nooderg/pipiSpot/internal/core/domain"
	"github.com/nooderg/pipiSpot/internal/core/ports"
	"gorm.io/gorm"
)

type CommentRepo struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepo {
	return &CommentRepo{db: db}
}

var _ ports.CommentRepository = &CommentRepo{}

func (cr CommentRepo) CreateComment(comment *domain.Comment) error {
	return cr.db.Create(&comment).Error
}

func (cr CommentRepo) GetCommentByID(id uint) (*domain.Comment, error) {
	var comment *domain.Comment
	err := cr.db.Model(&domain.Comment{}).Where("id = ?", id).First(&comment).Error
	return comment, err
}

func (cr CommentRepo) UpdateComment(comment *domain.Comment) error {
	return cr.db.Model(&domain.Comment{}).Where("id = ?", comment.ID).Updates(&comment).Error
}

func (cr CommentRepo) DeleteComment(id uint) error {
	return cr.db.Delete(&domain.Comment{}, id).Error
}
