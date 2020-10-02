package httpresponses

type RouteResponse struct {
	Data     string            `json:"data,omitempty"`
	Link     map[string]string `json:"link,omitempty"`
	MimeType string            `json:"mime_type,omitempty"`
}
