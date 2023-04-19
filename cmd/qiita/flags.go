package main

import (
	"fmt"
	"strings"

	"github.com/koki-develop/qiita-cli/internal/flags"
	"github.com/koki-develop/qiita-cli/internal/printers"
)

// common
var (
	// --format
	flagFormat = &flags.String{Flag: &flags.Flag{
		Name:        "format",
		Description: fmt.Sprintf("output format (%s)", strings.Join(printers.ListFormats(), "|"))},
		Default: "table",
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
	// --page
	flagItemsSearchPage = &flags.Int{Flag: &flags.Flag{
		Name:        "page",
		Description: "page number (from 1 to 100)"},
		Default: 1,
	}

	// --per-page
	flagItemsSearchPerPage = &flags.Int{Flag: &flags.Flag{
		Name:        "per-page",
		Description: "records count per page (from 1 to 100)"},
		Default: 100,
	}

	// --query
	flagItemsSearchQuery = &flags.String{Flag: &flags.Flag{
		Name:        "query",
		ShortName:   "q",
		Description: "search query"},
	}
)
