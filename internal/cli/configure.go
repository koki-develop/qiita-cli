package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/koki-develop/qiita-cli/internal/config"
	"github.com/koki-develop/qiita-cli/internal/flags"
	"github.com/koki-develop/qiita-cli/internal/printers"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"gopkg.in/yaml.v3"
)

type ConfigureParameters struct {
	Command         *cobra.Command
	Writer          io.Writer
	FlagAccessToken *flags.String // --access-token
	FlagFormat      *flags.String // --format
}

func Configure(params *ConfigureParameters) error {
	cfg, err := config.Load()
	if err != nil {
		cfg = &config.Config{}
	}

	interactive := !params.FlagAccessToken.Changed(params.Command) && !params.FlagFormat.Changed(params.Command)
	if interactive {
		// access token
		fmt.Fprint(params.Writer, "Qiita Access Token:")
		tkn, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return err
		}
		fmt.Fprintln(params.Writer)
		cfg.AccessToken = string(tkn)

		// format
		var format string
		fmt.Fprintf(params.Writer, "Default Output Format (%s): ", strings.Join(printers.ListFormats(), "|"))
		if _, err := fmt.Scanln(&format); err != nil {
			return err
		}
		cfg.Format = format
	} else {
		if params.FlagAccessToken.Changed(params.Command) {
			cfg.AccessToken = *params.FlagAccessToken.Get(params.Command, false)
		}
		if params.FlagFormat.Changed(params.Command) {
			cfg.Format = *params.FlagFormat.Get(params.Command, false)
		}
	}

	y, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	p, err := config.Path()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return err
	}

	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(y); err != nil {
		return err
	}

	fmt.Fprintf(params.Writer, "Configure successfully: %s\n", p)
	return nil
}
