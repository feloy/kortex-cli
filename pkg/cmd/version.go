package cmd

import (
	"github.com/kortex-hub/kortex-cli/pkg/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of kortex-cli",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("kortex-cli version %s\n", version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
