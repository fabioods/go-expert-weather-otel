package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/fabioods/go-expert-call-weather/configs"
	"github.com/fabioods/go-expert-call-weather/internal/domain"
	"github.com/fabioods/go-expert-call-weather/pkg/errorformated"
	"github.com/fabioods/go-expert-call-weather/pkg/trace"
)

type WeatherClient struct {
	WeatherApiURL     string
	WeatherApiTimeout int
}

type WeatherResponse struct {
	Cep                   string  `json:"cep"`
	CelsiusTemperature    float64 `json:"celsius_temperature"`
	FahrenheitTemperature float64 `json:"fahrenheit_temperature"`
	KelvinTemperature     float64 `json:"kelvin_temperature"`
}

func (w *WeatherResponse) ToCep() *domain.Cep {
	return domain.NewCep(w.Cep, w.CelsiusTemperature, w.KelvinTemperature, w.FahrenheitTemperature)
}

func NewWeatherClient(config *configs.Config) *WeatherClient {
	return &WeatherClient{
		WeatherApiURL:     config.WeatherApiURL,
		WeatherApiTimeout: config.WeatherApiTimeout,
	}
}

func (w *WeatherClient) WeatherByCep(ctx context.Context, cep string) (domain.Cep, error) {
	contextWithTimeOut, cancel := context.WithTimeout(ctx, time.Duration(w.WeatherApiTimeout)*time.Millisecond)
	defer cancel()

	path := fmt.Sprintf("%s/%s", w.WeatherApiURL, url.QueryEscape(cep))
	req, err := http.NewRequestWithContext(contextWithTimeOut, http.MethodGet, path, nil)
	if err != nil {
		return domain.Cep{}, errorformated.UnexpectedError(trace.GetTrace(), "error_creating_request", "error creating request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return domain.Cep{}, errorformated.UnexpectedError(trace.GetTrace(), "error_requesting_address", "error requesting address: %v", err)
	}
	defer res.Body.Close()

	bytesResponse, err := io.ReadAll(res.Body)
	if err != nil {
		return domain.Cep{}, errorformated.UnexpectedError(trace.GetTrace(), "error_reading_response", "error reading response: %v", err)
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(bytesResponse, &weatherResponse)
	if err != nil {
		return domain.Cep{}, errorformated.UnexpectedError(trace.GetTrace(), "error_unmarshalling_response_weather", "error unmarshalling response weather api: %v", err)
	}

	return *weatherResponse.ToCep(), nil
}
