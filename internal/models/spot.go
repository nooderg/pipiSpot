package models

type Spot struct {
	ID              string    `json:"id"`
	UserID          int       `json:"user_id"`
	Author          User      `json:"author"`
	Title           string    `json:"title"`
	Content         string    `json:"content"`
	Long            int       `json:"long"`
	Lat             int       `json:"lat"`
	RatingsStanding Ratings   `json:"ratingsStading"`
	RatingsSitting  Ratings   `json:"ratingsSitting"`
	Comments        []Comment `json:"comments"`
}

type Comment struct {
	UserID  int    `json:"user_id"`
	Author  User   `json:"author"`
	Content string `json:"content"`
	Origin  int    `json:"origin"`
}
