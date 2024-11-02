package configs

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Definindo variáveis de ambiente para o teste
	os.Setenv("WEATHER_API_URL", "https://api.weather.com")
	os.Setenv("WEATHER_API_TIMEOUT", "30")
	os.Setenv("PORT", "8080")
	defer os.Clearenv()

	config := LoadConfig()

	if config.WeatherApiURL != "https://api.weather.com" {
		t.Errorf("Expected WeatherApiURL to be 'https://api.weather.com', but got %s", config.WeatherApiURL)
	}

	if config.WeatherApiTimeout != 30 {
		t.Errorf("Expected WeatherApiTimeout to be 30, but got %d", config.WeatherApiTimeout)
	}

	if config.Port != "8080" {
		t.Errorf("Expected Port to be '8080', but got %s", config.Port)
	}
}

func TestLoadConfig_DefaultPort(t *testing.T) {
	// Configurando variáveis de ambiente sem a porta
	os.Setenv("WEATHER_API_URL", "https://api.weather.com")
	os.Setenv("WEATHER_API_TIMEOUT", "30")
	defer os.Clearenv()

	config := LoadConfig()

	// Verificando se o valor padrão da porta foi atribuído corretamente
	if config.Port != "8081" {
		t.Errorf("Expected default Port to be '8081', but got %s", config.Port)
	}
}

func TestLoadConfig_InvalidTimeout(t *testing.T) {
	// Configurando uma variável de ambiente com um timeout inválido
	os.Setenv("WEATHER_API_URL", "https://api.weather.com")
	os.Setenv("WEATHER_API_TIMEOUT", "invalid_timeout")
	defer os.Clearenv()

	// Verificando se ocorre um panic ao carregar um valor inválido
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic due to invalid WEATHER_API_TIMEOUT, but no panic occurred")
		}
	}()

	LoadConfig()
}
