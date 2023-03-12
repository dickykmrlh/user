package main

import (
	"github.com/dickykmrlh/user/config"
	"github.com/dickykmrlh/user/internal/repository"
)

func main() {
	config.NewConfig()
	repository.NewDBConn()
	repository.RunningMigration()
}
