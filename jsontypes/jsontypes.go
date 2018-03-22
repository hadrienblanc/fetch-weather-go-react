package weathertypes

// Weather information format from the json openweather API

// WWeatherInfo is the main type from the API to get the day weather condition.
// API page example : http://samples.openweathermap.org/data/2.5/weather?q=London,uk&appid=b6907d289e10d714a6e88b30761fae22
type WWeatherInfo struct {
	Coord      Position   `json:"coord"`
	Weather    []FWeather `json:"weather"`
	Base       string     `json:"base"`
	Main       WMain      `json:"main"`
	Visibility int        `json:"visibility"`
	Wind       WWind      `json:"wind"`
	Clouds     WCloud     `json:"clouds"`
	Dt         int        `json:"dt"`
	Sys        WSys       `json:"sys"`
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Cod        int        `json:"cod"`
}

type WSys struct {
	Type    int
	ID      int
	Message float64
	Country string
	Sunrise int
	Sunset  int
}

type WMain struct {
	Temp     float64
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
	Pressure float64
	Humidity float64
}

type WWind struct {
	Speed float64
	deg   float64
}

type WCloud struct {
	All int
}

type WeatherInfo struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

// Position contains the latitude and the longitude of a geolocalisation
type Position struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// ForecastMain is the main type from the API to get the forecast on the future
// days
type ForecastMain struct {
	Cod     string  `json:"cod"`
	Message float64 `json:"message"`
	Cnt     int
	List    []ForecastList
	City    FCity
}

type FCity struct {
	Id         int
	Name       string
	Coords     Position
	Country    string
	Population float64
}

type ForecastList struct {
	Dt      int
	Main    FMain
	Weather []FWeather
	Clouds  FCloud
	Wind    FWind
	Rain    FRain
	Sys     FSys
	Dt_txt  string
}

type FMain struct {
	Temp       float64
	Temp_min   float64
	Temp_max   float64
	Pressure   float64
	Sea_level  float64
	Grnd_level float64
	Humidity   float64
	Temp_kf    float64
}

type FWeather struct {
	Id          int
	Main        string
	Description string
	Icon        string
}

type FCloud struct {
	All int
}

type FWind struct {
	Speed float64
	Deg   float64
}

type FRain struct {
	//  "3h"  float64
}

type FSys struct {
	Pod string
}

type PollutionMain struct {
	Time     string
	Location Position
	Data     []PollutionData
}

type PollutionData struct {
	Precision float64
	Pressure  float64
	Value     float64
}
