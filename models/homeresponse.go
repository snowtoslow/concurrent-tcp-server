package models

// struct to get home response

type HomeResponse struct {
	Message string `json:"msg,omitempty"`
	MainLinks    `json:"link,omitempty"`
}
