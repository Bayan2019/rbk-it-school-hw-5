package model

import "time"

type OpenMeteoResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		Windspeed   float64 `json:"windspeed"`
		Weathercode int     `json:"weathercode"`
		Time        string  `json:"time"`
	} `json:"current_weather"`
}

type Weather struct {
	RequestedAt time.Time `db:"requested_at" json:"requested_at,omitempty"`
	Temperature float64   `db:"temperature" json:"temperature"`
	Description string    `db:"description" json:"description"`
}

type WeatherHistory struct {
	UserID      int64     `db:"user_id" json:"user_id"`
	City        string    `db:"city" json:"city,omitempty"`
	RequestedAt time.Time `db:"requested_at" json:"requested_at,omitempty"`
	Temperature float64   `db:"temperature" json:"temperature"`
	Description string    `db:"description" json:"description"`
}

// type WeatherResult struct {
// 	Latitude       float64 `json:"latitude"`
// 	Longitude      float64 `json:"longitude"`
// 	Temperature    float64 `json:"temperature"`
// 	WindSpeed      float64 `json:"wind_speed"`
// 	WeatherCode    int     `json:"weather_code"`
// 	Time           string  `json:"time"`
// 	Description    string  `json:"description"`
// 	Recommendation string  `json:"recommendation"`
// }
