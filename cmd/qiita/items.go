package main

import (
	"github.com/koki-develop/qiita-cli/internal/cli"
	"github.com/spf13/cobra"
)

var itemsCmd = &cobra.Command{
	Use:     "items",
	Aliases: []string{"item"},
	Short:   "Manage items",
	Long:    "Manage items.",
}

var itemsSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search items",
	Long:  "Search items.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newCLI(cmd, flagItemsColumns)
		if err != nil {
			return err
		}

		if err := c.ItemsSearch(&cli.ItemsSearchParameters{
			FlagPage:    flagPage,             // --page
			FlagPerPage: flagPerPage,          // --per-page
			FlagQuery:   flagItemsSearchQuery, // --query
		}); err != nil {
			return err
		}

		return nil
	},
}

var itemsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List own items",
	Long:  "List own items.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newCLI(cmd, flagItemsColumns)
		if err != nil {
			return err
		}

		if err := c.ItemsList(&cli.ItemsListParameters{
			FlagPage:    flagPage,    // --page
			FlagPerPage: flagPerPage, // --per-page
		}); err != nil {
			return err
		}

		return nil
	},
}

var itemsGetCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "Get an item",
	Long:  "Get an item.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newCLI(cmd, flagItemsColumns)
		if err != nil {
			return err
		}

		if err := c.ItemsGet(&cli.ItemsGetParameters{
			Args: args,
		}); err != nil {
			return err
		}

		return nil
	},
}

var itemsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an item",
	Long:  "Create an item.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newCLI(cmd, flagItemsColumns)
		if err != nil {
			return err
		}

		if err := c.ItemsCreate(&cli.ItemsCreateParameters{
			FlagFile:    flagItemsCreateFile,    // --file
			FlagWrite:   flagItemsCreateWrite,   // --write
			FlagTitle:   flagItemsCreateTitle,   // --title
			FlagBody:    flagItemsCreateBody,    // --body
			FlagTags:    flagItemsCreateTags,    // --tags
			FlagPrivate: flagItemsCreatePrivate, // --private
			FlagTweet:   flagItemsCreateTweet,   // --tweet
		}); err != nil {
			return err
		}

		return nil
	},
}

var itemsUpdateCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update an item",
	Long:  "Update an item.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newCLI(cmd, flagItemsColumns)
		if err != nil {
			return err
		}

		if err := c.ItemsUpdate(&cli.ItemsUpdateParameters{
			Args:        args,
			FlagFile:    flagItemsUpdateFile,    // --file
			FlagWrite:   flagItemsUpdateWrite,   // --write
			FlagTitle:   flagItemsUpdateTitle,   // --title
			FlagTags:    flagItemsUpdateTags,    // --tags
			FlagBody:    flagItemsUpdateBody,    // --body
			FlagPrivate: flagItemsUpdatePrivate, // --private
		}); err != nil {
			return err
		}

		return nil
	},
}

var itemsDeleteCmd = &cobra.Command{
	Use:   "delete [ids]",
	Short: "Delete an item",
	Long:  "Delete an item.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newCLI(cmd, nil)
		if err != nil {
			return err
		}

		if err := c.ItemsDelete(&cli.ItemsDeleteParameters{
			Args: args,
		}); err != nil {
			return err
		}

		return nil
	},
}

var itemsNewCmd = &cobra.Command{
	Use:   "new [file]",
	Short: "Create a new item file",
	Long:  "Create a new item file.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newCLI(cmd, nil)
		if err != nil {
			return err
		}

		if err := c.ItemsNew(&cli.ItemsNewParameters{
			Args:        args,
			FlagTitle:   flagItemsNewTitle,   // --title
			FlagTags:    flagItemsNewTags,    // --tags
			FlagPrivate: flagItemsNewPrivate, // --private
		}); err != nil {
			return err
		}

		return nil
	},
}

var itemsPushCmd = &cobra.Command{
	Use:   "push [files]",
	Short: "Upload items",
	Long:  "Upload items.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newCLI(cmd, nil)
		if err != nil {
			return err
		}

		if err := c.ItemsPush(&cli.ItemsPushParameters{
			Args:      args,
			FlagWrite: flagItemsPushWrite, // --write
		}); err != nil {
			return err
		}

		return nil
	},
}

var itemsPullCmd = &cobra.Command{
	Use:   "pull [ids]",
	Short: "Download items",
	Long:  "Download items.",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newCLI(cmd, nil)
		if err != nil {
			return err
		}

		if err := c.ItemsPull(&cli.ItemsPullParameters{
			Args:    args,
			FlagAll: flagItemsPullAll, // --all
			FlagOut: flagItemsPullOut, // --out
		}); err != nil {
			return err
		}

		return nil
	},
}
