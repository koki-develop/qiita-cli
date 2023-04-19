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
		cfg, err := loadConfig()
		if err != nil {
			return err
		}

		p, err := printers.Get(*flagFormat.Get(cmd, true))
		if err != nil {
			return err
		}

		cl := qiita.New(cfg.AccessToken)

		params := &qiita.ListItemsParameters{
			Page:    flagPage.Get(cmd, true),
			PerPage: flagPerPage.Get(cmd, true),
			Query:   flagItemsSearchQuery.Get(cmd, false),
		}
		items, err := cl.ListItems(params)
		if err != nil {
			return err
		}

		if err := p.Print(os.Stdout, *flagItemColumns.Get(cmd, true), items); err != nil {
			return err
		}

		return nil
	},
}

var itemsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list own items",
	Long:  "list own items.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig()
		if err != nil {
			return err
		}

		p, err := printers.Get(*flagFormat.Get(cmd, true))
		if err != nil {
			return err
		}

		cl := qiita.New(cfg.AccessToken)

		params := &qiita.ListAuthenticatedUserItemsParameters{
			Page:    flagPage.Get(cmd, true),
			PerPage: flagPerPage.Get(cmd, true),
		}
		items, err := cl.ListAuthenticatedUserItems(params)
		if err != nil {
			return err
		}

		if err := p.Print(os.Stdout, *flagItemColumns.Get(cmd, true), items); err != nil {
			return err
		}
		return nil
	},
}

var itemsGetCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "get an item",
	Long:  "get an item.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		cfg, err := loadConfig()
		if err != nil {
			return err
		}

		p, err := printers.Get(*flagFormat.Get(cmd, true))
		if err != nil {
			return err
		}

		cl := qiita.New(cfg.AccessToken)

		item, err := cl.GetItem(id)
		if err != nil {
			return err
		}

		if err := p.Print(os.Stdout, *flagItemColumns.Get(cmd, true), item); err != nil {
			return err
		}
		return nil
	},
}

var itemsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create an item",
	Long:  "create an item.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig()
		if err != nil {
			return err
		}

		p, err := printers.Get(*flagFormat.Get(cmd, true))
		if err != nil {
			return err
		}

		cl := qiita.New(cfg.AccessToken)

		params := &qiita.CreateItemParameters{}
		if flagItemsCreateTitle.Changed(cmd) {
			params.Title = flagItemsCreateTitle.Get(cmd, false)
		}
		if flagItemsCreateTags.Changed(cmd) {
			tags := qiita.TagsFromStrings(*flagItemsCreateTags.Get(cmd, false))
			params.Tags = &tags
		}
		if flagItemsCreateBody.Changed(cmd) {
			params.Body = flagItemsCreateBody.Get(cmd, false)
		}
		if flagItemsCreatePrivate.Changed(cmd) {
			params.Private = flagItemsCreatePrivate.Get(cmd, false)
		}
		if flagItemsCreateTweet.Changed(cmd) {
			params.Tweet = flagItemsCreateTweet.Get(cmd, false)
		}

		item, err := cl.CreateItem(params)
		if err != nil {
			return err
		}

		if err := p.Print(os.Stdout, *flagItemColumns.Get(cmd, true), item); err != nil {
			return err
		}
		return nil
	},
}
