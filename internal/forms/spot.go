package forms

import "github.com/nooderg/pipiSpot/internal/models"

type Spot struct {
	Title             string    `json:"title"`
	Content           string    `json:"content"`
	Long              string     `json:"long"`
	Lat               string     `json:"lat"`
}

func (s Spot) GetSpot(userID uint) models.Spot {
	return models.Spot{
		UserID: userID,
		Title: s.Title,
		Content: s.Content,
		Long: s.Long,
		Lat: s.Lat,
	}
}