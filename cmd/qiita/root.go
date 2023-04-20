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
		itemsGetCmd,    // items get
		itemsCreateCmd, // items create
		itemsUpdateCmd, // items update
		itemsDeleteCmd, // items delete
		itemsNewCmd,    // items new
		itemsPullCmd,   // items pull
	)

	/*
	 * flags
	 */
	disableSortFlags(rootCmd)

	/* configure */
	flags.Flags{
		flagConfigureAccessToken, // --access-token
		flagConfigureFormat,      // --format
	}.AddTo(configureCmd)

	/* items search */
	flags.Flags{
		flagFormat,           // --format
		flagItemColumns,      // --columns
		flagPage,             // --page
		flagPerPage,          // --per-page
		flagItemsSearchQuery, // --query
	}.AddTo(itemsSearchCmd)

	/* items list */
	flags.Flags{
		flagFormat,      // --format
		flagItemColumns, // --columns
		flagPage,        // --page
		flagPerPage,     // --per-page
	}.AddTo(itemsListCmd)

	/* items get */
	flags.Flags{
		flagFormat,      // --format
		flagItemColumns, // --columns
	}.AddTo(itemsGetCmd)

	/* items create */
	flags.Flags{
		flagFormat,             // --format
		flagItemColumns,        // --columns
		flagItemsCreateFile,    // --file
		flagItemsCreateWrite,   // --write
		flagItemsCreateTitle,   // --title
		flagItemsCreateTags,    // --tags
		flagItemsCreateBody,    // --body
		flagItemsCreatePrivate, // --private
		flagItemsCreateTweet,   // --tweet
	}.AddTo(itemsCreateCmd)

	/* items update */
	flags.Flags{
		flagFormat,             // --format
		flagItemColumns,        // --columns
		flagItemsUpdateFile,    // --file
		flagItemsUpdateWrite,   // --write
		flagItemsUpdateTitle,   // --title
		flagItemsUpdateTags,    // --tags
		flagItemsUpdateBody,    // --body
		flagItemsUpdatePrivate, // --private
	}.AddTo(itemsUpdateCmd)

	/* items new */
	flags.Flags{
		flagItemsNewTitle,   // --title
		flagItemsNewTags,    // --tags
		flagItemsNewPrivate, // --private
	}.AddTo(itemsNewCmd)

	/* items pull */
	flags.Flags{
		flagItemsPullOut, // --out
	}.AddTo(itemsPullCmd)
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
