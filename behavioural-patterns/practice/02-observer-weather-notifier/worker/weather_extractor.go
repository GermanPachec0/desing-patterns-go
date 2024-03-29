package worker

import (
	"fmt"

	"notifier/observer_pattern"
	openweaher "notifier/open_weaher"
	"time"
)

type Worker interface {
	Start()
}

type WeatherService struct {
	owClient openweaher.IOpenWeather
	notifier observer_pattern.Subject
}

func NewWeatherService(ow openweaher.IOpenWeather, notifier observer_pattern.Subject) *WeatherService {
	return &WeatherService{
		owClient: ow,
		notifier: notifier,
	}
}

func (w WeatherService) Start() error {

	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)
	fmt.Println("Starting Ticker")
	go func() {
		for {
			defer ticker.Stop() // Ensure ticker is stopped when goroutine exits

			select {
			case <-done:
				return
			case <-ticker.C:
				wtd, err := w.owClient.GetWeather()
				if err != nil {
					panic(err.Error())
				}
				w.notifier.NotifyAll(wtd.Weather[0].Description)
			}

		}
	}()
	return nil
}
