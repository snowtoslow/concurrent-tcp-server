package responses

import "concurrent-tcp-server/models"

// interface where is defined main methods to continue work
type ResponseRepository interface {
	GetTokenAndHomeLink(link string) (*models.RegisterResponse, error)
	GetAllRoutes(link string, token string) (*models.HomeResponse, error)
}
