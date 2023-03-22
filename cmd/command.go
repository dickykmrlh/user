package main

import (
	"github.com/spf13/cobra"
)

type command struct {
	cm *cobra.Command
}

func newCommand() *command {
	return &command{
		cm: &cobra.Command{
			Use:   "/main [sub]",
			Short: "user service",
		},
	}
}

func (c *command) addCommand(use, shortDesc string, runFunc func() error) {
	c.cm.AddCommand(&cobra.Command{
		Use:   use,
		Short: shortDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runFunc()
		},
	})
}

func (c *command) run() int {
	if err := c.cm.Execute(); err != nil {
		return 1
	}

	return 0
}
