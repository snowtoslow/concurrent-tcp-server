package main

import (
	"concurrent-tcp-server/config"
	"concurrent-tcp-server/models/constant"
	"concurrent-tcp-server/responses/repository"
	"github.com/joho/godotenv"
	"log"
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

	responseRepository := repository.NewResponseRepository()
	homeAndToken, err := responseRepository.GetTokenAndHomeLink("http://" + initializedConfigs.Host + initializedConfigs.RemoteServerPort + constant.TokenUri)
	if err != nil {
		log.Println("home and token error:", err)
	}

	myRoutes, err := responseRepository.GetAllRoutes("http://"+initializedConfigs.Host+initializedConfigs.RemoteServerPort+homeAndToken.HomeLink, homeAndToken.AccessToken)
	if err != nil {
		log.Println(err)
	}

	log.Println(myRoutes)
}
