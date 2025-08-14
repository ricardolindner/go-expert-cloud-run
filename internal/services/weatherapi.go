package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var WeatherAPIClient = &http.Client{}
var WeatherAPIBaseURL = "http://api.weatherapi.com/v1"

type Weather struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func GetWeather(city string) (*Weather, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	encodedCity := url.QueryEscape(city)

	url := fmt.Sprintf("%s/current.json?key=%s&q=%s", WeatherAPIBaseURL, apiKey, encodedCity)

	resp, err := WeatherAPIClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not find weather for city: %s", city)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("WeatherAPI complete response: ", string(bodyBytes))

	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var weather Weather
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, err
	}

	return &weather, nil
}
