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

func (server *Server) RunServer(port string) {
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

		go server.handleConnection()

	}
}

func (server *Server) handleConnection() {
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

		validateCommand := &validate{
			device: server,
		}

		parseCommand := &parse{
			device: server,
		}

		command := []Command{validateCommand, parseCommand}

		response := &response{
			command: command,
		}

		response.handleData(&myInput)

	}

	err := server.connection.Close()
	if err != nil {
		log.Println(err)
	}
}
