package main

import (
	"log"
	"os"

	"github.com/dickykmrlh/user/api/server"
	"github.com/dickykmrlh/user/config"
	"github.com/dickykmrlh/user/database"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("fail to load env variables, ", err)
	}
	config.Init()

	database.Init()

	server, err := server.New()
	if err != nil {
		log.Fatal("fail to create new server, ", err)
	}

	cmd := RootCommand()
	cmd.add(&cobra.Command{
		Use:   "server",
		Short: "starting web server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return server.Run()
		},
	})

	os.Exit(cmd.run())
}
