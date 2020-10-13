package server

type response struct {
	command []Command
}

func (r *response) handleData(input *string) {
	for _, c := range r.command {
		c.execute(input)
	}
}
