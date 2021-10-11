package models

type Ratings struct {
	SpotID                string `json:"spot_id"`
	AccessibilityLikes    uint   `json:"accessibility_likes"`
	AccessibilityDislikes uint   `json:"accessibility_dislikes"`
	ExpositionLikes       uint   `json:"exposition_likes"`
	ExpositionDislikes    uint   `json:"exposition_dislikes"`
	CleanlinessLikes      uint   `json:"cleanliness_likes"`
	CleanlinessDislikes   uint   `json:"cleanliness_dislikes"`
}
