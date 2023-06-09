package main

import (
	"os"
	"runtime/debug"

	"github.com/koki-develop/qiita-cli/internal/cli"
	"github.com/koki-develop/qiita-cli/internal/flags"
	"github.com/koki-develop/qiita-cli/internal/notify"
	"github.com/spf13/cobra"
)

var (
	version string
)

var rootCmd = &cobra.Command{
	Use:          "qiita",
	Short:        "CLI for Qiita",
	Long:         "CLI for Qiita.",
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
		itemsPushCmd,   // items push
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
		flagItemsColumns,     // --columns
		flagPage,             // --page
		flagPerPage,          // --per-page
		flagItemsSearchQuery, // --query
	}.AddTo(itemsSearchCmd)

	/* items list */
	flags.Flags{
		flagFormat,       // --format
		flagItemsColumns, // --columns
		flagPage,         // --page
		flagPerPage,      // --per-page
	}.AddTo(itemsListCmd)

	/* items get */
	flags.Flags{
		flagFormat,       // --format
		flagItemsColumns, // --columns
	}.AddTo(itemsGetCmd)

	/* items create */
	flags.Flags{
		flagFormat,             // --format
		flagItemsColumns,       // --columns
		flagItemsCreateTitle,   // --title
		flagItemsCreateTags,    // --tags
		flagItemsCreateBody,    // --body
		flagItemsCreatePrivate, // --private
		flagItemsCreateTweet,   // --tweet
	}.AddTo(itemsCreateCmd)

	/* items update */
	flags.Flags{
		flagFormat,             // --format
		flagItemsColumns,       // --columns
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

	/* items push */
	flags.Flags{
		flagItemsPushWrite, // --write
	}.AddTo(itemsPushCmd)

	/* items pull */
	flags.Flags{
		flagItemsPullOut, // --out
		flagItemsPullAll, // --all
	}.AddTo(itemsPullCmd)

	_ = notify.NotifyNewRelease(os.Stderr, version)
}

func disableSortFlags(cmd *cobra.Command) {
	cmd.Flags().SortFlags = false
	for _, child := range cmd.Commands() {
		disableSortFlags(child)
	}
}

func newCLI(cmd *cobra.Command, columns *flags.StringSlice) (*cli.CLI, error) {
	return cli.New(&cli.Config{
		Command:     cmd,
		Writer:      os.Stdout,
		ErrWriter:   os.Stderr,
		FlagFormat:  flagFormat, // --format
		FlagColumns: columns,    // --columns
	})
}
