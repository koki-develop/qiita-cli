package main

import (
	"github.com/koki-develop/qiita-cli/internal/config"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use: "configure",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := config.Configure(&config.Config{}); err != nil {
			return err
		}

		return nil
	},
}
