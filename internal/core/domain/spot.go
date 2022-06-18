package domain

import "time"

type Spot struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Long      string    `json:"long"`
	Lat       string    `json:"lat"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Comments  []Comment `json:"comments" gorm:"-"`
}
