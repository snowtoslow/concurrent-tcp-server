package main

import (
	"concurrent-tcp-server/config"
	"concurrent-tcp-server/models"
	"concurrent-tcp-server/models/constant"
	"encoding/json"
	"fmt"
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
	initializedConfigs:= config.New()
	log.Println("http://"+initializedConfigs.Host+initializedConfigs.RemoteServerPort+constant.TokenUri)
	test("http://"+initializedConfigs.Host+initializedConfigs.RemoteServerPort+constant.TokenUri)
}

func test(path string){
	url := path
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	var tokenResponse models.RegisterResponse

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		panic(err)
	}

	log.Println(tokenResponse)
}




