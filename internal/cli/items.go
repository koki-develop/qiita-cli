package cli

import (
	"fmt"
	"os"
	"strings"

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

// $ qiita items list
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

type ItemsGetParameters struct {
	Args []string
}

// $ qiita items get
func (c *CLI) ItemsGet(params *ItemsGetParameters) error {
	id := params.Args[0]

	item, err := c.client.GetItem(id)
	if err != nil {
		return err
	}

	if err := c.printer.Print(c.writer, item); err != nil {
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

type ItemsUpdateParameters struct {
	Args        []string
	FlagFile    *flags.String      // --file
	FlagWrite   *flags.Bool        // --write
	FlagTitle   *flags.String      // --title
	FlagTags    *flags.StringSlice // --tags
	FlagBody    *flags.String      // --body
	FlagPrivate *flags.Bool        // --private
}

func (c *CLI) ItemsUpdate(params *ItemsUpdateParameters) error {
	if params.FlagWrite.Changed(c.command) && !params.FlagFile.Changed(c.command) {
		return ErrWriteWithoutFile
	}
	file := params.FlagFile.Get(c.command, false)

	var id string
	p := &qiita.UpdateItemParameters{}

	if file != nil {
		md, fm, err := c.readMarkdown(*file)
		if err != nil {
			return err
		}
		if fm.ID != nil {
			id = *fm.ID
		}
		p.Title = fm.Title
		p.Tags = fm.QiitaTags()
		p.Body = &md
		p.Private = fm.Private
	}

	if len(params.Args) > 0 {
		id = params.Args[0]
	}
	if id == "" {
		return ErrIDRequired
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

	item, err := c.client.UpdateItem(id, p)
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

type ItemsDeleteParameters struct {
	Args []string
}

func (c *CLI) ItemsDelete(params *ItemsDeleteParameters) error {
	for _, id := range params.Args {
		if err := c.client.DeleteItem(id); err != nil {
			return err
		}
	}

	return nil
}

type ItemsNewParameters struct {
	Args        []string
	FlagTitle   *flags.String      // --title
	FlagTags    *flags.StringSlice // --tags
	FlagPrivate *flags.Bool        // --private
}

func (c *CLI) ItemsNew(params *ItemsNewParameters) error {
	filename := params.Args[0]
	if !strings.HasSuffix(filename, ".md") {
		filename += ".md"
	}

	ext, err := util.Exists(filename)
	if err != nil {
		return err
	}
	if ext {
		return fmt.Errorf("file already exists: %s", filename)
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	fm := qiita.ItemFrontMatter{
		Title:   params.FlagTitle.Get(c.command, true),
		Tags:    params.FlagTags.Get(c.command, true),
		Private: params.FlagPrivate.Get(c.command, true),
	}

	if err := util.WriteMarkdown(f, "", fm); err != nil {
		return err
	}

	fmt.Fprintf(c.writer, "Created: %s\n", filename)
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
