package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var cfg *Config

func NewConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(fmt.Errorf("fail to load env variables, %s", err))
	}

	cfg = &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Fatal(fmt.Errorf("fail to process config, %s", err))
	}
}

func GetConfig() *Config {
	if cfg == nil {
		log.Fatal("config is not initialize")
	}

	return cfg
}

func (c *Config) GetDBConfig() *DB {
	return c.DB
}

type Config struct {
	DB *DB
}

type DB struct {
	RunMigration    bool   `envconfig:"DB_RUN_MIGRATION"`
	Host            string `envconfig:"DB_HOST"`
	Port            int    `envconfig:"DB_PORT"`
	DBName          string `envconfig:"DB_NAME"`
	Username        string `envconfig:"DB_USERNAME"`
	Password        string `envconfig:"DB_PASSWORD"`
	SSLMode         string `envconfig:"DB_SSL_MODE"`
	PingTimeout     int    `envconfig:"DB_TIMEOUT" default:"3"`
	MaxIdleConns    int    `envconfig:"DB_MAX_IDLE_CONNS" default:"5"`
	MaxOpenCons     int    `envconfig:"DB_MAX_OPEN_CON" default:"5"`
	ConnMaxLifetime int    `envconfig:"DB_CONN_MAX_LIFE_TIME" default:"10"`
	ConnMaxIdleTime int    `envconfig:"DB_CONN_MAX_IDLE_TIME" default:"10"`
}
