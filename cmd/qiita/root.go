package main

import (
	"os"
	"runtime/debug"

	"github.com/koki-develop/qiita-cli/internal/flags"
	"github.com/spf13/cobra"
)

var (
	version string
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

	/* items */
	rootCmd.AddCommand(
		configureCmd, // configure
		itemsCmd,     // items
	)
	itemsCmd.AddCommand(
		itemsSearchCmd, // items search
	)

	/*
	 * flags
	 */
	for _, cmd := range []*cobra.Command{rootCmd, itemsCmd, itemsSearchCmd} {
		cmd.Flags().SortFlags = false
	}

	/* common */
	flags.Flags{
		flagFormat, // --format
	}.AddTo(
		// items
		itemsSearchCmd, // items search
	)

	/* items */
	flags.Flags{
		flagItemColumns, // --columns
	}.AddTo(
		itemsSearchCmd, // items search
	)

	/* items search */
	flags.Flags{
		flagItemsSearchPage,    // --page
		flagItemsSearchPerPage, // --per-page
		flagItemsSearchQuery,   // --query
	}.AddTo(itemsSearchCmd)
}
