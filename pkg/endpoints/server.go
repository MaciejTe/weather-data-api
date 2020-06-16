package endpoints

import (
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"weatherapp/pkg/config"
	"weatherapp/pkg/openweather"
)

// Server structure holds HTTP client and configuration for all API endpoints.
type Server struct {
	weatherClient *openweather.Client
	configuration *config.Config
}

// NewServer initializes Server structure and returns pointer to it.
func NewServer(httpClient resty.Client) *Server {

	configuration := config.Get()
	log.Infof("Loaded configuration parameters. API port: %v, OpenWeather API key: %v", configuration.APIPort, configuration.OpenWeatherAPIKey)

	return &Server{
		weatherClient: openweather.NewClient(configuration.OpenWeatherAPIKey, httpClient),
		configuration: configuration,
	}
}

// GetConfig returns pointer to configuration structure (config.Config).
func (s *Server) GetConfig() *config.Config {
	return s.configuration
}
