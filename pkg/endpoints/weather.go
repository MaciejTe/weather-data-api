package endpoints

import (
	"fmt"
	"github.com/MaciejTe/weatherapp/pkg/helpers"
	"github.com/MaciejTe/weatherapp/pkg/openweather"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// GetWeatherByName collects weather information from OpenWeather API and aggregates them in JSON list.
// Available query parameters:
// 1. cities: list of city names, comma separated, i.e. New York,Warszawa,Berlin
func (s *Server) GetWeatherByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		queryParams := r.URL.Query()
		citiesStr := queryParams.Get("cities")
		if citiesStr == "" {
			helpers.RespondWithError(w, http.StatusBadRequest, "You need to provide at least one city")
			return
		}
		cities := strings.Split(citiesStr, ",")

		var response []openweather.WeatherData
		for _, cityName := range cities {
			if cityData, found := s.cacheClient.Get(cityName); found && r.Header.Get("cache-Control") != "no-cacheClient" {
				switch typeFound := cityData.(type) {
				case *openweather.WeatherData:
					log.Debugf("Using cacheClient on city %v", cityName)
					response = append(response, *cityData.(*openweather.WeatherData))
					continue
				default:
					log.Errorf("cache error. City data details: %v. Type found: %T", cityData, typeFound)
					continue
				}
			} else {
				weatherResponse, err := s.weatherClient.SearchByName(cityName)
				if err != nil {
					log.Error(err)
					helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
					return
				}
				if weatherResponse.StatusCode() == http.StatusOK {
					weatherData, err := openweather.NewWeatherData(weatherResponse.Body())
					if err != nil {
						log.Error(err)
						helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
						return
					}
					response = append(response, *weatherData)
					log.Debugf("Adding city %v to cacheClient...", cityName)
					s.cacheClient.Set(cityName, weatherData)
				} else {
					errorData, err := openweather.NewErrorResponse(weatherResponse.Body())
					if err != nil {
						helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
						return
					}
					log.Errorf("OpenWeather API returned error HTTP code: %v. JSON body: %v", weatherResponse.StatusCode(), *errorData)
					openWeatherErrorMsg := fmt.Sprintf("OpenWeather API error (city %v): %v", cityName, errorData.Message)
					helpers.RespondWithError(w, http.StatusBadRequest, openWeatherErrorMsg)
					return
				}
			}
		}

		helpers.RespondWithJSON(w, http.StatusOK, response)
		return
	default:
		helpers.RespondWithError(w, http.StatusNotFound, "not found")
	}
}
