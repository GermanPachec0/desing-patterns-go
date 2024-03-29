package openweaher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notifier/config"
)

type IOpenWeather interface {
	GetWeather() (*Weather, error)
}

type OpenWeather struct {
	config     config.OpenWeatherConfig
	httpClient *http.Client
}

func New(cfg config.OpenWeatherConfig) *OpenWeather {
	client := http.DefaultClient
	return &OpenWeather{
		config:     cfg,
		httpClient: client,
	}
}

func (o *OpenWeather) GetWeather() (*Weather, error) {
	path := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?appid=%s&units=%s&lang=%s&lat=%s&lon=%s",
		o.config.AppID, o.config.Units, o.config.Lang, o.config.Lat, o.config.Lon)
	weather := Weather{}
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	res, err := o.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&weather)
	if err != nil {
		return nil, err
	}
	return &weather, nil

}
