package spot_command

type UpdateCommand struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Long    string `json:"long"`
	Lat     string `json:"lat"`
}

func (c *UpdateCommand) Validate() {
	// TODO
}
