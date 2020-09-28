package models

// struct to get home response

type HomeResponse struct {
	Message string            `json:"msg,omitempty"`
	Link    map[string]string `json:"link,omitempty"`
}
