package responses

import (
	"concurrent-tcp-server/models/httpresponses"
)

// interface where is defined main methods to continue work
type ResponseRepository interface {
	GetTokenAndHomeLink(link string) (*httpresponses.RegisterResponse, error)
	GetAllRoutes(link string, token string) (*httpresponses.HomeResponse, error)
}
