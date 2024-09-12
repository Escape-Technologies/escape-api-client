package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Escape-Technologies/escape-api-client/api"
	"github.com/Escape-Technologies/escape-api-client/internal"

	"github.com/spf13/cobra"
)

var introspectionResponse string

var uploadIntrospectionCmd = &cobra.Command{
	Use:   "upload-introspection [applicationId]",
	Short: "Upload introspection data for an application",
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

		body := api.PostApplicationsApplicationIdUploadIntrospectionJSONRequestBody{
			IntrospectionResponse: introspectionResponse,
		}
		resp, err := client.PostApplicationsApplicationIdUploadIntrospectionWithResponse(
			context.Background(),
			args[0],
			body,
		)
		if err != nil {
			return fmt.Errorf("error while uploading introspection: %s", err)
		}
		if resp.StatusCode() != http.StatusOK {
			return fmt.Errorf("error: %s", resp.Status())
		}

		fmt.Printf("%s\n", resp.Body)
		return nil
	},
}

func init() { // nolint: gochecknoinits
	uploadIntrospectionCmd.Flags().
		StringVar(&apiKey, "api-key", "", "API key for authorization")
	uploadIntrospectionCmd.Flags().
		StringVar(&introspectionResponse, "introspection-response", "", "Introspection response (required)")
	_ = uploadIntrospectionCmd.MarkFlagRequired("introspection-response")
	rootCmd.AddCommand(uploadIntrospectionCmd)
}
