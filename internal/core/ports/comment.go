package ports

import "github.com/nooderg/pipiSpot/internal/core/domain"


type CommentRepository interface {
	CreateComment(*domain.Comment) error
	GetCommentByID(uint) (*domain.Comment, error)
	UpdateComment(*domain.Comment) error
	DeleteComment(uint) error
}