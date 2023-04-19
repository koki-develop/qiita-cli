package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/koki-develop/qiita-cli/internal/config"
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

	rootCmd.AddCommand(
		configureCmd, // configure
		itemsCmd,     // items
	)
	/* items */
	itemsCmd.AddCommand(
		itemsSearchCmd, // items search
		itemsListCmd,   // items list
	)

	/*
	 * flags
	 */
	disableSortFlags(rootCmd)

	/* common */
	flags.Flags{
		flagFormat, // --format
	}.AddTo(
		// items
		itemsSearchCmd, // items search
		itemsListCmd,   // items list
	)

	/* configure */
	flags.Flags{
		flagConfigureAccessToken, // --access-token
		flagConfigureFormat,      // --format
	}.AddTo(configureCmd)

	/* items */
	flags.Flags{
		flagItemColumns, // --columns
	}.AddTo(
		itemsSearchCmd, // items search
		itemsListCmd,   // items list
	)

	/* items search */
	flags.Flags{
		flagItemsSearchPage,    // --page
		flagItemsSearchPerPage, // --per-page
		flagItemsSearchQuery,   // --query
	}.AddTo(itemsSearchCmd)

	/* items list */
	flags.Flags{
		flagItemsListPage,    // --page
		flagItemsListPerPage, // --per-page
	}.AddTo(itemsListCmd)
}

func loadConfig() (*config.Config, error) {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to load config.\nIf you have not configured yet, please run `qiita configure`.")
		return nil, err
	}

	return cfg, nil
}

func disableSortFlags(cmd *cobra.Command) {
	cmd.Flags().SortFlags = false
	for _, child := range cmd.Commands() {
		disableSortFlags(child)
	}
}
