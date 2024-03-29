package main

import (
	"notifier/config"
	"notifier/observer_pattern"

	openweaher "notifier/open_weaher"
	"notifier/worker"
)

func main() {

	cfg := config.OpenWeatherConfig{
		AppID: "f369635965b00ad16ced5da4da4b9f3b",
		Units: "metrics",
		Lang:  "es",
		Lat:   "-32.9595",
		Lon:   "-60.661542",
	}
	owClient := openweaher.New(cfg)

	weatherNoti := observer_pattern.NewWeatherNotifier()
	member1 := &observer_pattern.Member{
		Id: "pepe@gmail.com",
	}
	weatherNoti.Register(member1)
	member2 := &observer_pattern.Member{
		Id: "german@gmail.com",
	}
	weatherNoti.Register(member2)

	worker := worker.NewWeatherService(owClient, weatherNoti)

	err := worker.Start()
	if err != nil {
		panic(err)
	}
	select {}

}
