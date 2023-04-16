package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "qiita",
	Short:        "CLI for Qiita",
	Long:         "CLI for Qiita",
	SilenceUsage: true,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(
		itemsCmd,
	)

	itemsCmd.AddCommand(
		itemsListCmd,
	)
}
