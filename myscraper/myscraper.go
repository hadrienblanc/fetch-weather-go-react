package myscraper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hadrienblanc/challenge-go-weather/jsontypes"
)

// This function give the API key of a free account to access the API openweather
// a best practice is to put this API_KEY NOT in the code :)
func getAPIKey() string {
	return "6b15abf544dbe3102fe17afb98c4578e"
}

// getPolution will get the current pollution given the latitude and longitude
func getPollution(city string, lat float64, lon float64) (float64, error) {
	apiKey := getAPIKey()
	url := fmt.Sprintf("http://api.openweathermap.org/pollution/v1/co/%.0f,%.0f/current.json?appid=%s", lat, lon, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return 0, err
	}

	pollutionFetched := weathertypes.PollutionMain{}

	if err := json.Unmarshal(body, &pollutionFetched); err != nil {
		return 0, err
	}

	dataLength := len(pollutionFetched.Data)
	sum := 0.0

	for _, dataPoint := range pollutionFetched.Data {
		sum += dataPoint.Value
	}

	if dataLength == 0 {
		return 0, nil
	}

	return (sum / float64(dataLength)), nil
}

// ForecastToMyWeatherInfo convert the forecast information from the json
// to the MyWeatherInfo format used in our scraper.
func ForecastToMyWeatherInfo(forecastInfo weathertypes.ForecastMain) (MyWeatherInfo, MyWeatherInfo) {
	info := [2]MyWeatherInfo{}

	if len(forecastInfo.List) > 2 {
		for i := 0; i < 2; i++ {
			info[i].Temperature = forecastInfo.List[i+1].Main.Temp
			info[i].Humidity = forecastInfo.List[i+1].Main.Humidity
			if len(forecastInfo.List[i].Weather) > 1 {
				if forecastInfo.List[i].Weather[0].Main == "Rain" {
					info[i].IsRaining = 100.0
				} else {
					info[i].IsRaining = 0.0
				}
			}
		}
	}

	return info[0], info[1]
}

// getForecast will fetch the information of the weather in 3 hours and in
// 6 hours.
func getForecast(city string) (MyWeatherInfo, MyWeatherInfo, error) {
	apiKey := getAPIKey()
	inThreeHours := MyWeatherInfo{}
	inSixHours := MyWeatherInfo{}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return inThreeHours, inSixHours, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	forecastInfo := weathertypes.ForecastMain{}

	err2 := json.Unmarshal(body, &forecastInfo)

	if err2 != nil {
		return inThreeHours, inSixHours, err2
	}

	inThreeHours, inSixHours = ForecastToMyWeatherInfo(forecastInfo)

	return inThreeHours, inSixHours, nil
}

// WWeatherInfoToMyWeatherInfo takes the information of the json API struct format
// and return the MyWeatherInfo struct format
func WWeatherInfoToMyWeatherInfo(weather weathertypes.WWeatherInfo) MyWeatherInfo {
	info := MyWeatherInfo{}

	info.Temperature = weather.Main.Temp
	info.Humidity = weather.Main.Humidity
	info.IsRaining = 0.0
	info.Lat = weather.Coord.Lat
	info.Lon = weather.Coord.Lon

	if weather.Weather[0].Main == "Rain" {
		info.IsRaining = 100.0
	}

	return info
}

// weatherURL returns the url of the openweather API to have the information of
// one day
func weatherURL(city string) string {
	return fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s", city, getAPIKey())
}

// Myscraper takes the city name in arguments and returns the weather information
// of 3 moments : now, in 3 hours and in 6 hours.
func Myscraper(city string) (MyWeatherInfo, MyWeatherInfo, MyWeatherInfo, error) {
	infoNow := MyWeatherInfo{}
	infoTreeHours := MyWeatherInfo{}
	infoSixHours := MyWeatherInfo{}

	url := weatherURL(city)
	resp, err := http.Get(url)
	if err != nil {
		return infoNow, infoTreeHours, infoSixHours, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	weather := weathertypes.WWeatherInfo{}

	err = json.Unmarshal(body, &weather)

	if err != nil {
		return infoNow, infoTreeHours, infoSixHours, err
	}

	infoNow = WWeatherInfoToMyWeatherInfo(weather)

	infoTreeHours, infoSixHours, err = getForecast(city)

	infoNow.Pollution, err = getPollution(city, infoNow.Lat, infoNow.Lon)

	if err != nil {
		return infoNow, infoTreeHours, infoSixHours, err
	}

	return infoNow, infoTreeHours, infoSixHours, nil
}
