package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabioods/go-expert-call-weather/internal/domain"
	"github.com/fabioods/go-expert-call-weather/internal/usecase"
	"github.com/fabioods/go-expert-call-weather/internal/usecase/mocks"
	"github.com/fabioods/go-expert-call-weather/pkg/errorformated"
	"github.com/fabioods/go-expert-call-weather/pkg/trace"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWeatherByCepHandler_Success(t *testing.T) {
	mockUseCase := new(mocks.WeatherByCepUseCase)
	handler := NewWeatherByCepHandler(mockUseCase)

	inputDTO := usecase.InputDTO{Cep: "12345678"}
	outputDTO := domain.Cep{Cep: "12345678", CelsiusTemperature: 25.0}

	// Configurando o comportamento esperado do mock
	mockUseCase.On("Execute", mock.Anything, inputDTO).Return(outputDTO, nil)

	// Criando uma requisição de teste
	r := httptest.NewRequest(http.MethodGet, "/weather/12345678", nil)
	w := httptest.NewRecorder()

	// Definindo o parâmetro da URL
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("cep", "12345678")
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	handler.Handle(w, r)

	// Verificando o resultado
	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var responseBody domain.Cep
	json.NewDecoder(resp.Body).Decode(&responseBody)
	assert.Equal(t, outputDTO, responseBody)

	mockUseCase.AssertExpectations(t)
}

func TestWeatherByCepHandler_InvalidCep(t *testing.T) {
	mockUseCase := new(mocks.WeatherByCepUseCase)
	handler := NewWeatherByCepHandler(mockUseCase)

	// Criando uma requisição com um CEP inválido
	r := httptest.NewRequest(http.MethodGet, "/weather/invalid_cep", nil)
	w := httptest.NewRecorder()

	// Definindo o parâmetro da URL
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("cep", "invalid_cep")
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	handler.Handle(w, r)

	// Verificando o resultado
	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
}

func TestWeatherByCepHandler_UseCaseError(t *testing.T) {
	mockUseCase := new(mocks.WeatherByCepUseCase)
	handler := NewWeatherByCepHandler(mockUseCase)

	inputDTO := usecase.InputDTO{Cep: "12345678"}
	errorMessage := errorformated.UnexpectedError(trace.GetTrace(), "error_creating_request", "error creating request: %v", nil)

	// Configurando o comportamento esperado do mock para retornar um erro
	mockUseCase.On("Execute", mock.Anything, inputDTO).Return(domain.Cep{}, errorMessage)

	// Criando uma requisição de teste
	r := httptest.NewRequest(http.MethodGet, "/weather/12345678", nil)
	w := httptest.NewRecorder()

	// Definindo o parâmetro da URL
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("cep", "12345678")
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	handler.Handle(w, r)

	// Verificando o resultado
	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	var responseBody map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&responseBody)
	assert.Equal(t, "error creating request: <nil>", responseBody["message"])

	mockUseCase.AssertExpectations(t)
}
