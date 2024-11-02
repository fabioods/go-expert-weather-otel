package domain

type Cep struct {
	Cep                   string  `json:"cep"`
	CelsiusTemperature    float64 `json:"celsius_temperature"`
	KelvinTemperature     float64 `json:"kelvin_temperature"`
	FahrenheitTemperature float64 `json:"fahrenheit_temperature"`
}

func NewCep(cep string, celsiusTemperature float64, kelvinTemperature float64, fahrenheitTemperature float64) *Cep {
	return &Cep{
		Cep:                   cep,
		CelsiusTemperature:    celsiusTemperature,
		KelvinTemperature:     kelvinTemperature,
		FahrenheitTemperature: fahrenheitTemperature,
	}
}
