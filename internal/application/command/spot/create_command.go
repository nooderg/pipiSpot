package spot_command

type CreateCommand struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Long      string    `json:"long"`
	Lat       string    `json:"lat"`
}

func (c *CreateCommand) Validate() {
	// TODO: impl validate
}