package responses

import "concurrent-tcp-server/models"

// interface where is defined main methods to continue work
type ResponseRepository interface {
	GetTokenAndHomeLink(link string, registerResponse *models.RegisterResponse)(error, *models.RegisterResponse)
	GetAllRoutes(registerResponse *models.RegisterResponse)(error, *models.HomeResponse)
}
