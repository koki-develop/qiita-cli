package main

import (
	"github.com/spf13/cobra"
)

var itemsCmd = &cobra.Command{
	Use:   "items",
	Short: "Manage items",
	Long:  "Manage items.",
}

var itemsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list items",
	Long:  "list items.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
