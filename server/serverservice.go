package server

import (
	"concurrent-tcp-server/models/constant"
	"concurrent-tcp-server/utils"
	"fmt"
	"log"
	"strings"
)

type response struct {
	command []Command
}

func (r *response) handleData(input string) {
	for _, c := range r.command {
		c.execute(input)
	}
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
	p.device.PrintResponse(input)
}

//receiver interface

type device interface {
	Validate(string)
	/*getData(string)
	errorReturn(bool)*/
	PrintResponse(string)
}

func (server *Server) Validate(input string) {
	log.Println("Validation")
	if strings.Contains(input, constant.ExpectedInput) && len(strings.Fields(input)) == 2 {
		input = utils.ToSnakeCase(strings.Fields(input)[1])
	} else {
		input = "null"
	}
	log.Println("MY NEW INPUT:", input)
}

func (server *Server) PrintResponse(input string) {
	log.Println("RESPONSE:", input)
	if input == "null" {
		server.connection.Write([]byte(fmt.Sprintf("%v\n", NotValidInput.Error())))
	} else {
		server.getData(input)
	}
}

func (server *Server) getData(input string) {
	log.Println("GET DATA:", input)
	for i := 0; i < len(server.myMap); i++ {
		if v, found := server.myMap[i][input]; found {
			server.connection.Write([]byte(fmt.Sprintf("%v\n", v)))
		} else {
			server.errorReturn(found)
		}
	}
}

func (server *Server) errorReturn(found bool) {
	if !found {
		server.connection.Write([]byte(fmt.Sprintf("%v\n", ErrFieldNotFound.Error())))
		return
	}
}
