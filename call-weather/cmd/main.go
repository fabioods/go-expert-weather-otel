package main

import (
	"github.com/fabioods/go-expert-call-weather/configs"
	"github.com/fabioods/go-expert-call-weather/internal/handler"
	"github.com/fabioods/go-expert-call-weather/internal/infra/client"
	"github.com/fabioods/go-expert-call-weather/internal/infra/webserver"
	"github.com/fabioods/go-expert-call-weather/internal/usecase"
)

func main() {
	c := configs.LoadConfig()
	ws := webserver.NewWebServer(c.Port)
	weatherClient := client.NewWeatherClient(c)
	useCase := usecase.NewWeatherByCepUseCase(weatherClient)
	weatherHandler := handler.NewWeatherByCepHandler(useCase)
	ws.AddHandler("/weather/cep/{cep}", weatherHandler.Handle)
	ws.Start()
}
