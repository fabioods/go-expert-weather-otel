package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabioods/go-expert-call-weather/configs"
	"github.com/fabioods/go-expert-call-weather/pkg/errorformated"
	"github.com/stretchr/testify/assert"
)

func TestWeatherByCep_Success(t *testing.T) {
	// Criação de um servidor de teste que retorna uma resposta simulada
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := WeatherResponse{
			Cep:                   "12345678",
			CelsiusTemperature:    25.0,
			FahrenheitTemperature: 77.0,
			KelvinTemperature:     298.15,
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Configuração do cliente com a URL do servidor de teste
	config := &configs.Config{
		WeatherApiURL:     server.URL,
		WeatherApiTimeout: 1000,
	}
	client := NewWeatherClient(config)

	cep, err := client.WeatherByCep(context.Background(), "12345678")
	assert.NoError(t, err)
	assert.Equal(t, "12345678", cep.Cep)
	assert.Equal(t, 25.0, cep.CelsiusTemperature)
	assert.Equal(t, 77.0, cep.FahrenheitTemperature)
	assert.Equal(t, 298.15, cep.KelvinTemperature)
}

func TestWeatherByCep_RequestError(t *testing.T) {
	// Configurando um cliente com uma URL inválida para simular erro de requisição
	config := &configs.Config{
		WeatherApiURL:     "http://invalid-url",
		WeatherApiTimeout: 1000,
	}
	client := NewWeatherClient(config)

	_, err := client.WeatherByCep(context.Background(), "12345678")
	var formattedErr *errorformated.ErrorFormated
	assert.ErrorAs(t, err, &formattedErr)
	assert.Equal(t, "error_requesting_address", err.(*errorformated.ErrorFormated).Code)
}

func TestWeatherByCep_ResponseReadError(t *testing.T) {
	// Criação de um servidor de teste que fecha a conexão imediatamente, simulando erro de leitura
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "1")
		// Não escreve nada no corpo, causando erro de leitura
	}))
	defer server.Close()

	config := &configs.Config{
		WeatherApiURL:     server.URL,
		WeatherApiTimeout: 1000,
	}
	client := NewWeatherClient(config)

	_, err := client.WeatherByCep(context.Background(), "12345678")
	var formattedErr *errorformated.ErrorFormated
	assert.ErrorAs(t, err, &formattedErr)
	assert.Equal(t, "error_reading_response", err.(*errorformated.ErrorFormated).Code)
}

func TestWeatherByCep_UnmarshalError(t *testing.T) {
	// Criação de um servidor de teste que retorna um JSON inválido
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("invalid json"))
	}))
	defer server.Close()

	config := &configs.Config{
		WeatherApiURL:     server.URL,
		WeatherApiTimeout: 1000,
	}
	client := NewWeatherClient(config)

	_, err := client.WeatherByCep(context.Background(), "12345678")
	var formattedErr *errorformated.ErrorFormated
	assert.ErrorAs(t, err, &formattedErr)
	assert.Equal(t, "error_unmarshalling_response_weather", err.(*errorformated.ErrorFormated).Code)
}
