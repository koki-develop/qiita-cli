package cli

import (
	"os"

	"github.com/koki-develop/qiita-cli/internal/flags"
	"github.com/koki-develop/qiita-cli/internal/qiita"
	"github.com/koki-develop/qiita-cli/internal/util"
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

type ItemsListParameters struct {
	FlagPage    *flags.Int // --page
	FlagPerPage *flags.Int // --per-page
}

func (c *CLI) ItemsList(params *ItemsListParameters) error {
	items, err := c.client.ListAuthenticatedUserItems(&qiita.ListAuthenticatedUserItemsParameters{
		Page:    params.FlagPage.Get(c.command, true),
		PerPage: params.FlagPerPage.Get(c.command, true),
	})
	if err != nil {
		return err
	}

	if err := c.printer.Print(c.writer, items); err != nil {
		return err
	}

	return nil
}

type ItemsCreateParameters struct {
	FlagFile    *flags.String      // --file
	FlagWrite   *flags.Bool        // --write
	FlagTitle   *flags.String      // --title
	FlagTags    *flags.StringSlice // --tags
	FlagBody    *flags.String      // --body
	FlagPrivate *flags.Bool        // --private
	FlagTweet   *flags.Bool        // --tweet
}

// $ qiita items create
func (c *CLI) ItemsCreate(params *ItemsCreateParameters) error {
	if params.FlagWrite.Changed(c.command) && !params.FlagFile.Changed(c.command) {
		return ErrWriteWithoutFile
	}
	file := params.FlagFile.Get(c.command, false)

	p := &qiita.CreateItemParameters{}

	if file != nil {
		md, fm, err := c.readMarkdown(*file)
		if err != nil {
			return err
		}
		if fm.ID != nil {
			return ErrCreateWithID
		}
		p.Title = fm.Title
		p.Tags = fm.QiitaTags()
		p.Body = &md
		p.Private = fm.Private
	}

	if params.FlagTitle.Changed(c.command) {
		p.Title = params.FlagTitle.Get(c.command, true)
	}
	if params.FlagTags.Changed(c.command) {
		p.Tags = util.Ptr(qiita.TagsFromStrings(*params.FlagTags.Get(c.command, true)))
	}
	if params.FlagBody.Changed(c.command) {
		p.Body = params.FlagBody.Get(c.command, true)
	}
	if params.FlagPrivate.Changed(c.command) {
		p.Private = params.FlagPrivate.Get(c.command, true)
	}
	if params.FlagTweet.Changed(c.command) {
		p.Tweet = params.FlagTweet.Get(c.command, true)
	}

	item, err := c.client.CreateItem(p)
	if err != nil {
		return err
	}

	if *params.FlagWrite.Get(c.command, true) {
		if err := c.writeMarkdown(*file, item); err != nil {
			return err
		}
	}

	if err := c.printer.Print(c.writer, item); err != nil {
		return err
	}

	return nil
}

func (c *CLI) readMarkdown(file string) (string, qiita.ItemFrontMatter, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", qiita.ItemFrontMatter{}, err
	}
	defer f.Close()

	var fm qiita.ItemFrontMatter
	md, err := util.ReadMarkdown(f, &fm)
	if err != nil {
		return "", qiita.ItemFrontMatter{}, err
	}

	return md, fm, nil
}

func (c *CLI) writeMarkdown(file string, item qiita.Item) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := util.WriteMarkdown(f, item.Body(), item.FrontMatter()); err != nil {
		return err
	}

	return nil
}
