package main

import (
	"github.com/dickykmrlh/profile/config"
	"github.com/dickykmrlh/profile/internal/repository"
)

func main() {
	config.NewConfig()
	repository.NewDBConn()
}
