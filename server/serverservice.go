package server

import (
	"concurrent-tcp-server/models/constant"
	"concurrent-tcp-server/utils"
	"fmt"
	"strings"
)

//command interface
type Command interface {
	execute(input *string)
}

//concrete validate command
type validate struct {
	device device
}

func (v *validate) execute(input *string) {
	v.device.Validate(input)
}

//concrete command parse
type parse struct {
	device device
}

func (p *parse) execute(input *string) {
	p.device.PrintResponse(input)
}

//receiver interface

type device interface {
	Validate(*string)
	PrintResponse(*string)
}

func (server *Server) Validate(input *string) {
	if strings.Contains(*input, constant.ExpectedInput) && len(strings.Fields(*input)) == 2 {
		*input = utils.ToSnakeCase(strings.Fields(*input)[1])
	} else {
		*input = "null"
	}
}

func (server *Server) PrintResponse(input *string) {
	if *input == "null" {
		server.connection.Write([]byte(fmt.Sprintf("%v\n", NotValidInput.Error())))
	} else {
		server.getData(*input)
	}
}

func (server *Server) getData(input string) {
	foundFlag := false
	for i := 0; i < len(server.myMap); i++ {
		if v, found := server.myMap[i][input]; found {
			foundFlag = true
			server.connection.Write([]byte(fmt.Sprintf("%v\n", v)))
		} else {
			if foundFlag == false {
				server.connection.Write([]byte(fmt.Sprintf("%v\n", ErrFieldNotFound.Error())))
				break
			}
			continue
		}
	}
}
