package main

import (
	"os"

	"github.com/koki-develop/qiita-cli/internal/cli"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:     "configure",
	Aliases: []string{"config"},
	Short:   "Configure Qiita CLI",
	Long:    "Configure Qiita CLI.",
	Args:    cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cli.Configure(&cli.ConfigureParameters{
			Command:         cmd,
			Writer:          os.Stdout,
			FlagAccessToken: flagConfigureAccessToken, // --access-token
			FlagFormat:      flagConfigureFormat,      // --format
		})
		if err != nil {
			return err
		}

		return nil
	},
}
