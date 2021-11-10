package forms

import "github.com/nooderg/pipiSpot/internal/models"

type User struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u User) GetUser() models.User {
	return models.User{
		Firstname: u.Firstname,
		Lastname: u.Lastname,
		Email: u.Email,
		Password: u.Password,
	}
}