package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/fabioods/go-expert-call-weather/internal/domain"
	"github.com/fabioods/go-expert-call-weather/internal/infra/client/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInputDTO_DefineCep_ValidCep(t *testing.T) {
	input := InputDTO{}
	err := input.DefineCep("12345678")
	assert.NoError(t, err)
	assert.Equal(t, "12345678", input.Cep)
}

func TestInputDTO_DefineCep_InvalidCep(t *testing.T) {
	input := InputDTO{}
	err := input.DefineCep("12345")
	assert.Error(t, err)
	assert.Equal(t, "cep is invalid", err.Error())
}

func TestInputDTO_DefineCep_EmptyCep(t *testing.T) {
	input := InputDTO{}
	err := input.DefineCep("")
	assert.Error(t, err)
	assert.Equal(t, "cep is required", err.Error())
}

func TestWeatherByCepUseCase_Execute_Success(t *testing.T) {
	mockWeatherClient := new(mocks.WeatherByCepClient)
	useCase := NewWeatherByCepUseCase(mockWeatherClient)

	input := InputDTO{Cep: "12345678"}
	expectedCep := domain.Cep{
		Cep:                   "12345678",
		CelsiusTemperature:    25.0,
		FahrenheitTemperature: 77.0,
		KelvinTemperature:     298.15,
	}

	// Configurando o mock para retornar o valor esperado
	mockWeatherClient.On("WeatherByCep", mock.Anything, "12345678").Return(expectedCep, nil)

	result, err := useCase.Execute(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expectedCep, result)

	mockWeatherClient.AssertExpectations(t)
}

func TestWeatherByCepUseCase_Execute_ClientError(t *testing.T) {
	mockWeatherClient := new(mocks.WeatherByCepClient)
	useCase := NewWeatherByCepUseCase(mockWeatherClient)

	input := InputDTO{Cep: "12345678"}
	expectedError := errors.New("client error")

	// Configurando o mock para retornar um erro
	mockWeatherClient.On("WeatherByCep", mock.Anything, "12345678").Return(domain.Cep{}, expectedError)

	result, err := useCase.Execute(context.Background(), input)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, domain.Cep{}, result)

	mockWeatherClient.AssertExpectations(t)
}
