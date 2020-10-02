package main

import (
	"concurrent-tcp-server/config"
	"concurrent-tcp-server/models/constant"
	"concurrent-tcp-server/parser"
	"concurrent-tcp-server/responses/repository"
	"github.com/joho/godotenv"
	"log"
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

var mutex = &sync.Mutex{}
var wg = &sync.WaitGroup{}

func main() {

	initializedConfigs := config.New()
	var myDataInArray = make(map[string]string)

	runtime.GOMAXPROCS(7)
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
		wg.Add(1)
		go func(value string) {
			defer wg.Done()
			responseRepository.GetLinkResponse("http://"+initializedConfigs.Host+initializedConfigs.RemoteServerPort+value, homeAndToken.AccessToken, myDataInArray)
			if err != nil {
				log.Println(err)
			}
			runtime.Gosched()
		}(v)

	}

	wg.Wait()

	for k, v := range myDataInArray {
		if dataToParse, exists := parser.ParsePlatforms[v]; exists {
			dataToParse.Parse(k)
		}
	}
}
