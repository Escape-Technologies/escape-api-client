package cmd

import (
	"fmt"
	"os"

	"github.com/Escape-Technologies/escape-api-client/internal"
	"github.com/spf13/cobra"
)

var (
	apiKey string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "escape-api-client",
	Short: "Escape API CLI Client",
	Long: `Escape API CLI Client is a command-line interface client built using Cobra for interacting with the Escape API.

This application allows users to fetch scan details, list applications in an organization, start scans, and upload introspection data.`,
	Version: fmt.Sprintf(
		"%s, built on %s",
		internal.BuildVersion,
		internal.BuildDate,
	),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() { // nolint: gochecknoinits
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.escape-api-client.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
