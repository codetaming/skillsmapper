package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	ServerURI = getEnv("SERVER_URI", "http://localhost:8080")
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "cli is scripts CLI for Skills Mapper",
	Long:  ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
