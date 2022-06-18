package user_command

type DeleteCommand struct {
	ID uint `json:"id"`
}

func (c *DeleteCommand) Validate() {
	// TODO
}
