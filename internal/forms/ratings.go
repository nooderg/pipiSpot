package forms

import "github.com/nooderg/pipiSpot/internal/models"

type Ratings struct {
	Position      string `json:"position"`
	Accessibility bool   `json:"accessibility"`
	Exposition    bool   `json:"exposition"`
	Cleanliness   bool   `json:"cleanliness"`
}

func (r *Ratings) GetRatings(userID uint, spotID uint) models.Ratings {
	return models.Ratings{
		SpotID: spotID,
		UserID: userID,
		Position: r.Position,
		Accessibility: r.Accessibility,
		Exposition: r.Exposition,
		Cleanliness: r.Cleanliness,
	}
}
