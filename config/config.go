package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host   string
	DSN    string
	Secret string
}

func SetupConfig() (Config, error) {
	godotenv.Load()

	host := os.Getenv("HOST")
	if len(host) < 1 {
		return Config{}, errors.New("couldnt get host from env")
	}

	dsn := os.Getenv("DSN")
	if len(dsn) < 1 {
		return Config{}, errors.New("dsn error")
	}

	secret := os.Getenv("SECRET")
	if len(secret) < 1 {
		return Config{}, errors.New("secret error")
	}
	return Config{Host: host, DSN: dsn, Secret: secret}, nil
}
