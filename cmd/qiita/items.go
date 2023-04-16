package main

import (
	"os"

	"github.com/koki-develop/qiita-cli/internal/printer"
	"github.com/koki-develop/qiita-cli/internal/qiita"
	"github.com/spf13/cobra"
)

var itemsCmd = &cobra.Command{
	Use:   "items",
	Short: "Manage items",
	Long:  "Manage items.",
}

var (
	flagItemColumns []string

	flagItemsSearchPage    int
	flagItemsSearchPerPage int
	flagItemsSearchQuery   string
)

var itemsSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "search items",
	Long:  "search items.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		cl := qiita.New("")

		params := &qiita.ListItemsParameters{}
		if cmd.Flags().Changed("page") {
			params.Page = &flagItemsSearchPage
		}
		if cmd.Flags().Changed("per-page") {
			params.PerPage = &flagItemsSearchPerPage
		}
		if cmd.Flags().Changed("query") {
			params.Query = &flagItemsSearchQuery
		}

		items, err := cl.ListItems(params)
		if err != nil {
			return err
		}

		p, err := printer.Get(flagFormat)
		if err != nil {
			return err
		}

		if err := p.Print(os.Stdout, items); err != nil {
			return err
		}

		return nil
	},
}
