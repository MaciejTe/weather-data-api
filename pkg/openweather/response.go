package openweather

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

// Coordinates structure holds information about given city coordinates.
type Coordinates struct {
	Longitude float32 `json:"lon"`
	Latitude float32 `json:"lat"`
}

// WeatherDescription structure holds information about given city weather description.
type WeatherDescription struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// MainWeatherData structure holds the most important information about given city.
type MainWeatherData struct {
	Temperature float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	TempMin float32 `json:"temp_min"`
	TempMax float32 `json:"temp_max"`
	Pressure int `json:"pressure"`
	Humidity int `json:"humidity"`
}

// WindData structure holds information about given city's wind conditions.
type WindData struct {
	Speed float32 `json:"speed"`
	Degrees int `json:"deg"`
}

// OtherData structure holds additional information about given city's weather conditions.
type OtherData struct {
	Type    int     `json:"type"`
	ID      int     `json:"id"`
	Message float32 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}

// WeatherData structure holds all, gathered information about given city's weather conditions.
type WeatherData struct {
	Coord      Coordinates          `json:"coord"`
	Weather    []WeatherDescription `json:"weather"`
	Base       string               `json:"base"`
	Main       MainWeatherData      `json:"main"`
	Visibility int                  `json:"visibility"`
	Wind       WindData             `json:"wind"`
	Cloudiness struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt       int       `json:"dt"`
	Sys      OtherData `json:"sys"`
	Timezone int       `json:"timezone"`
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Cod      int       `json:"cod"`
}

// ErrorData structure holds information about potential issue with OpenWeather API usage.
type ErrorData struct {
	Cod int `json:"cod,int,string"` // OpenWeather sometimes responds with int and sometimes with string...
	Message string `json:"message"`
}

// NewWeatherData initializes WeatherData structure from given response body.
func NewWeatherData(weatherResponseBody []byte) (weatherResult *WeatherData, err error) {
	weatherResult = &WeatherData{}
	err = json.Unmarshal(weatherResponseBody, &weatherResult)
	if err != nil {
		log.Errorf("Failed to unmarshal weather search response! Details: %v", err)
		return nil, err
	}
	return weatherResult, nil
}

// NewErrorResponse initializes ErrorData structure from given error response body.
func NewErrorResponse(errorResponseBody []byte) (errorResponse *ErrorData, err error) {
	errorResponse = &ErrorData{}
	err = json.Unmarshal(errorResponseBody, &errorResponse)
	if err != nil {
		log.Errorf("Failed to unmarshal error response! Details: %v", err)
		return nil, err
	}
	return errorResponse, nil
}
