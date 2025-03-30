package config

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

// Config structure to hold application settings
type Config struct {
	ApiUrl         string
	ApiPort        string
	AllowedOrigins string

	DBHost     string
	DBPort     uint16
	DBUser     string
	DBPassword string
	DBName     string
}

var instantiated *Config
var once sync.Once

func New() *Config {
	once.Do(func() {
		// set the .env file as the config file
		viper.SetConfigFile(".env")

		// set default values
		viper.SetDefault("API_PORT", "8080")

		if err := viper.ReadInConfig(); err != nil {
			//log.Fatal("Failed to read config file, ensure .env file exists in the root directory.")
			log.Fatal(err)
		}

		// load environment variables and configuration into the Config structure
		instantiated = &Config{
			ApiUrl:         viper.GetString("API_URL"),
			ApiPort:        viper.GetString("API_PORT"),
			AllowedOrigins: viper.GetString("ALLOWED_ORIGINS"),

			DBHost:     viper.GetString("DB_HOST"),
			DBPort:     uint16(viper.GetInt("DB_PORT")),
			DBUser:     viper.GetString("DB_USER"),
			DBPassword: viper.GetString("DB_PASSWORD"),
			DBName:     viper.GetString("DB_NAME"),
		}
	})

	return instantiated
}
