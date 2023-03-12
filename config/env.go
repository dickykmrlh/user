package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var cfg *Config

func NewConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	cfg = &Config{}
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetConfig() *Config {
	if cfg == nil {
		log.Fatal("config is not initialize")
	}

	return cfg
}

type Config struct {
	DB DB
}

type DB struct {
	RunMigration bool   `envconfig:"DB_RUN_MIGRATION"`
	Host         string `env:"DB_HOST" envDefault:"localhost"`
	Port         int    `env:"DB_PORT" envDefault:"5432"`
	DBName       string `env:"DB_NAME" envDefault:"gym"`
	Username     string `env:"DB_USERNAME"`
	Password     string `env:"DB_PASSWORD" secret:"GYM_DB_PASSWORD"`
}
