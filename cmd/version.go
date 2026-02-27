package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "0.1.0-next"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of kortex-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kortex-cli version %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
