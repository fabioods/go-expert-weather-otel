package domain

import (
	"testing"
)

func TestNewCep(t *testing.T) {
	cep := "12345678"
	celsiusTemperature := 25.0
	kelvinTemperature := 298.15
	fahrenheitTemperature := 77.0

	result := NewCep(cep, celsiusTemperature, kelvinTemperature, fahrenheitTemperature)

	if result.Cep != cep {
		t.Errorf("Expected Cep to be '%s', but got '%s'", cep, result.Cep)
	}

	if result.CelsiusTemperature != celsiusTemperature {
		t.Errorf("Expected CelsiusTemperature to be %f, but got %f", celsiusTemperature, result.CelsiusTemperature)
	}

	if result.KelvinTemperature != kelvinTemperature {
		t.Errorf("Expected KelvinTemperature to be %f, but got %f", kelvinTemperature, result.KelvinTemperature)
	}

	if result.FahrenheitTemperature != fahrenheitTemperature {
		t.Errorf("Expected FahrenheitTemperature to be %f, but got %f", fahrenheitTemperature, result.FahrenheitTemperature)
	}
}
