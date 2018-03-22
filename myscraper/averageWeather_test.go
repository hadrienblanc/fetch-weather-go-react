package myscraper

import (
	"testing"
)

func TestKelvinToCelsius(t *testing.T) {

	kelvin := 290.0
	celsius := KelvinToCelsius(kelvin)

	celsiusExpected := (290.0 - 273.0)
	if celsius != celsius_expected {
		t.Errorf("Expected the celsius temperature '%f', have '%f'\n", celsiusExpected, celsius)
	}
}
