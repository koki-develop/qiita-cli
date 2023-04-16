package main

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "qiita",
	Short: "CLI for Qiita",
	Long:  "CLI for Qiita",
}

func init() {
	rootCmd.AddCommand(
		itemsCmd,
	)
}
