package models

type Ratings struct {
	ID            uint   `json:"id"`
	SpotID        uint   `json:"spot_id"`
	UserID        uint   `json:"user_id"`
	Position      string `json:"position"`
	Accessibility bool   `json:"accessibility"`
	Exposition    bool   `json:"exposition"`
	Cleanliness   bool   `json:"cleanliness"`
}
