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
func (e *Server) GetWeatherByName(w http.ResponseWriter, r *http.Request) {
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
			weatherResponse, err := e.weatherClient.SearchByName(cityName)
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
			} else {
				errorData, err := openweather.NewErrorResponse(weatherResponse.Body())
				if err != nil {
					helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
					return
				}
				log.Errorf("OpenWeather API returned error HTTP code: %v. JSON body: %v", weatherResponse.StatusCode(), *errorData)
				openWeatherErrorMsg := fmt.Sprintf("OpenWeather API error: %v", errorData.Message)
				helpers.RespondWithError(w, http.StatusBadRequest, openWeatherErrorMsg)
				return
			}

		}
		helpers.RespondWithJSON(w, http.StatusOK, response)
		return
	default:
		helpers.RespondWithError(w, http.StatusNotFound, "not found")
	}
}
