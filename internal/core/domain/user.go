package domain

import (
	"crypto"
	"html"
	"strings"
	"time"
)

type User struct {
	ID        uint       `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type PublicUser struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (u *User) BeforeSave() error {
	h := crypto.SHA1.New()
	_, err := h.Write([]byte(u.Password))
	if err != nil {
		return err
	}

	u.Password = string(h.Sum(nil))
	return nil
}


func (u *User) Prepare() {
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

type Users []User

func (us Users) PublicUsers() []interface{} {
	res := make([]interface{}, len(us))
	for i, u := range us {
		res[i] = u.PublicUser()
	}
	return res
}

func (u *User) PublicUser() interface{} {
	return &PublicUser{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}
