package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Escape-Technologies/escape-api-client/api"
	"github.com/Escape-Technologies/escape-api-client/internal"

	"github.com/spf13/cobra"
)

var (
	configurationOverride string
	commitHash            string
	introspection         string
)

var startScanCmd = &cobra.Command{
	Use:   "start-scan [applicationId]",
	Short: "Start a scan for an application",
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		if apiKey == "" {
			return errors.New("apiKey is required")
		}

		client, err := api.NewClientWithResponses(
			internal.APIEndpoint,
			api.WithHTTPClient(&http.Client{
				Transport: internal.NewTransportWithAuth(apiKey),
			}),
		)
		if err != nil {
			return fmt.Errorf("error while creating client: %s", err)
		}

		body := api.PostApplicationApplicationIdStartScanJSONRequestBody{
			ConfigurationOverride: &configurationOverride,
			CommitHash:            &commitHash,
			Introspection:         &introspection,
		}
		resp, err := client.PostApplicationApplicationIdStartScanWithResponse(
			context.Background(),
			args[0],
			body,
		)
		if err != nil {
			return fmt.Errorf("error while starting scan: %s", err)
		}
		if resp.StatusCode() != http.StatusOK {
			return fmt.Errorf("error: %s", resp.Status())
		}

		fmt.Printf("%s\n", resp.Body)
		return nil
	},
}

func init() { // nolint: gochecknoinits
	startScanCmd.Flags().
		StringVar(&apiKey, "api-key", "", "API key for authorization")
	startScanCmd.Flags().
		StringVar(&configurationOverride, "config", "", "Configuration override")
	startScanCmd.Flags().StringVar(&commitHash, "commit", "", "Commit hash")
	startScanCmd.Flags().
		StringVar(&introspection, "introspection", "", "Introspection JSON")
	rootCmd.AddCommand(startScanCmd)
}
