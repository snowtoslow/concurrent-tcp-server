package main

//add it in line 57 in case of some pizdetz runtime.Gosched()
import (
	"concurrent-tcp-server/config"
	"concurrent-tcp-server/models/constant"
	"concurrent-tcp-server/parser"
	"concurrent-tcp-server/responses/repository"
	"concurrent-tcp-server/server"
	"github.com/joho/godotenv"
	"log"
	"net"
	"net/http"
	"runtime"
	"sync"
)

// function to load env-project variables
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	initializedConfigs := config.New()
	var myDataInArray = make(map[string]string)
	filledGroupedData := parser.GroupedData{}
	mutex := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	connection := new(net.Conn)

	runtime.GOMAXPROCS(7) // allow the run-time support to utilize more than one OS thread(in this case 7)
	var client = new(http.Client)

	responseRepository := repository.NewResponseRepository(client, mutex, wg)
	homeAndToken, err := responseRepository.GetTokenAndHomeLink("http://" + initializedConfigs.Host + initializedConfigs.RemoteServerPort + constant.TokenUri)
	if err != nil {
		log.Println("home and token error:", err)
	}

	myRoutes, err := responseRepository.GetAllRoutes("http://"+initializedConfigs.Host+initializedConfigs.RemoteServerPort+homeAndToken.HomeLink, homeAndToken.AccessToken)
	if err != nil {
		log.Println(err)
	}

	var mainMap = myRoutes.Link

	for _, v := range mainMap {
		wg.Add(1) // spawn goroutines
		go func(value string) {
			defer wg.Done()
			responseRepository.GetLinkResponse("http://"+initializedConfigs.Host+initializedConfigs.RemoteServerPort+value, homeAndToken.AccessToken, myDataInArray)
			if err != nil {
				log.Println(err)
			}

		}(v)

	}

	wg.Wait()

	for k, v := range myDataInArray {
		if dataToParse, exists := parser.ParsePlatforms[v]; exists {
			if err := dataToParse.Parse(k, &filledGroupedData); err != nil {
				log.Println(err)
			}
		}
	}

	myServer := server.NewServer(*connection, filledGroupedData.FullMap)

	myServer.RunServer(initializedConfigs.TcpServerPort)

}
