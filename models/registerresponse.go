package models

type TokenResponse struct{
	AccessToken string `json:"access_token"`
	HomeLink string `json:"link"`
}
