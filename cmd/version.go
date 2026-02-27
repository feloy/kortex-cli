package cmd

import (
	"github.com/spf13/cobra"
)

var Version = "0.1.0-next"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of kortex-cli",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("kortex-cli version %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
