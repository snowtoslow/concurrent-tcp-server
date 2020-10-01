package repository

import (
	"concurrent-tcp-server/models"
	"concurrent-tcp-server/models/constant"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// this is something like a service where the logic of methods are implemented

type ResponseRepository struct {
	client *http.Client
	mutex  *sync.Mutex
	wg     *sync.WaitGroup
}

func NewResponseRepository(client *http.Client, mutex *sync.Mutex, wg *sync.WaitGroup) *ResponseRepository {
	return &ResponseRepository{
		client: client,
		mutex:  mutex,
		wg:     wg,
	}
}

// method to get token and home link
func (responseRepository ResponseRepository) GetTokenAndHomeLink(link string) (*models.RegisterResponse, error) {
	request, err := responseRepository.client.Get(link)
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
	resp, err := responseRepository.client.Do(req)
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

func (responseRepository ResponseRepository) GetLinkResponse(link string, token string, data *models.ResponseData) {
	var responseData *models.RouteResponse
	responseRepository.mutex.Lock()

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Fatal(err)
	}
	responseRepository.mutex.Unlock()
	req.Header.Set(constant.HeaderAccessToken, token)
	resp, err := responseRepository.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Println(err)
	}

	responseRepository.responseCallTest(responseData, token, data)
}

func (responseRepository ResponseRepository) responseCallTest(response *models.RouteResponse, token string, data *models.ResponseData) {
	if response.Link != nil {
		responseRepository.createArrayOfData(response.Data, data)
		for _, v := range response.Link {
			go responseRepository.GetLinkResponse("http://localhost:5000"+v, token, data)
		}
	} else {
		responseRepository.createArrayOfData(response.Data, data)
	}
}

func (responseRepository ResponseRepository) createArrayOfData(data string, responseData *models.ResponseData) {
	*responseData = append(*responseData, data)
}
