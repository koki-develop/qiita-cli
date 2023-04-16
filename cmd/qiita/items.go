package main

import "github.com/spf13/cobra"

var itemsCmd = &cobra.Command{
	Use:   "items",
	Short: "Manage items",
	Long:  "Manage items.",
}
