package usecase

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"fmt"
	"html/template"

	dt "x/domain"
)


func TemplateHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to load template")
	}

	info := updateWeather()

	t.Execute(w, info)
}

func updateWeather() dt.WeatherDetail {
	water := rand.Intn(99) + 1
	wind := rand.Intn(99) + 1

	weather := dt.Weather{}

	weather.Status.Wind = wind
	weather.Status.Water = water

	updateFile(weather, "file.json")

	return getStatusInfo(water, wind)
}

func getStatusInfo(water int, wind int) dt.WeatherDetail {
	weatherDetail := dt.WeatherDetail{
		Water:       water,
		WaterStatus: getWaterStatus(water),
		Wind:        wind,
		WindStatus:  getWindStatus(wind),
	}

	return weatherDetail
}

func getWaterStatus(water int) string {
	if water <= 5 {
		return "Aman"
	} else if water >= 6 && water <= 8 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}

func getWindStatus(wind int) string {
	if wind <= 6 {
		return "Aman"
	} else if wind >= 7 && wind <= 15 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}


func updateFile(data dt.Weather, filename string) error {
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, value, 0644)
}
