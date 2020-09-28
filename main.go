package main

import (
	"concurrent-tcp-server/config"
	"concurrent-tcp-server/models"
	"concurrent-tcp-server/models/constant"
	"concurrent-tcp-server/responses/repository"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
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
var mutex = &sync.Mutex{}
var test13 string
var responses []string

func main() {
	initializedConfigs := config.New()

	runtime.GOMAXPROCS(7)
	var client = new(http.Client)

	responseRepository := repository.NewResponseRepository(client)
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
		wg.Add(1)
		v := v
		go func() {
			defer wg.Done()
			myStr, err = test("http://"+initializedConfigs.Host+initializedConfigs.RemoteServerPort+v, homeAndToken.AccessToken)
			if err != nil {
				log.Println(err)
			}
			if myStr.Link != nil {
				for _, v := range myStr.Link {
					myData, err := test("http://"+initializedConfigs.Host+initializedConfigs.RemoteServerPort+v, homeAndToken.AccessToken)
					if err != nil {
						log.Println("ERROR:", err)
					}
					test13 = myData.Data
				}
			}
			responses = append(responses, test13)
			runtime.Gosched()
		}()

	}

	wg.Wait() // await completion of miser and spendthrift
	fmt.Println("Final balance: ", responses)
}

func test(link string, token string) (homeResponse *models.ResponseTest, err error) {
	mutex.Lock()

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Fatal(err)
	}
	mutex.Unlock()
	req.Header.Set(constant.HeaderAccessToken, token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &homeResponse)
	if err != nil {
		return nil, err
	}

	return
}
