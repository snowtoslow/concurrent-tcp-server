package httpresponses

type RegisterResponse struct {
	AccessToken string `json:"access_token"`
	HomeLink    string `json:"link"`
}
