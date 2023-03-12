package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var cfg *Config

func NewConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("fail to load env variables, ", err)
	}

	cfg = &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Fatal("fail to process config, ", err)
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
	MigrationRun    bool   `envconfig:"DB_MIGRATION_RUN"`
	MigrationPath   string `envconfig:"DB_MIGRATION_PATH"`
}
