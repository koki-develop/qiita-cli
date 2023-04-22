package cli

import (
	"github.com/koki-develop/qiita-cli/internal/flags"
	"github.com/koki-develop/qiita-cli/internal/qiita"
)

type ItemsSearchParameters struct {
	FlagPage    *flags.Int    // --page
	FlagPerPage *flags.Int    // --per-page
	FlagQuery   *flags.String // --query
}

// $ qiita items search
func (c *CLI) ItemsSearch(params *ItemsSearchParameters) error {
	items, err := c.client.ListItems(&qiita.ListItemsParameters{
		Page:    params.FlagPage.Get(c.command, true),
		PerPage: params.FlagPerPage.Get(c.command, true),
		Query:   params.FlagQuery.Get(c.command, false),
	})
	if err != nil {
		return err
	}

	if err := c.printer.Print(c.writer, items); err != nil {
		return err
	}

	return nil
}
