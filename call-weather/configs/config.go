package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	WeatherApiURL     string `mapstructure:"WEATHER_API_URL"`
	WeatherApiTimeout int    `mapstructure:"WEATHER_API_TIMEOUT"`
	Port              string `mapstructure:"PORT"`
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	weatherTimeout, err := strconv.Atoi(os.Getenv("WEATHER_API_TIMEOUT"))
	if err != nil {
		panic(err)
	}

	cfg := &Config{
		WeatherApiURL:     os.Getenv("WEATHER_API_URL"),
		WeatherApiTimeout: weatherTimeout,
		Port:              os.Getenv("PORT"),
	}
	if cfg.Port == "" {
		cfg.Port = "8081"
	}
	return cfg
}
