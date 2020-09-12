package repository

import (
	"concurrent-tcp-server/models"
	"concurrent-tcp-server/models/constant"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// this is something like a service where the logic of methods are implemented

type ResponseRepository struct {

}

func NewResponseRepository() *ResponseRepository {
	return &ResponseRepository{

	}
}

// method to get token and home link
func (responseRepository ResponseRepository) GetTokenAndHomeLink(link string) (*models.RegisterResponse, error) {
	request, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer request.Body.Close()

	responseBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	var registerResponse models.RegisterResponse

	err = json.Unmarshal(responseBody, &registerResponse)
	if err != nil {
		return nil, err
	}

	return &registerResponse, nil

}

// method to get token for home link and access all routes
func (responseRepository ResponseRepository) GetAllRoutes(link string, token string) (*models.HomeResponse, error) {
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
	var homeResponse models.HomeResponse

	err = json.Unmarshal(body, &homeResponse)
	if err != nil {
		return nil, err
	}

	return &homeResponse, nil
}
