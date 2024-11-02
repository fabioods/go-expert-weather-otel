package client

import (
	"context"

	"github.com/fabioods/go-expert-call-weather/internal/domain"
)

//go:generate mockery --all --case=underscore --disable-version-string
type WeatherByCepClient interface {
	WeatherByCep(ctx context.Context, cep string) (domain.Cep, error)
}
