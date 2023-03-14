package main

import (
	"log"

	"github.com/dickykmrlh/user/config"
	"github.com/dickykmrlh/user/internal/repository"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("fail to load env variables, ", err)
	}

	config.InitConfig()
	repository.InitDBConn()
}
