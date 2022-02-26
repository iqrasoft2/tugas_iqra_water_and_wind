package domain

type Weather struct {
	Status Status `json:"status"`
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type WeatherDetail struct {
	Water       int    `json:"water"`
	WaterStatus string `json:"waterStatus"`
	Wind        int    `json:"wind"`
	WindStatus  string `json:"windStatus"`
}