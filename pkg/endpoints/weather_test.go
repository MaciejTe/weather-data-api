package endpoints

import (
	"github.com/MaciejTe/weatherapp/pkg/cache"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func PkgDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func TestGetWeatherByName(t *testing.T) {
	inputURL := "api/v1/weather?cities=Szczecin"
	expectedPayload := `[{"coord":{"lon":14.55,"lat":53.43},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"base":"stations","main":{"temp":293.14,"feels_like":290.72,"temp_min":292.59,"temp_max":293.71,"pressure":1015,"humidity":49},"visibility":0,"wind":{"speed":3.13,"deg":40},"clouds":{"all":3},"dt":1592243597,"sys":{"type":3,"id":19799,"message":0,"country":"PL","sunrise":1592188361,"sunset":1592249526},"timezone":7200,"id":3083829,"name":"Szczecin","cod":200}]`
	expectedHTTPCode := 200

	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	cacheClient := cache.NewCache(10*time.Minute, 10*time.Minute)
	defer httpmock.DeactivateAndReset()

	inputOpenWeatherResponseBody, err := ioutil.ReadFile(filepath.Join(PkgDir(),"testdata", t.Name()+".golden"))
	if err != nil {
		t.Fatalf("failed reading .golden: %s", err)
	}
	httpmock.RegisterResponder("GET", `http://api.openweathermap.org/data/2.5/weather`,
		httpmock.NewStringResponder(200, string(inputOpenWeatherResponseBody)))

	req, err := http.NewRequest("GET", inputURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	server := NewServer(*client, cacheClient)
	handler := http.HandlerFunc(server.GetWeatherByName)

	handler.ServeHTTP(recorder, req)

	if recorder.Code != expectedHTTPCode {
		t.Errorf("handler returned wrong HTTP code: got %v want %v",
			recorder.Code, expectedHTTPCode)
	}

	if recorder.Body.String() != expectedPayload {
		t.Errorf("handler returned wrong JSON payload: got %v want %v",
			recorder.Body.String(), expectedPayload)
	}

	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()
	log.Infof("Call count info: %v", info)
}


func TestGetWeatherByNameNegative(t *testing.T) {
	testCaseTable := []struct{
		inputURL         string
		expectedPayload  string
		expectedHTTPCode int
	} {
		{
			inputURL:         "api/v1/weather",
			expectedPayload:  "{\"error\":\"You need to provide at least one city\"}",
			expectedHTTPCode: 400,
		},
		{
			inputURL:         "api/v1/weather?cities=NonexistingCity",
			expectedPayload:  `{"error":"OpenWeather API error (city NonexistingCity): city not found"}`,
			expectedHTTPCode: 400,
		},
	}
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	cacheClient := cache.NewCache(10*time.Minute, 10*time.Minute)
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", `http://api.openweathermap.org/data/2.5/weather`,
		httpmock.NewStringResponder(404, `{"cod":"404", "message": "city not found"}`))

	server := NewServer(*client, cacheClient)

	for _, testCase := range testCaseTable {
		req, err := http.NewRequest("GET", testCase.inputURL, nil)
		if err != nil {
			t.Fatal(err)
		}
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(server.GetWeatherByName)

		handler.ServeHTTP(recorder, req)

		if recorder.Code != testCase.expectedHTTPCode {
			t.Errorf("handler returned wrong HTTP code: got %v want %v",
				recorder.Code, testCase.expectedHTTPCode)
		}

		if recorder.Body.String() != testCase.expectedPayload {
			t.Errorf("handler returned wrong JSON payload: got %v want %v",
				recorder.Body.String(), testCase.expectedPayload)
		}
	}

	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()
	log.Infof("Call count info: %v", info)
}
