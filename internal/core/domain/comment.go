package domain

type Comment struct {
	ID      uint   `json:"id"`
	SpotID  uint   `json:"spot_id"`
	UserID  uint   `json:"-" gorm:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
