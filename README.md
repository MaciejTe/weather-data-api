# Weather data REST API
[![Build Status](https://travis-ci.com/MaciejTe/TWFjaWVqIFRvbWN6dWsgcmVjcnVpdG1lbnQgdGFzaw-.svg?branch=master)](https://travis-ci.com/MaciejTe/TWFjaWVqIFRvbWN6dWsgcmVjcnVpdG1lbnQgdGFzaw-) 
[![Coverage Status](https://codecov.io/gh/MaciejTe/TWFjaWVqIFRvbWN6dWsgcmVjcnVpdG1lbnQgdGFzaw-/branch/master/graph/badge.svg)](https://codecov.io/gh/MaciejTe/TWFjaWVqIFRvbWN6dWsgcmVjcnVpdG1lbnQgdGFzaw-?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/MaciejTe/TWFjaWVqIFRvbWN6dWsgcmVjcnVpdG1lbnQgdGFzaw-)](https://goreportcard.com/report/github.com/MaciejTe/TWFjaWVqIFRvbWN6dWsgcmVjcnVpdG1lbnQgdGFzaw-)
[![Documentation](https://godoc.org/github.com/MaciejTe/TWFjaWVqIFRvbWN6dWsgcmVjcnVpdG1lbnQgdGFzaw-?status.svg)](https://godoc.org/github.com/MaciejTe/TWFjaWVqIFRvbWN6dWsgcmVjcnVpdG1lbnQgdGFzaw-)

Small Go microservice enabling users to retrieve information about the weather in the places of their choosing.

# Endpoints

The only endpoint is **/api/v1/weather**, which takes list of city names and aggregates their weather information into JSON list.

## Example

Request: ```http://<IP_ADDRESS>:<API_PORT>/api/v1/weather?cities=Szczecin,Warszawa```

Response: 
```json
[
     {
         "coord": {
             "lon": 14.55,
             "lat": 53.43
         },
         "weather": [
             {
                 "id": 800,
                 "main": "Clear",
                 "description": "clear sky",
                 "icon": "01n"
             }
         ],
         "base": "stations",
         "main": {
             "temp": 289.95,
             "feels_like": 288.5,
             "temp_min": 289.26,
             "temp_max": 290.37,
             "pressure": 1016,
             "humidity": 63
         },
         "visibility": 0,
         "wind": {
             "speed": 2.02,
             "deg": 360
         },
         "clouds": {
             "all": 0
         },
         "dt": 1592252418,
         "sys": {
             "type": 3,
             "id": 19799,
             "message": 0,
             "country": "PL",
             "sunrise": 1592188361,
             "sunset": 1592249526
         },
         "timezone": 7200,
         "id": 3083829,
         "name": "Szczecin",
         "cod": 200
     },
     {
         "coord": {
             "lon": 21.01,
             "lat": 52.23
         },
         "weather": [
             {
                 "id": 801,
                 "main": "Clouds",
                 "description": "few clouds",
                 "icon": "02n"
             }
         ],
         "base": "stations",
         "main": {
             "temp": 292.66,
             "feels_like": 291.49,
             "temp_min": 292.04,
             "temp_max": 293.15,
             "pressure": 1014,
             "humidity": 52
         },
         "visibility": 10000,
         "wind": {
             "speed": 1.5,
             "deg": 10
         },
         "clouds": {
             "all": 13
         },
         "dt": 1592252418,
         "sys": {
             "type": 1,
             "id": 1713,
             "message": 0,
             "country": "PL",
             "sunrise": 1592187247,
             "sunset": 1592247540
         },
         "timezone": 7200,
         "id": 756135,
         "name": "Warsaw",
         "cod": 200
     }
 ]
 ```

## Configuration
Currently file with following environment variables need to be given as --env-file parameter during Docker container startup:
1. API_PORT (default 8080) - port on which REST API listens for connections
2. OPEN_WEATHER_API_KEY - application ID for OpenWeather REST API

## Docker
 There are two dockerfiles: 
 * **Dockerfile** - production, multi stage build image
 * **Dockerile.dev** - development Dockerfile, one stage build with all necessary tools
 
### How to launch development environment
1. Enter project directory
2. Issue ```make build_image_dev```
3. Issue ```make dev```

### How to launch production environment
1. Enter project directory
2. Issue ```make build_image```
3. Prepare config.env file with proper parameters
4. Issue ```make run```

# TODO
- [ ] Use goroutines instead of for loop for getting weather data
- [ ] Rate limiting
- [x] Separate business logic from cache implementation
- [ ] Separate business logic from weather provider implementation
- [x] More tests (cache test)
- [ ] Add API performance profiling metrics
- [ ] Add cache expiration parameters (CACHE_EXPIRATION_TIME, CACHE_PURGE_TIME)
