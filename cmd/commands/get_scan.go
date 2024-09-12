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

var getScanCmd = &cobra.Command{
	Use:   "get-scan [scanId]",
	Short: "Get details of a specific scan",
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
		resp, err := client.GetScansScanIdWithResponse(
			context.Background(),
			args[0],
		)
		if err != nil {
			return fmt.Errorf("error while getting scan: %s", err)
		}
		if resp.StatusCode() != http.StatusOK {
			return fmt.Errorf("error: %s", resp.Status())
		}

		fmt.Printf("%s\n", resp.Body)
		return nil
	},
}

func init() { // nolint: gochecknoinits
	getScanCmd.Flags().
		StringVar(&apiKey, "api-key", "", "API key for authorization")
	rootCmd.AddCommand(getScanCmd)
}
