package server

import (
	"concurrent-tcp-server/models/constant"
	"concurrent-tcp-server/utils"
	"fmt"
	"log"
	"strings"
)

func (server *Server) Validate(input *string) {
	log.Println("VALIDATE")
	if strings.Contains(*input, constant.ExpectedInput) && len(strings.Fields(*input)) == 2 {
		*input = utils.ToSnakeCase(strings.Fields(*input)[1])
	} else {
		*input = "null"
	}
}

func (server *Server) PrintResponse(input *string) {
	log.Println("Print response")
	if *input == "null" {
		server.connection.Write([]byte(fmt.Sprintf("%v\n", NotValidInput.Error())))
	} else {
		server.getData(*input)
	}
}

func (server *Server) getData(input string) {
	log.Println("GET DATA!")
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
