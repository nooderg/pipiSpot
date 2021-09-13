package models

type Ratings struct {
	SpotID                int    `json:"spot_id"`
	Position              string `json:"position"`
	AccessibilityLikes    int    `json:"accessibility_likes"`
	AccessibilityDislikes int    `json:"accessibility_dislikes"`
	ExpositionLikes       int    `json:"exposition_likes"`
	ExpositionDislikes    int    `json:"exposition_dislikes"`
	CleanlinessLikes      int    `json:"cleanliness_likes"`
	CleanlinessDislikes   int    `json:"cleanliness_dislikes"`
}
