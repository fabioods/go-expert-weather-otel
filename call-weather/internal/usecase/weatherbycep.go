package usecase

import (
	"context"
	"errors"
	"regexp"

	"github.com/fabioods/go-expert-call-weather/internal/domain"
	"github.com/fabioods/go-expert-call-weather/internal/infra/client"
	"github.com/fabioods/go-expert-call-weather/pkg/otel"
)

type InputDTO struct {
	Cep string `json:"cep"`
}

func (w *InputDTO) DefineCep(cep string) error {
	err := w.ValidateCep(cep)
	if err != nil {
		return err
	}
	w.Cep = cep
	return nil
}

func (w *InputDTO) ValidateCep(cep string) error {
	if cep == "" {
		return errors.New("cep is required")
	}
	cepRegex := `^\d{8}$`
	matched, err := regexp.MatchString(cepRegex, cep)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("cep is invalid")
	}
	return nil
}

//go:generate mockery --all --case=underscore --disable-version-string
type WeatherByCepUseCase interface {
	Execute(context context.Context, input InputDTO) (domain.Cep, error)
}

type weatherByCepUseCase struct {
	weatherClient client.WeatherByCepClient
}

func NewWeatherByCepUseCase(weatherClient client.WeatherByCepClient) *weatherByCepUseCase {
	return &weatherByCepUseCase{
		weatherClient: weatherClient,
	}
}

func (w *weatherByCepUseCase) Execute(context context.Context, input InputDTO) (domain.Cep, error) {
	tracer := otel.TracerFromContext(context)
	ctx, span := tracer.Start(context, "usecase")
	defer span.End()
	city, err := w.weatherClient.WeatherByCep(ctx, input.Cep)
	if err != nil {
		return domain.Cep{}, err
	}

	return city, nil
}
