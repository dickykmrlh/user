package main

import (
	"log"

	"github.com/dickykmrlh/user/config"
	service "github.com/dickykmrlh/user/internal/core/services"
	"github.com/dickykmrlh/user/internal/handler"
	"github.com/dickykmrlh/user/internal/repository"
	"github.com/joho/godotenv"
)

func main() {
	//config
	if err := godotenv.Load(); err != nil {
		log.Fatal("fail to load env variables, ", err)
	}
	config.InitConfig()

	//repository
	repository.InitDBConn(config.GetConfig().GetDBConfig())
	repository.RunMigration(config.GetConfig().GetDBConfig())
	db, err := repository.GetDB()
	if err != nil {
		log.Fatal("fail to get db, ", err)
	}
	userRepo := repository.NewUserRepo(db)

	//service
	userSvc := service.NewUserService(userRepo)

	//handler
	handler := handler.NewUserHandler(userSvc)

	//cmd := newCommand()
}
