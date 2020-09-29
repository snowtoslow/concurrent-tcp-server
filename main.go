package main

import (
	"concurrent-tcp-server/config"
	"concurrent-tcp-server/models"
	"concurrent-tcp-server/models/constant"
	"concurrent-tcp-server/responses/repository"
	"fmt"
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

// balance for shared bank account

var test13 string
var responses []string
var mutex = &sync.Mutex{}
var myStr *models.ResponseTest

func main() {
	initializedConfigs := config.New()

	runtime.GOMAXPROCS(7)
	var client = new(http.Client)

	responseRepository := repository.NewResponseRepository(client, mutex)
	homeAndToken, err := responseRepository.GetTokenAndHomeLink("http://" + initializedConfigs.Host + initializedConfigs.RemoteServerPort + constant.TokenUri)
	if err != nil {
		log.Println("home and token error:", err)
	}

	myRoutes, err := responseRepository.GetAllRoutes("http://"+initializedConfigs.Host+initializedConfigs.RemoteServerPort+homeAndToken.HomeLink, homeAndToken.AccessToken)
	if err != nil {
		log.Println(err)
	}

	var wg sync.WaitGroup
	var mainMap = myRoutes.Link
	var myStr *models.ResponseTest

	for _, v := range mainMap {
		wg.Add(7)
		v := v
		go func() {
			defer wg.Done()
			myStr, err = responseRepository.GetLinkResponse("http://"+initializedConfigs.Host+initializedConfigs.RemoteServerPort+v, homeAndToken.AccessToken)
			if err != nil {
				log.Println(err)
			}
			if myStr.Link != nil {
				for _, v := range myStr.Link {
					myData, err := responseRepository.GetLinkResponse("http://"+initializedConfigs.Host+initializedConfigs.RemoteServerPort+v, homeAndToken.AccessToken)
					if err != nil {
						log.Println("ERROR:", err)
					}
					log.Println("DATA:", myData)
				}
			}
			//responses = append(responses, test13)
			runtime.Gosched() // check this shit
		}()

	}

	wg.Wait()
	fmt.Println("Final balance: ", responses)
}

/*PUT IT IN MAIN*/
/*
initializedConfigs := config.New()

runtime.GOMAXPROCS(7)
var client = new(http.Client)

responseRepository := repository.NewResponseRepository(client, mutex)
homeAndToken, err := responseRepository.GetTokenAndHomeLink("http://" + initializedConfigs.Host + initializedConfigs.RemoteServerPort + constant.TokenUri)
if err != nil {
log.Println("home and token error:", err)
}

myRoutes, err := responseRepository.GetAllRoutes("http://"+initializedConfigs.Host+initializedConfigs.RemoteServerPort+homeAndToken.HomeLink, homeAndToken.AccessToken)
if err != nil {
log.Println(err)
}


var wg sync.WaitGroup
var mainMap = myRoutes.Link
//var myStr *models.ResponseTest


for _, v := range mainMap {
wg.Add(7)

go responseRepository.GetLinkResponse("http://"+initializedConfigs.Host+initializedConfigs.RemoteServerPort+v, homeAndToken.AccessToken)

runtime.Gosched() // check this shit
}
wg.Done()
wg.Wait()
fmt.Println("Final balance: ", responses)*/
