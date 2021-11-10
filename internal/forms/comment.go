package forms

import "github.com/nooderg/pipiSpot/internal/models"

type Comment struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (c *Comment) GetComment(spotID uint) models.Comment {
	return models.Comment{
		SpotID: spotID,
		Title: c.Title,
		Content: c.Content,
	}
}
