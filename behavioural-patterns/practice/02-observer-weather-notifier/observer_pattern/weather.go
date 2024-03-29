package observer_pattern

import "fmt"

type WeatherNotifier struct {
	observerList []Observer
	weather      string
}

func NewWeatherNotifier() *WeatherNotifier {
	return &WeatherNotifier{}
}

func (w *WeatherNotifier) Register(o Observer) {
	w.observerList = append(w.observerList, o)
}

func (w *WeatherNotifier) NotifyAll(weather string) {
	fmt.Printf("the actual weather was updated %s\n", w.weather)

	for _, m := range w.observerList {
		m.update(w.weather)
	}
}

func (w *WeatherNotifier) Deregister(o Observer) {
	w.observerList = removeFromSlice(w.observerList, o)
}

func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
