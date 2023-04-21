package main

import (
	"github.com/koki-develop/qiita-cli/internal/config"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure Qiita CLI",
	Long:  "Configure Qiita CLI.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := &config.Config{
			AccessToken: *flagConfigureAccessToken.Get(cmd, true),
			Format:      *flagConfigureFormat.Get(cmd, true),
		}
		if err := config.Configure(cfg); err != nil {
			return err
		}

		return nil
	},
}
