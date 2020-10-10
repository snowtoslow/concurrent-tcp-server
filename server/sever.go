package server

//could be include in
//server.connection.Write([]byte(fmt.Sprintf("%10v", v)))
/*if i%10 == 0 {
	server.connection.Write([]byte("\n"))
}*/

import (
	"bufio"
	"concurrent-tcp-server/models/constant"
	"concurrent-tcp-server/utils"
	"fmt"
	"log"
	"net"
	"strings"
)

type Server struct {
	connection net.Conn
}

func NewServer(connection net.Conn) *Server {
	return &Server{
		connection: connection,
	}
}

func (server Server) RunServer(inputMap []map[string]interface{}, port string) {
	log.Println("Server start running on port: ", port)
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("Error while listening", err)
	}
	defer listen.Close()

	for {
		server.connection, err = listen.Accept()
		if err != nil {
			log.Println("Accept error: ", err)
			return
		}

		go server.handleConnection(inputMap)

	}
}

func (server Server) handleConnection(inputMap []map[string]interface{}) {
	log.Println("Handle connection is started!")

	for {
		netData, err := bufio.NewReader(server.connection).ReadString('\n')
		if err != nil {
			log.Fatal("Error reading: ", err)
		}

		if strings.TrimSpace(netData) == "STOP" {
			break
		}

		if err, myValidString := server.validateServerInput(strings.TrimSpace(netData), constant.ExpectedInput); err != nil {
			server.connection.Write([]byte(err.Error() + "\n"))
		} else {
			if err = server.searchInParsedData(inputMap, myValidString); err != nil {
				server.connection.Write([]byte(ErrFieldNotFound.Error() + "\n"))
			}
		}
	}

	err := server.connection.Close()
	if err != nil {
		log.Println(err)
	}

}

func (server Server) searchInParsedData(inputMap []map[string]interface{}, inputWord string) (err error) {
	foundFlag := false
	for i := 0; i < len(inputMap); i++ {
		if v, found := inputMap[i][inputWord]; found {
			foundFlag = true
			err = nil
			server.connection.Write([]byte(fmt.Sprintf("%v\n", v)))
		} else {
			if foundFlag == false {
				err = ErrFieldNotFound
			}
			continue
		}
	}

	return
}

func (server Server) validateServerInput(input string, validationString string) (error, string) {
	if strings.Contains(input, validationString) && len(strings.Fields(input)) == 2 {
		return nil, utils.ToSnakeCase(strings.Fields(input)[1])
	}
	return NotValidInput, input
}
