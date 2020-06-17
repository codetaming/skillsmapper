package cmd

import (
	"fmt"
	"github.com/codetaming/skillsmapper/cli/version"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cli",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:\t", version.Version)
		fmt.Println("BuildTime:\t", version.BuildTime)
		fmt.Println("GitCommit:\t", version.GitCommit)
	},
}
