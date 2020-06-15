package config

import (
	log "github.com/sirupsen/logrus"
	"os"
)

// Config structure holds configuration parameters taken from environment variables.
type Config struct {
	APIPort           string
	OpenWeatherAPIKey string
}

// Get retrieves all configuration parameters from environment variables. List of currently used environment variables:
// 1. API_PORT (default 8080) - port on which REST API listens for connections
// 2. OPEN_WEATHER_API_KEY - application ID for OpenWeather REST API
func Get() *Config {
	conf := &Config{
		APIPort: "8080",
	}

	conf.APIPort = os.Getenv("API_PORT")
	if conf.APIPort == "" {
		log.Warn("API port configuration environment variable (API_PORT) not set. Setting default value: 8080")
	}
	conf.OpenWeatherAPIKey = os.Getenv("OPEN_WEATHER_API_KEY")
	if conf.OpenWeatherAPIKey == "" {
		log.Fatal("OpenWeather API key not set. Please set OPEN_WEATHER_API_KEY environment variable.")
	}

	return conf
}
