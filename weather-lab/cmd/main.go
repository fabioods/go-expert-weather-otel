package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/fabioods/go-expert-wheater-lab/configs"
	"github.com/fabioods/go-expert-wheater-lab/internal/handler"
	"github.com/fabioods/go-expert-wheater-lab/internal/infra/client/viacepclient"
	"github.com/fabioods/go-expert-wheater-lab/internal/infra/client/weatherclient"
	"github.com/fabioods/go-expert-wheater-lab/internal/infra/webserver"
	"github.com/fabioods/go-expert-wheater-lab/internal/usecase"
	otelPkg "github.com/fabioods/go-expert-wheater-lab/pkg/otel"
	"go.opentelemetry.io/otel"
)

func main() {
	c := configs.LoadConfig()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	shutdown, err := otelPkg.InitProvider(os.Getenv("OTEL_SERVICE_NAME"), os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider: %w", err)
		}
	}()

	tracer := otel.Tracer("weather-lab")

	ws := webserver.NewWebServer(c.Port, tracer)
	cepClient := viacepclient.NewViaCepClient(c)
	weatherClient := weatherclient.NewWeatherClient(c)
	useCase := usecase.NewWeatherByCepUseCase(weatherClient, cepClient)
	weatherHandler := handler.NewWeatherByCepHandler(useCase)
	ws.AddHandler("/weather/cep/{cep}", weatherHandler.Handle)
	ws.Start()
}
