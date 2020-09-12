package repository

import (
	"concurrent-tcp-server/models"
	"concurrent-tcp-server/models/constant"
	"encoding/json"
	"io/ioutil"
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
func (responseRepository ResponseRepository) GetTokenAndHomeLink(link string, registerResponse *models.RegisterResponse)(error,*models.RegisterResponse){
	request, err := http.Get(link)
	if err != nil {
		return err, nil
	}
	defer request.Body.Close()

	responseBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err, nil
	}
	err = json.Unmarshal(responseBody, &registerResponse)
	if err != nil {
		return err, nil
	}

	return nil,registerResponse

}

// method to get token for home link and access all routes
func (responseRepository ResponseRepository) GetAllRoutes(registerResponse *models.RegisterResponse)(error, *models.HomeResponse){
	request, err := http.Get(registerResponse.HomeLink)
	if err!=nil {
		return err,nil
	}

	request.Header.Set(constant.HeaderAccessToken,registerResponse.AccessToken)

	responseBody, err := ioutil.ReadAll(request.Body)
	if err!=nil {
		return err,nil
	}

	var homeResponse models.HomeResponse

	err = json.Unmarshal(responseBody,&homeResponse)
	if err!=nil {
		return err,nil
	}

	return nil, &homeResponse
}