package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"github.com/koki-develop/qiita-cli/internal/printers"
	"github.com/spf13/cobra"
)

var (
	version string
)

var (
	flagFormat string
)

var rootCmd = &cobra.Command{
	Use:          "qiita",
	Short:        "CLI for Qiita",
	Long:         "CLI for Qiita",
	SilenceUsage: true,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	/*
	 * version
	 */
	if version == "" {
		if info, ok := debug.ReadBuildInfo(); ok {
			version = info.Main.Version
		}
	}
	rootCmd.Version = version

	/*
	 * commands
	 */

	// items
	rootCmd.AddCommand(itemsCmd)
	itemsCmd.AddCommand(
		itemsSearchCmd,
	)

	/*
	 * flags
	 */
	for _, cmd := range []*cobra.Command{rootCmd, itemsCmd, itemsSearchCmd} {
		cmd.Flags().SortFlags = false
	}

	// format
	for _, cmd := range []*cobra.Command{itemsSearchCmd} {
		cmd.Flags().StringVarP(&flagFormat, "format", "f", "json", fmt.Sprintf("output format (%s)", strings.Join(printers.ListFormats(), "|")))
	}

	// items
	for _, cmd := range []*cobra.Command{itemsSearchCmd} {
		cmd.Flags().StringSliceVar(&flagItemColumns, "columns", []string{"id", "title", "user", "url"}, "properties that are going to be presented as columns (table format only)")
	}

	// items search
	itemsSearchCmd.Flags().IntVar(&flagItemsSearchPage, "page", 1, "page number (from 1 to 100)")
	itemsSearchCmd.Flags().IntVar(&flagItemsSearchPerPage, "per-page", 100, "records count per page (from 1 to 100)")
	itemsSearchCmd.Flags().StringVarP(&flagItemsSearchQuery, "query", "q", "", "search query")
}
