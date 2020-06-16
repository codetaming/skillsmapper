package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of smctl",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("smctl 0.0.1")
	},
}
