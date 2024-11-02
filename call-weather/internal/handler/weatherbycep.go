package handler

import (
	"encoding/json"
	"net/http"

	"github.com/fabioods/go-expert-call-weather/internal/usecase"
	"github.com/fabioods/go-expert-call-weather/pkg/errorformated"
	"github.com/go-chi/chi/v5"
)

//go:generate mockery --all --case=underscore --disable-version-string
type WeatherByCepHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type weatherByCepHandler struct {
	useCase usecase.WeatherByCepUseCase
}

func NewWeatherByCepHandler(useCase usecase.WeatherByCepUseCase) *weatherByCepHandler {
	return &weatherByCepHandler{
		useCase: useCase,
	}
}

func (h *weatherByCepHandler) Handle(w http.ResponseWriter, r *http.Request) {
	input := usecase.InputDTO{}

	err := input.DefineCep(chi.URLParam(r, "cep"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		json.NewEncoder(w).Encode(err)
		return
	}

	output, err := h.useCase.Execute(r.Context(), input)
	if err != nil {
		statusCode := err.(*errorformated.ErrorFormated).StatusCode()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
