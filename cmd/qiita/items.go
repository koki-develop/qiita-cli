package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/koki-develop/qiita-cli/internal/cli"
	"github.com/koki-develop/qiita-cli/internal/qiita"
	"github.com/koki-develop/qiita-cli/internal/util"
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
	Use:   "delete [id]",
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
		filename := args[0]
		// 末尾が .md でない場合は .md を付ける
		if !strings.HasSuffix(filename, ".md") {
			filename += ".md"
		}

		// もしファイルが既に存在している場合はエラーを返す
		if _, err := os.Stat(filename); err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		} else {
			return fmt.Errorf("file already exists: %s", filename)
		}

		f, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer f.Close()

		fm := qiita.ItemFrontMatter{
			Title:   flagItemsNewTitle.Get(cmd, true),
			Tags:    flagItemsNewTags.Get(cmd, true),
			Private: flagItemsNewPrivate.Get(cmd, true),
		}

		if err := util.WriteMarkdown(f, "", fm); err != nil {
			return err
		}

		fmt.Printf("Created: %s\n", filename)

		return nil
	},
}

var itemsPullCmd = &cobra.Command{
	Use:   "pull [ids]",
	Short: "Download items",
	Long:  "Download items.",
	RunE: func(cmd *cobra.Command, args []string) error {
		all := *flagItemsPullAll.Get(cmd, true)
		if all && len(args) > 0 {
			return ErrIDsWithAll
		}

		cfg, err := loadConfig()
		if err != nil {
			return err
		}

		cl := qiita.New(cfg.AccessToken)

		fmt.Println("Pulling items...")

		var items qiita.Items
		if all {
			for i := 0; i < 100; i++ {
				p := &qiita.ListAuthenticatedUserItemsParameters{
					PerPage: util.Int(100),
					Page:    util.Int(i + 1),
				}
				is, err := cl.ListAuthenticatedUserItems(p)
				if err != nil {
					return err
				}
				items = append(items, is...)
				if len(is) < 100 {
					break
				}
			}
		} else {
			for _, id := range args {
				item, err := cl.GetItem(id)
				if err != nil {
					return err
				}
				items = append(items, item)
			}
		}

		out := *flagItemsPullOut.Get(cmd, true)
		if err := os.MkdirAll(out, os.ModePerm); err != nil {
			return err
		}

		for _, item := range items {
			filename := path.Join(out, fmt.Sprintf("%s.md", strings.ReplaceAll(item.Title(), "/", "_")))
			f, err := util.CreateFile(filename)
			if err != nil {
				return err
			}
			defer f.Close()
			if err := util.WriteMarkdown(f, item.Body(), item.FrontMatter()); err != nil {
				return err
			}
		}

		fmt.Println("Done.")
		return nil
	},
}
