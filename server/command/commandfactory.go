package command

type Response struct {
	Command []Command
}

func (r *Response) HandleData(input *string) {
	for _, c := range r.Command {
		c.Execute(input)
	}
}
