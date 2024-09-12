package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Escape-Technologies/escape-api-client/api"
	"github.com/Escape-Technologies/escape-api-client/internal"

	"github.com/spf13/cobra"
)

var organizationApplicationsCmd = &cobra.Command{
	Use:   "organization-applications [organizationId]",
	Short: "Get a list of applications bound to an organization",
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		client, err := api.NewClientWithResponses(
			internal.APIEndpoint,
			api.WithHTTPClient(&http.Client{
				Transport: internal.NewTransportWithAuth(apiKey),
			}),
		)
		if err != nil {
			return fmt.Errorf("error while creating client: %s", err)
		}
		resp, err := client.GetOrganizationOrganizationIdApplicationsWithResponse(
			context.Background(),
			args[0],
		)
		if err != nil {
			return fmt.Errorf("error while getting applications: %s", err)
		}
		if resp.StatusCode() != http.StatusOK {
			return fmt.Errorf("error: %s", resp.Status())
		}

		fmt.Printf("%s\n", resp.Body)
		return nil
	},
}

func init() { // nolint: gochecknoinits
	organizationApplicationsCmd.Flags().
		StringVar(&apiKey, "api-key", "", "API key for authorization")
	rootCmd.AddCommand(organizationApplicationsCmd)
}
