package openweather

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

const openWeatherURL string = "http://api.openweathermap.org/data/2.5"

// Client structure allows to communicate with OpenWeather REST API.
type Client struct {
	apiKey string
	client *resty.Client
}

// NewClient creates new OpenWeather REST API Client object.
func NewClient(apiKey string, httpClient resty.Client) *Client {
	return &Client{apiKey: apiKey, client: &httpClient}
}

// SearchByName returns weather information for given city name.
func (u *Client) SearchByName(cityName string) (*resty.Response, error) {
	endpoint := fmt.Sprintf("%s/weather?q=%s&appid=%s", openWeatherURL, cityName, u.apiKey)
	return u.client.R().SetHeader("Content-Type", "application/json").Get(endpoint)
}
