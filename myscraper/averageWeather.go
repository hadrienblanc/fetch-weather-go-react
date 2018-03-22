package myscraper

import (
	"fmt"
)

// KelvinToCelsius allows to use the Conversion Formula : °Celsius = °Kelvin - 273.15
func KelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273
}

// ConvertWeatherToColor converts the MyWeatherInfo type, the information about
// the weather, into a html color code. example #ff5a42
func ConvertWeatherToColor(info MyWeatherInfo) string {
	blueValue := ((info.Humidity) / 100 * 255) + (125 * info.IsRaining)

	temperatureTmp := (KelvinToCelsius(info.Temperature) + 20)
	if temperatureTmp < 0 {
		temperatureTmp = 0
	}
	if temperatureTmp > 50 {
		temperatureTmp = 50
	}
	redValue := (temperatureTmp / 50) * 255

	greenValue := (info.Pollution / 0.00001) * 255
	if greenValue > 255 {
		greenValue = 255
	}

	colorHTML := fmt.Sprintf("#%02X%02X%02X", int(redValue), int(greenValue), int(blueValue))
	return colorHTML
}

// AverageWeather returns the average weather information between now,
// in 3 hours and in six hours.
// note that the pollution is a value fetched only from the current time.
func AverageWeather(city string) (MyWeatherInfo, error) {
	average := MyWeatherInfo{}

	infoNow, infoTreeHours, infoSixHours, err := Myscraper(city)

	if err != nil {
		return average, err
	}

	average.Temperature = (infoNow.Temperature + infoTreeHours.Temperature + infoSixHours.Temperature) / 3
	average.Humidity = (infoNow.Humidity + infoTreeHours.Humidity + infoSixHours.Humidity) / 3
	average.IsRaining = (infoNow.IsRaining + infoTreeHours.IsRaining + infoSixHours.IsRaining) / 3
	average.Pollution = infoNow.Pollution

	return average, nil
}
