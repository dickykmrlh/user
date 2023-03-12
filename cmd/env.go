package main

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
