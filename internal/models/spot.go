package models

type Spot struct {
	ID                string    `json:"id"`
	UserID            string    `json:"user_id"`
	User              User      `json:"user"`
	Title             string    `json:"title"`
	Content           string    `json:"content"`
	Long              int64     `json:"long"`
	Lat               int64     `json:"lat"`
	RatingsStandingID string    `json:"-" gorm:"ratings_standing_id"`
	RatingsSittingID  string    `json:"-" gorm:"ratings_sitting_id"`
	RatingsStanding   Ratings   `json:"ratings_stading"`
	RatingsSitting    Ratings   `json:"ratings_sitting"`
	Comments          []Comment `json:"comments" gorm:"-"`
}

type Comment struct {
	SpotID  string `json:"spot_id"`
	UserID  string `json:"-" gorm:"user_id"`
	User    User   `json:"user" gorm:"-"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
