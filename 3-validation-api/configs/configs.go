package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Email    EmailConfig
	Password PasswordConfig
	Address  AddressConfig
}

type EmailConfig struct {
	Email string
}

type PasswordConfig struct {
	Password string
}

type AddressConfig struct {
	Address string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}

	return &Config{
		Email:    EmailConfig{Email: os.Getenv("EMAIL")},
		Password: PasswordConfig{Password: os.Getenv("PASSWORD")},
		Address:  AddressConfig{Address: os.Getenv("ADDRESS")},
	}
}
