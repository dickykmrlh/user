package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var cfg Config

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	cfg = Config{}
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%+v\n", cfg)
}
