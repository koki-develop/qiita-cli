package main

import (
	"os"

	"github.com/koki-develop/qiita-cli/internal/printers"
	"github.com/koki-develop/qiita-cli/internal/qiita"
	"github.com/spf13/cobra"
)

var itemsCmd = &cobra.Command{
	Use:   "items",
	Short: "Manage items",
	Long:  "Manage items.",
}

var itemsSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "search items",
	Long:  "search items.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := printers.Get(flagFormat.Value)
		if err != nil {
			return err
		}

		cl := qiita.New("")

		params := &qiita.ListItemsParameters{
			Page:    flagItemsSearchPage.Get(cmd, true),
			PerPage: flagItemsSearchPerPage.Get(cmd, true),
			Query:   flagItemsSearchQuery.Get(cmd, false),
		}
		items, err := cl.ListItems(params)
		if err != nil {
			return err
		}

		if err := p.Print(os.Stdout, flagItemColumns.Value, items); err != nil {
			return err
		}

		return nil
	},
}
