package user_command

type UpdateCommand struct {
	ID uint `json:"id"`	
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (c *UpdateCommand) Validate() {
	// TODO
}