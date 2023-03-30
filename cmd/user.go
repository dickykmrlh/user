package main

import (
	"log"
	"os"

	"github.com/dickykmrlh/user/cmd/app"
	"github.com/dickykmrlh/user/config"
	"github.com/dickykmrlh/user/database"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "/main [sub]",
		Short: "user service",
	}
}

func run(cmd *cobra.Command) int {
	if err := cmd.Execute(); err != nil {
		return 1
	}

	return 0
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("fail to load env variables, ", err)
	}
	config.Init()

	database.Init()

	server, err := app.NewServer()
	if err != nil {
		log.Fatal("fail to create new server, ", err)
	}

	cmd := newRootCommand()
	cmd.AddCommand(&cobra.Command{
		Use:   "server",
		Short: "starting web server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return server.Run()
		},
	})

	os.Exit(run(cmd))
}
