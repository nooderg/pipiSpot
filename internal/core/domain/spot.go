package domain

type Spot struct {
	ID       uint      `json:"id"`
	UserID   uint      `json:"user_id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Long     string    `json:"long"`
	Lat      string    `json:"lat"`
	Comments []Comment `json:"comments" gorm:"-"`
}
