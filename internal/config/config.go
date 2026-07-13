package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	BotToken string `env:"BOT_TOKEN" env-required:"true"`
	Env      string `env:"env" env-required:"true"`
	DB       DB
}

type DB struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
}

func Load() (*Config, error) {
	var cfg Config
	if _, err := os.Stat(".env"); err != nil {
		log.Fatalf("project dont have a config file")
		return nil, err
	}
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		log.Fatalf("error reading config file")
		return nil, err
	}

	return &cfg, nil
}
