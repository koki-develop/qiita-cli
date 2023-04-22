package cli

import (
	"fmt"
	"io"

	"github.com/koki-develop/qiita-cli/internal/config"
	"github.com/koki-develop/qiita-cli/internal/flags"
	"github.com/koki-develop/qiita-cli/internal/printers"
	"github.com/koki-develop/qiita-cli/internal/qiita"
	"github.com/koki-develop/qiita-cli/internal/util"
	"github.com/spf13/cobra"
)

type CLI struct {
	command *cobra.Command

	writer    io.Writer
	errWriter io.Writer

	config  *config.Config
	client  *qiita.Client
	printer printers.Printer
}

type Config struct {
	Command *cobra.Command

	Writer    io.Writer
	ErrWriter io.Writer

	FlagFormat  *flags.String      // --format
	FlagColumns *flags.StringSlice // --columns
}

func New(cfg *Config) (*CLI, error) {
	c := &CLI{
		command:   cfg.Command,
		writer:    cfg.Writer,
		errWriter: cfg.ErrWriter,
	}

	qiitacfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(c.errWriter, "Failed to load config.\nIf you have not configured yet, please run `qiita configure`.")
		return nil, err
	}
	c.config = qiitacfg
	c.client = qiita.New(c.config.AccessToken)

	if cfg.FlagFormat != nil && cfg.FlagColumns != nil {
		f := cfg.FlagFormat.Get(cfg.Command, false)
		if f == nil {
			f = &qiitacfg.Format
		}
		if *f == "" {
			f = util.Ptr("table")
		}

		p, err := printers.Get(*f)
		if err != nil {
			return nil, err
		}
		c.printer = p
		c.printer.SetColumns(*cfg.FlagColumns.Get(cfg.Command, true))
	}

	return c, nil
}
