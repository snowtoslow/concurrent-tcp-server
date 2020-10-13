package server

//could be include in
//server.connection.Write([]byte(fmt.Sprintf("%10v", v)))
/*if i%10 == 0 {
	server.connection.Write([]byte("\n"))
}*/

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type Server struct {
	connection net.Conn
	myMap      []map[string]interface{}
}

func NewServer(connection net.Conn, myMap []map[string]interface{}) *Server {
	return &Server{
		connection: connection,
		myMap:      myMap,
	}
}

func (server *Server) RunServer(inputMap []map[string]interface{}, port string) {
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

func (server *Server) handleConnection(inputMap []map[string]interface{}) {
	log.Println("Handle connection is started!")

	for {
		netData, err := bufio.NewReader(server.connection).ReadString('\n')
		if err != nil {
			log.Fatal("Error reading: ", err)
		}

		myInput := strings.TrimSpace(netData)

		if myInput == "STOP" {
			break
		}

		onCommand := &validate{
			device: server,
		}

		ofCommand := &parse{
			device: server,
		}

		command := []Command{onCommand, ofCommand}

		response := &response{
			command: command,
		}

		response.handleData(myInput)

	}

	err := server.connection.Close()
	if err != nil {
		log.Println(err)
	}

}

/*func (server *Server) searchInParsedData(inputMap []map[string]interface{}, inputWord string) (err error) {
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

func (server *Server) validateServerInput(input string, validationString string) (error, string) {
	if strings.Contains(input, validationString) && len(strings.Fields(input)) == 2 {
		return nil, utils.ToSnakeCase(strings.Fields(input)[1])
	}
	return NotValidInput, input
}*/
