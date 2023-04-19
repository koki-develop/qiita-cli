package main

import (
	"fmt"
	"strings"

	"github.com/koki-develop/qiita-cli/internal/config"
	"github.com/koki-develop/qiita-cli/internal/flags"
	"github.com/koki-develop/qiita-cli/internal/printers"
)

// common
var (
	// --format
	flagFormat = &flags.String{Flag: &flags.Flag{
		Name:        "format",
		Description: fmt.Sprintf("output format (%s)", strings.Join(printers.ListFormats(), "|"))},
		Default: func() string {
			cfg, err := config.Load()
			if err != nil {
				return "table"
			}
			if cfg.Format != "" {
				return cfg.Format
			}
			return "table"
		}(),
	}

	// --page
	flagPage = &flags.Int{Flag: &flags.Flag{
		Name:        "page",
		Description: "page number (from 1 to 100)"},
		Default: 1,
	}

	// --per-page
	flagPerPage = &flags.Int{Flag: &flags.Flag{
		Name:        "per-page",
		Description: "records count per page (from 1 to 100)"},
		Default: 100,
	}
)

// configure
var (
	flagConfigureAccessToken = &flags.String{Flag: &flags.Flag{
		Name:        "access-token",
		Description: "qiita access token"},
	}
	flagConfigureFormat = &flags.String{Flag: &flags.Flag{
		Name:        "format",
		Description: "default output format"},
	}
)

// items
var (
	// --columns
	flagItemColumns = &flags.StringSlice{Flag: &flags.Flag{
		Name:        "columns",
		Description: "properties that are going to be presented as columns (table format only)"},
		Default: []string{"id", "title", "user", "url"},
	}
)

// items search
var (
	// --query
	flagItemsSearchQuery = &flags.String{Flag: &flags.Flag{
		Name:        "query",
		ShortName:   "q",
		Description: "search query"},
	}
)

// items create
var (
	// --file
	flagItemsCreateFile = &flags.String{Flag: &flags.Flag{
		Name:        "file",
		Description: "create an item from a file",
	}}

	// --write
	flagItemsCreateWrite = &flags.Bool{Flag: &flags.Flag{
		Name:        "write",
		Description: "write information about the created item to a file",
	}}

	// --title
	flagItemsCreateTitle = &flags.String{Flag: &flags.Flag{
		Name:        "title",
		Description: "the title of the item",
	}}

	// --tags
	flagItemsCreateTags = &flags.StringSlice{Flag: &flags.Flag{
		Name:        "tags",
		Description: "a list of tags"},
	}

	// --body
	flagItemsCreateBody = &flags.String{Flag: &flags.Flag{
		Name:        "body",
		Description: "item body in markdown"},
	}

	// --private
	flagItemsCreatePrivate = &flags.Bool{Flag: &flags.Flag{
		Name:        "private",
		Description: "whether the item is private"},
	}

	// --tweet
	flagItemsCreateTweet = &flags.Bool{Flag: &flags.Flag{
		Name:        "tweet",
		Description: "whether to post a tweet"},
	}
)

// items update
var (
	// --title
	flagItemsUpdateTitle = &flags.String{Flag: &flags.Flag{
		Name:        "title",
		Description: "the title of the item",
	}}
	// --tags
	flagItemsUpdateTags = &flags.StringSlice{Flag: &flags.Flag{
		Name:        "tags",
		Description: "a list of tags"},
	}
	// --body
	flagItemsUpdateBody = &flags.String{Flag: &flags.Flag{
		Name:        "body",
		Description: "item body in markdown"},
	}
	// --private
	flagItemsUpdatePrivate = &flags.Bool{Flag: &flags.Flag{
		Name:        "private",
		Description: "whether the item is private"},
	}
)
