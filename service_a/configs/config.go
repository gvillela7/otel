package config

import (
	"errors"
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API          APIConfig
	WeatherAPI   WeatherAPI
	ZipkinConfig ZipkinConfig
}

type APIConfig struct {
	Port        string
	Environment string
}

type WeatherAPI struct {
	Key string
}
type ZipkinConfig struct {
	Endpoint string
}

func init() {
	//Service
	viper.SetDefault("api.port", "8000")
	viper.SetDefault("api.environment", "dev")
	viper.SetDefault("zipkin.endpoint", "")
	viper.SetDefault("weather.key", "")

}

func Load(viperPath ...string) error {
	if len(viperPath) > 0 {
		viper.AddConfigPath(viperPath[0])
	} else {
		viper.AddConfigPath(".")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port:        viper.GetString("api.port"),
		Environment: viper.GetString("api.environment"),
	}

	cfg.WeatherAPI = WeatherAPI{
		Key: viper.GetString("weather.key"),
	}

	cfg.ZipkinConfig = ZipkinConfig{
		Endpoint: viper.GetString("zipkin.endpoint"),
	}
	return nil
}

func GetAPIConfig() APIConfig {
	return cfg.API
}
func GetZipkinConfig() ZipkinConfig {
	return cfg.ZipkinConfig
}
func GetWeatherAPI() WeatherAPI {
	return cfg.WeatherAPI
}
