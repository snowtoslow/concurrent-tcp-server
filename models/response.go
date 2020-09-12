package models

type Response map[string]interface{}

type ResponseTest struct {
	Message  string `json:"msg,omitempty"`
	Data     string `json:"data,omitempty"`
	Link     Link   `json:"link,omitempty"`
	MimeType string `json:"mime_type,omitempty"`
}
