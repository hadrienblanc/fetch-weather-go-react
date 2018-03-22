package main

import (
	"fmt"
	"github.com/hadrienblanc/challenge-go-weather/myscraper"
	"log"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func writeAnError(w http.ResponseWriter) {
	fmt.Fprintf(w, "{\"error\" : \"An error occured. Please try again.\"}\n")
}

func handler(w http.ResponseWriter, r *http.Request) {

	requestedCity := r.URL.Path[1:]

	if len(requestedCity) > 3 {
		enableCors(&w)

		info, err := myscraper.AverageWeather(requestedCity)

		if err != nil {
			writeAnError(w)
		} else {
			colorHTML := myscraper.ConvertWeatherToColor(info)
			fmt.Fprintf(w, "{\"color\" : \"%s\",\"city\" : \"%s\", \"temperature\" : %.2f}\n",
				colorHTML, requestedCity, myscraper.KelvinToCelsius(info.Temperature))
		}

	} else {
		writeAnError(w)
	}

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
