package server

import (
	"concurrent-tcp-server/utils"
	"fmt"
	"strings"
)

type response struct {
	command Command
}

func (r *response) handleData(input string) {
	r.command.execute(input)
}

//command interface
type Command interface {
	execute(input string)
}

//concrete validate command
type validate struct {
	device device
}

func (v *validate) execute(input string) {
	v.device.Validate(input)
}

//concrete command parse

type parse struct {
	device device
}

func (p *parse) execute(input string) {
	p.device.Parse(input)
}

//receiver interface

type device interface {
	Validate(input string)
	Parse(input string)
}

func (server *Server) Validate(input string) {
	if strings.Contains(input, server.validationStr) && len(strings.Fields(input)) == 2 {
		server.input = utils.ToSnakeCase(strings.Fields(input)[1])
	} else {
		server.connection.Write([]byte(NotValidInput.Error() + "\n"))
	}
}

func (server *Server) Parse(input string) {
	for i := 0; i < len(server.myMap); i++ {
		if v, found := server.myMap[i][input]; found {
			server.connection.Write([]byte(fmt.Sprintf("%v\n", v)))
		} else {
			server.connection.Write([]byte(ErrFieldNotFound.Error() + "\n"))
		}
	}
}
