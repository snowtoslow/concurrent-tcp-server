package models

// struct to get home response

type HomeResponse struct {
	Message string `json:"msg,omitempty"`
	Link    `json:"link,omitempty"`
}
