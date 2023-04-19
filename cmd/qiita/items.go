package main

import (
	"errors"
	"os"

	"github.com/koki-develop/qiita-cli/internal/printers"
	"github.com/koki-develop/qiita-cli/internal/qiita"
	"github.com/koki-develop/qiita-cli/internal/util"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type itemFrontMatter struct {
	ID      *string   `yaml:"id,omitempty"`
	Title   *string   `yaml:"title,omitempty"`
	Tags    *[]string `yaml:"tags,flow,omitempty"`
	Private *bool     `yaml:"private,omitempty"`
}

func (fm *itemFrontMatter) QiitaTags() *qiita.Tags {
	if fm.Tags == nil {
		return nil
	}

	var tags qiita.Tags
	for _, t := range *fm.Tags {
		tags = append(tags, &qiita.Tag{Name: t})
	}
	return &tags
}

var itemsCmd = &cobra.Command{
	Use:   "items",
	Short: "Manage items",
	Long:  "Manage items.",
}

var itemsSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search items",
	Long:  "Search items.",
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
	Short: "List own items",
	Long:  "List own items.",
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
	Short: "Get an item",
	Long:  "Get an item.",
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
	Short: "Create an item",
	Long:  "Create an item.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig()
		if err != nil {
			return err
		}

		if flagItemsCreateWrite.Changed(cmd) && !flagItemsCreateFile.Changed(cmd) {
			return errors.New("`--write` can be used when `--file` is set")
		}

		p, err := printers.Get(*flagFormat.Get(cmd, true))
		if err != nil {
			return err
		}

		cl := qiita.New(cfg.AccessToken)

		params := &qiita.CreateItemParameters{}
		if flagItemsCreateFile.Changed(cmd) {
			f, err := os.Open(*flagItemsCreateFile.Get(cmd, true))
			if err != nil {
				return err
			}
			defer f.Close()

			var fm itemFrontMatter
			md, err := util.ReadMarkdown(f, &fm)
			if err != nil {
				return err
			}
			if fm.ID != nil {
				return errors.New("id cannot be set when creating an item")
			}
			params.Title = fm.Title
			params.Tags = fm.QiitaTags()
			params.Body = &md
			params.Private = fm.Private
		}

		if flagItemsCreateTitle.Changed(cmd) {
			params.Title = flagItemsCreateTitle.Get(cmd, true)
		}
		if flagItemsCreateTags.Changed(cmd) {
			tags := qiita.TagsFromStrings(*flagItemsCreateTags.Get(cmd, true))
			params.Tags = &tags
		}
		if flagItemsCreateBody.Changed(cmd) {
			params.Body = flagItemsCreateBody.Get(cmd, true)
		}
		if flagItemsCreatePrivate.Changed(cmd) {
			params.Private = flagItemsCreatePrivate.Get(cmd, true)
		}
		if flagItemsCreateTweet.Changed(cmd) {
			params.Tweet = flagItemsCreateTweet.Get(cmd, true)
		}

		item, err := cl.CreateItem(params)
		if err != nil {
			return err
		}

		if *flagItemsCreateWrite.Get(cmd, true) && flagItemsCreateFile.Changed(cmd) {
			fm := itemFrontMatter{
				ID:      util.String(item.ID()),
				Title:   util.String(item.Title()),
				Tags:    util.Strings(item.Tags().Names()),
				Private: util.Bool(item.Private()),
			}
			y, err := yaml.Marshal(fm)
			if err != nil {
				return err
			}
			f, err := os.Create(*flagItemsCreateFile.Get(cmd, true))
			if err != nil {
				return err
			}
			defer f.Close()
			f.WriteString("---\n")
			f.Write(y)
			f.WriteString("---\n\n")
			f.WriteString(item.Body())
		}

		if err := p.Print(os.Stdout, *flagItemColumns.Get(cmd, true), item); err != nil {
			return err
		}

		return nil
	},
}
