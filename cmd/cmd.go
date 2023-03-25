package main

import (
	"github.com/spf13/cobra"
)

type command struct {
	cm *cobra.Command
}

func RootCommand() *command {
	return &command{
		cm: &cobra.Command{
			Use:   "/main [sub]",
			Short: "user service",
		},
	}
}

func (c *command) add(cmd *cobra.Command) {
	c.cm.AddCommand(cmd)
}

func (c *command) run() int {
	if err := c.cm.Execute(); err != nil {
		return 1
	}

	return 0
}
