package main

import (
	"concurrent-tcp-server/config"
	"concurrent-tcp-server/models"
	"concurrent-tcp-server/models/constant"
	"concurrent-tcp-server/responses/repository"
	"encoding/json"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
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

	log.Println(test("http://"+initializedConfigs.Host+initializedConfigs.RemoteServerPort+myRoutes.Route1,homeAndToken.AccessToken))
}

func test(link string, token string) (*models.ResponseTest, error) {
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set(constant.HeaderAccessToken, token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var homeResponse models.ResponseTest

	err = json.Unmarshal(body, &homeResponse)
	if err != nil {
		return nil, err
	}

	return &homeResponse, nil
}
